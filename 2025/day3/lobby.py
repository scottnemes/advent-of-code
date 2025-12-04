import sys


def get_input(file):
    with open(file) as f:
        lines = f.readlines()
    return lines
  

def main():
    # Part 1
    # batteries = get_input("input.txt")
    # total = 0
    # for bank in batteries:
    #     first = bank[0]
    #     first_idx = 0
    #     second = "0"
    #     for i, battery in enumerate(bank[1:-2]):
    #         if battery > first:
    #             first = battery
    #             first_idx = i + 1
    #     for battery in bank[first_idx+1:]:
    #         if battery > second:
    #             second = battery
    #     joltage = int(first + second)
    #     total += joltage
    # print(total)

    # Part 2
    batteries = get_input("input.txt")
    total = 0
    batteries_needed = 12
    for bank in batteries:
        bank = bank[:-1]
        batteries_remaining = len(bank)
        joltage = ""
        max_removed = False
        for i, _ in enumerate(bank):
            if len(joltage) == batteries_needed:
                break
            if max_removed or batteries_remaining + len(joltage) <= batteries_needed:
                joltage += bank[i]
                batteries_remaining -= 1
                max_removed = True
                continue
            skipped = False
            for bat in bank[i+1:len(bank)-batteries_needed+len(joltage)+1]:
                if bat > bank[i]:
                    skipped = True
                    break
            if not skipped:
                joltage += bank[i]
            batteries_remaining -= 1
        total += int(joltage)
    print(total)


if __name__ == '__main__':
    main()
