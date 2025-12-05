import sys


def get_input(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


def check_space(x: int, y: int) -> bool:
    count = 0
    if x-1 >= 0 and grid[y][x-1] == "@":
        count += 1
    if x+1 < len(grid[0]) and grid[y][x+1] == "@":
        count += 1
    if y-1 >= 0 and grid[y-1][x] == "@":
        count += 1
    if y+1 < len(grid) and grid[y+1][x] == "@":
        count += 1
    if y-1 >= 0 and x-1 >= 0 and grid[y-1][x-1] == "@":
        count += 1
    if y-1 >= 0 and x+1 < len(grid[0]) and grid[y-1][x+1] == "@":
        count += 1
    if y+1 < len(grid) and x-1 >= 0 and grid[y+1][x-1] == "@":
        count += 1
    if y+1 < len(grid) and x+1 < len(grid[0]) and grid[y+1][x+1] == "@":
        count += 1
    if count < 4:
        return True 


grid = []


def main():
    global grid

    # # Part 1
    # input = get_input("sample.txt")
    # for row in input:
    #     rows = []
    #     for col in row[:-1]:
    #         rows.append(col)
    #     grid.append(rows)
    
    # rolls = 0
    # for y in range(len(grid)):
    #     for x in range(len(grid[0])):
    #         count = 0
    #         if grid[y][x] != "@":
    #             continue
    #         valid = check_space(x, y)
    #         if not valid:
    #             continue
    #         rolls += 1
    # print(rolls)

    # Part 2
    input = get_input("input.txt")
    global grid
    for row in input:
        rows = []
        for col in row[:-1]:
            rows.append(col)
        grid.append(rows)
    
    rolls = 0
    retries = []
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if grid[y][x] != "@":
                continue
            valid = check_space(x, y)
            if not valid:
                continue
            rolls += 1
            grid[y][x] = "."
            retries.append([y-1,x])
            retries.append([y+1,x])
            retries.append([y,x-1])
            retries.append([y,x+1])
            retries.append([y-1,x-1])
            retries.append([y-1,x+1])
            retries.append([y+1,x-1])
            retries.append([y+1,x+1])

    while len(retries) > 0:
        old_retries = retries
        retries = []
        for coords in old_retries:
            y = coords[0]
            x = coords[1]
            if x < 0 or x >= len(grid[0]) or y < 0 or y >= len(grid):
                continue
            if grid[y][x] != "@":
                continue
            valid = check_space(x, y)
            if not valid:
                continue
            rolls += 1
            grid[y][x] = "."
            retries.append([y-1,x])
            retries.append([y+1,x])
            retries.append([y,x-1])
            retries.append([y,x+1])
            retries.append([y-1,x-1])
            retries.append([y-1,x+1])
            retries.append([y+1,x-1])
            retries.append([y+1,x+1])

    print(rolls)


if __name__ == '__main__':
    main()
