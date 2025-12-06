import sys


def get_input(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


def main():
    input = get_input("input.txt")

    # # Part 1
    # problems = []
    # for row in input:
    #     row = row[:-1]
    #     problem = row.split()
    #     problems.append(problem)

    # total = 0
    # for row in range(len(problems[0])):
    #     nums = []
    #     for i in range(len(problems) - 1): # range is up to, plus the -1, so -2 in reality
    #         nums.append(problems[i][row])
    #     problem = f"{problems[-1][row]}".join(nums)
    #     total += eval(problem)
    # print(total)

    # Part 2
    
    parsed_problems = []
    problems = []
    for row in input:
        row = row[:-1]
        problems.append(list(row))
        parsed_problems.append(row.split())
   
    start_idxs = []
    for i, op in enumerate(problems[-1]):
        if op == " ":
            continue
        else:
            start_idxs.append(i)

    max_num_lengths = []
    for col in range(len(parsed_problems[0])):
        max_length = 0
        for row in range(len(parsed_problems)-1):
            length = len(parsed_problems[row][col])
            if length > max_length:
                max_length = length
        max_num_lengths.append(max_length)

    total = 0
    for i, idx in enumerate(start_idxs):
        problem = []
        for j, row in enumerate(problems[:-1]):
            num = ""
            for char in row[idx:idx + max_num_lengths[i]]:
                if char == " ":
                    char = "."
                num += str(char)
            problem.append(num)

        new_problem = []
        for k in range(len(problem[0])):
            num = ""
            for t in range(len(problem)):
                if problem[t][k] == ".":
                    continue
                num += problem[t][k]
            new_problem.append(num)
        
        total += eval(f"{parsed_problems[-1][i]}".join(new_problem))
    print(total)


if __name__ == '__main__':
    main()
