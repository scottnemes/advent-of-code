import sys


def get_input(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


def main():
    input = get_input("input.txt")

    # # Part 1  
    # ranges = {}
    # ids = []

    # hit_divider = False
    # for row in input:
    #     row = row[:-1]
    #     if row == "":
    #         hit_divider = True
    #         continue
    #     if not hit_divider:
    #         bounds = row.split("-")
    #         start = int(bounds[0])
    #         end = int(bounds[1])
    #         if start in ranges.keys():
    #             if ranges[start] < end:
    #                 ranges[start] = end
    #                 continue
    #         else:
    #             ranges[start] = end
    #     else:
    #         ids.append(int(row))

    # fresh = 0
    # for i in ids:
    #     for k, v in ranges.items():
    #         if i >= k and i <= v:
    #             fresh += 1
    #             break
    # print(len(fresh))

    # Part 2
    ranges = []

    for row in input:
        row = row[:-1]
        if row == "":
            break
        bounds = row.split("-")
        start = int(bounds[0])
        end = int(bounds[1])
        ranges.append([start, end])
        ranges.sort(key=lambda x: x[0])

    merged_ranges = []
    current_start = ranges[0][0]
    current_end = ranges[0][1]

    for start, end in ranges[1:]:
        # print(f"current_start: {current_start}")
        # print(f"current_end: {current_end}")
        # print(f"start: {start}")
        # print(f"end: {end}")
        if start <= current_end:
            # print("start <= currend_end")
            current_end = max(current_end, end)
        else:
            merged_ranges.append([current_start, current_end])
            current_start = start
            current_end = end
    
    merged_ranges.append([current_start, current_end])

    ids = 0
    for start, end in merged_ranges:
        ids += end - start + 1
    
    print(ids)


if __name__ == '__main__':
    main()
