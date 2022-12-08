"""
--- Day 1: Calorie Counting ---
Part 1
Santa's reindeer typically eat regular reindeer food, but they need a lot of magical energy to deliver presents on Christmas. For that, their favorite snack is a special type of star fruit that only grows deep in the jungle. The Elves have brought you on their annual expedition to the grove where the fruit grows.
To supply enough magical energy, the expedition needs to retrieve a minimum of fifty stars by December 25th. Although the Elves assure you that the grove has plenty of fruit, you decide to grab any fruit you see along the way, just in case.
Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!
The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).
The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

Part 2
By the time you calculate the answer to the Elves' question, they've already realized that the Elf carrying the most Calories of food might eventually run out of snacks.
To avoid this unacceptable situation, the Elves would instead like to know the total Calories carried by the top three Elves carrying the most Calories. That way, even if one of those Elves runs out of snacks, they still have two backups.
In the example above, the top three Elves are the fourth Elf (with 24000 Calories), then the third Elf (with 11000 Calories), then the fifth Elf (with 10000 Calories). The sum of the Calories carried by these three elves is 45000.

Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
"""

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