import sys


# load list of calories written by elves
def import_elf_calories(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


# sum up calories by elf
# a blank line signals the start of a new elf's calories
def process_elf_calories(calories_list):
    elf_calories = []
    calories = 0

    if len(calories_list) == 0:
        sys.exit(1)

    for calorie in calories_list:
        calorie = calorie.strip()
        if not calorie:
            elf_calories.append(calories)
            calories = 0
            continue
        calories += int(calorie.strip())
    return elf_calories


# return the sum of calories for the top N elves
def add_elf_calories(calories_list, top_number):
    calories_list.sort(reverse=True)
    calories = 0
    for i in range(top_number):
        calories += calories_list[i]
    return calories
    

def main():
    calories_input = import_elf_calories("input.txt")
    calories_list = process_elf_calories(calories_input)
    top_calories_1 = add_elf_calories(calories_list, 1)
    top_calories_3 = add_elf_calories(calories_list, 3)
    print(f"Top Calories for 1 Elves: {top_calories_1}")
    print(f"Top Calories for 3 Elves: {top_calories_3}")


if __name__ == '__main__':
    main()