"""
--- Day 3: Rucksack Reorganization ---

--- Part One ---
One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.
Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.
The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer to different types of items).
The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

For example, suppose you have the following list of contents from six rucksacks:
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
The fourth rucksack's compartments only share item type v.
The fifth rucksack's compartments only share item type t.
The sixth rucksack's compartments only share item type s.

To help prioritize item rearrangement, every item type can be converted to a priority:
Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.

In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?

--- Part Two ---
As you finish identifying the misplaced items, the Elves come to you with another issue.
For safety, the Elves are divided into groups of three. Every Elf carries a badge that identifies their group. For efficiency, within each group of three Elves, the badge is the only item type carried by all three Elves. That is, if a group's badge is item type B, then all three Elves will have item type B somewhere in their rucksack, and at most two of the Elves will be carrying any other item type.
The problem is that someone forgot to put this year's updated authenticity sticker on the badges. All of the badges need to be pulled out of the rucksacks so the new authenticity stickers can be attached.
Additionally, nobody wrote down which item type corresponds to each group's badges. The only way to tell which item type is the right one is by finding the one item type that is common between all three Elves in each group.

Every set of three lines in your list corresponds to a single group, but each group can have a different badge item type. So, in the above example, the first group's rucksacks are the first three lines:
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg

And the second group's rucksacks are the next three lines:
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

In the first group, the only item type that appears in all three rucksacks is lowercase r; this must be their badges. In the second group, their badge item type must be Z.
Priorities for these items must still be found to organize the sticker attachment efforts: here, they are 18 (r) for the first group and 52 (Z) for the second group. The sum of these is 70.

Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the priorities of those item types?
"""


import string
import sys


# load list of rucksacks / items
def import_rucksacks(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


# split rucksack items into comparments
# return list of items that are in both compartments
def find_duplicate_items(rucksack):
    duplicate_items = []
    middle = int(len(rucksack) / 2)
    compartment_1 = rucksack[:middle]
    compartment_2 = rucksack[middle:]
    for item in compartment_1:
        if item in compartment_2:
            if item not in duplicate_items:
                duplicate_items.append(item)
    return duplicate_items


# return item that is in the three given rucksacks
def find_group_badge(rucksacks):
    rucksack_1 = rucksacks[0]
    rucksack_2 = rucksacks[1]
    rucksack_3 = rucksacks[2]
    for item in rucksack_1:
        if item in rucksack_2 and item in rucksack_3:
            return item


# calculate the priority of the given item(s)
def calculate_priorities(items):
    total = 0
    for item in items:
        total += (string.ascii_letters.index(item) + 1)
    return total


def main():
    duplicate_items_priority_total = 0
    rucksacks = import_rucksacks("input.txt")

    # find items that are in both comparments of a given rucksack
    for rucksack in rucksacks:
        rucksack = rucksack.strip()
        duplicate_items = find_duplicate_items(rucksack)
        duplicate_items_priority_total += calculate_priorities(duplicate_items)
    
    print(f"Total priority of duplicate items (part 1): {duplicate_items_priority_total}")

    i = 0
    badge_priority_total = 0
    group_rucksacks = []
    # group rucksacks into groups of three
    # find item that is in all three rucksacks
    for rucksack in rucksacks:
        rucksack = rucksack.strip()
        if i > 2:
            i = 0
            badge = find_group_badge(group_rucksacks)
            badge_priority_total += calculate_priorities(badge)
            group_rucksacks = []
        group_rucksacks.append(rucksack)
        i += 1
    # process the last group
    badge = find_group_badge(group_rucksacks)
    badge_priority_total += calculate_priorities(badge)
    
    print(f"Total priorty of group badges (part 2): {badge_priority_total}")


if __name__ == '__main__':
    main()