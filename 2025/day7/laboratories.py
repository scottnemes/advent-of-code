import sys


def get_input(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


def main():
    # # Part 1
    # data = get_input("input.txt")
  
    # paths = {}
    # carrots = set()
    # splits = 0
    # for y, row in enumerate(data):
    #     row = row[:-1]
    #     for x, v in enumerate(row):
    #         if v == "S":
    #             paths[f"{y}-{x}"] = "X"
    #             continue
    #         if v == "^":
    #             carrots.add(f"{y}-{x}")
    #             paths[f"{y}-{x}"] = "^"
    #             paths[f"{y}-{x-1}"] = "X"
    #             paths[f"{y}-{x+1}"] = "X"

    # for y, row in enumerate(data):
    #     row = row[:-1]
    #     for x, v in enumerate(row):
    #         if v == "^":
    #             for k in sorted(paths.keys(), key=get_y, reverse=True):
    #                 #print(k)
    #                 y2, x2 = k.split("-")
    #                 if int(x2) != x:
    #                     continue
    #                 if int(y2) < y and paths[f"{y2}-{x2}"] == "^":
    #                     break
    #                 if int(y2) < y:
    #                     #print(y,x, y2, x2)
    #                     splits += 1
    #                     break
    # print(splits)

    # Part 2
    data = get_input("sample.txt")
    # 3244 too low
    paths = {}
    carrots = []
    for y, row in enumerate(data):
        row = row[:-1]
        for x, v in enumerate(row):
            if v == "S":
                paths[f"{y}-{x}"] = 1
                continue
            if v == "^":
                carrots.append(f"{y}-{x}")
                paths[f"{y}-{x}"] = "^"
                paths[f"{y}-{x-1}"] = paths.get(f"{y}-{x-1}", 0) + 1
                paths[f"{y}-{x+1}"] = paths.get(f"{y}-{x+1}", 0) + 1

    for y, row in enumerate(data):
        row = row[:-1]
        for x, v in enumerate(row):
            if v == "^":
                for k in sorted(paths.keys(), key=get_y, reverse=True):
                    #print(k)
                    y2, x2 = k.split("-")
                    if int(x2) != x:
                        continue
                    if int(y2) < y and paths[f"{y2}-{x2}"] == "^":
                        break
                    if int(y2) < y:
                        break
    count = 0
    for k, v in paths.items():
        if v != "^":
            count += int(v)
    print(count)


def get_y(cords: str) -> int:
    return int(cords.split("-")[0])


if __name__ == '__main__':
    main()
