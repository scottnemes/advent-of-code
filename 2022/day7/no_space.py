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