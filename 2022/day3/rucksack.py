import string


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