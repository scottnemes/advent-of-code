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
    # 837905266357201192774153975965072960276319014627339281004922249271022419931806730707108158287237275 too high
    # 147293551176578 too low
    batteries = get_input("sample.txt")
    total = 0
    batteries_needed = 12
    removable = len(batteries[0]) - 1 - batteries_needed
    for bank in batteries:
        bank = bank[:-1]
        batteries_remaining = len(bank)
        joltage = ""
        max_removed = False
        #print(bank)
        for i, battery in enumerate(bank):
            print(f"joltage: {joltage}, batteries_remaining: {batteries_remaining}, left: {batteries_remaining + len(joltage)}")
            if len(joltage) == batteries_needed:
                break
            if max_removed or batteries_remaining + len(joltage) <= batteries_needed:
                joltage += bank[i]
                batteries_remaining -= 1
                max_removed = True
                continue
            if bank[i] >= bank[i+1]:
                joltage += bank[i]
            batteries_remaining -= 1
        total += int(joltage)
        print(joltage)
    print(total)


if __name__ == '__main__':
    main()
