"""
--- Day 7: No Space Left On Device ---

--- Part One ---
You can hear birds chirping and raindrops hitting leaves as the expedition proceeds. Occasionally, you can even hear much louder sounds in the distance; how big do the animals get out here, anyway?

The device the Elves gave you has problems with more than just its communication system. You try to run a system update:

$ system-update --please --pretty-please-with-sugar-on-top
Error: No space left on device
Perhaps you can delete some files to make space for the update?

You browse around the filesystem to assess the situation and save the resulting terminal output (your puzzle input). For example:

$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
The filesystem consists of a tree of files (plain data) and directories (which can contain other directories or files). The outermost directory is called /. You can navigate around the filesystem, moving into or out of directories and listing the contents of the directory you're currently in.

Within the terminal output, lines that begin with $ are commands you executed, very much like some modern computers:

cd means change directory. This changes which directory is the current directory, but the specific result depends on the argument:
cd x moves in one level: it looks in the current directory for the directory named x and makes it the current directory.
cd .. moves out one level: it finds the directory that contains the current directory, then makes that directory the current directory.
cd / switches the current directory to the outermost directory, /.
ls means list. It prints out all of the files and directories immediately contained by the current directory:
123 abc means that the current directory contains a file named abc with size 123.
dir xyz means that the current directory contains a directory named xyz.
Given the commands and output in the example above, you can determine that the filesystem looks visually like this:

- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - j (file, size=4060174)
    - d.log (file, size=8033020)
    - d.ext (file, size=5626152)
    - k (file, size=7214296)
Here, there are four directories: / (the outermost directory), a and d (which are in /), and e (which is in a). These directories also contain files of various sizes.

Since the disk is full, your first step should probably be to find directories that are good candidates for deletion. To do this, you need to determine the total size of each directory. The total size of a directory is the sum of the sizes of the files it contains, directly or indirectly. (Directories themselves do not count as having any intrinsic size.)

The total sizes of the directories above can be found as follows:

The total size of directory e is 584 because it contains a single file i of size 584 and no other directories.
The directory a has total size 94853 because it contains files f (size 29116), g (size 2557), and h.lst (size 62596), plus file i indirectly (a contains e which contains i).
Directory d has total size 24933642.
As the outermost directory, / contains every file. Its total size is 48381165, the sum of the size of every file.
To begin, find all of the directories with a total size of at most 100000, then calculate the sum of their total sizes. In the example above, these directories are a and e; the sum of their total sizes is 95437 (94853 + 584). (As in this example, this process can count files more than once!)

Find all of the directories with a total size of at most 100000. What is the sum of the total sizes of those directories?

--- Part Two ---
Now, you're ready to choose a directory to delete.

The total disk space available to the filesystem is 70000000. To run the update, you need unused space of at least 30000000. You need to find a directory you can delete that will free up enough space to run the update.

In the example above, the total size of the outermost directory (and thus the total amount of used space) is 48381165; this means that the size of the unused space must currently be 21618835, which isn't quite the 30000000 required by the update. Therefore, the update still requires a directory with total size of at least 8381165 to be deleted before it can run.

To achieve this, you have the following options:

Delete directory e, which would increase unused space by 584.
Delete directory a, which would increase unused space by 94853.
Delete directory d, which would increase unused space by 24933642.
Delete directory /, which would increase unused space by 48381165.
Directories e and a are both too small; deleting them would not free up enough space. However, directories d and / are both big enough! Between these, choose the smallest: d, increasing unused space by 24933642.

Find the smallest directory that, if deleted, would free up enough space on the filesystem to run the update. What is the total size of that directory?
"""


# load list of commands / output
def import_command_history(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


# find directories below the given size limit and add up the total size
def sum_directories_below_size(directory_sizes, limit):
    total_size = 0
    for _, value in directory_sizes.items():
        if value <= limit:
            total_size += value
    return total_size


# find the smallest directory to delete which meets or exceeds the space needed
def find_smallest_directory_to_delete(directory_sizes, disk_space_needed):
    directory_to_delete = ""
    directory_to_delete_size = 0
    for directory, size in directory_sizes.items():
        if size >= disk_space_needed:
            if directory_to_delete_size == 0:
                directory_to_delete_size = size
                directory_to_delete = directory
            elif size < directory_to_delete_size:
                directory_to_delete_size = size
                directory_to_delete = directory
    return directory_to_delete, directory_to_delete_size


def main():
    directory_structure = {}
    current_directory = "/"
    directory_sizes = {}
    command_history = import_command_history("input.txt")

    disk_size = 70000000
    upgrade_disk_space_needed = 30000000

    for line in command_history:
        line = line.strip()
        line_split = line.split()
        if line_split[0] == "$":
            command = line_split[1]
            if command == "cd":
                target_directory = line_split[2]
                # changing directory
                if target_directory == "..":
                    # using a hash to separate the paths to not muck up root ("/")
                    # split the path on the hash, remove the last entry (current directory), and rejoin with a hash
                    # this effectively moves the current directory down one in the tree
                    current_directory = "#".join(current_directory.split("#")[:-1])
                else:
                    if target_directory != "/":
                        current_directory = f"{current_directory}#{target_directory}"
                    else:
                        current_directory = "/"
        elif line_split[0] == "dir":
            # directory listing
            # save it along with the current directory since we know that is the parent
            directory_name = line_split[1]
            directory_structure[f"{current_directory}#{directory_name}"] = {"parent": current_directory}
        else:
            # at this point we know this is a file
            # parse the size and then add the size to the current directory plus each parent directory
            file_size = int(line_split[0])
            try:
                directory_sizes[current_directory] += file_size
            except:
                directory_sizes[current_directory] = file_size
            temporary_directory = current_directory
            while True:
                if temporary_directory == "/":
                    break
                temporary_directory = "#".join(temporary_directory.split("#")[:-1])
                try:
                    directory_sizes[temporary_directory] += file_size
                except:
                    directory_sizes[temporary_directory] = file_size
    

    part_1_directory_sizes = sum_directories_below_size(directory_sizes, 100000)
    print(f"Sum of directories less-than 100000 in size (part 1): {part_1_directory_sizes}")

    current_free_disk_space = disk_size - directory_sizes["/"]
    disk_space_needed = abs(current_free_disk_space - upgrade_disk_space_needed)
    
    directory_to_delete, directory_to_delete_size = find_smallest_directory_to_delete(directory_sizes, disk_space_needed)
    print(f"Directory to delete for the upgrade (part 2): {directory_to_delete} [{directory_to_delete_size}]")

if __name__ == '__main__':
    main()