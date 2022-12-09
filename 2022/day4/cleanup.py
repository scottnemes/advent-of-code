# load list of cleaning schedules / ids
def import_cleaning_ids(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


# parse ranges from the schedule of ids
def get_cleaning_ranges(cleaning_ids):
    ranges = cleaning_ids.split(",")
    range_1 = ranges[0].split("-")
    range_1_start = int(range_1[0])
    range_1_end = int(range_1[1])
    range_2 = ranges[1].split("-")
    range_2_start = int(range_2[0])
    range_2_end = int(range_2[1])
    return range_1_start, range_1_end, range_2_start, range_2_end


# check if the two schedules fully overlap
def ranges_fully_overlap(cleaning_ids):
    range_1_start, range_1_end, range_2_start, range_2_end = get_cleaning_ranges(cleaning_ids)
    if range_1_start >= range_2_start and range_1_end <= range_2_end:
        return True
    elif range_2_start >= range_1_start and range_2_end <= range_1_end:
        return True
    else:
        return False


# check if the two schedules partially overlap
def ranges_partially_overlap(cleaning_ids):
    range_1_start, range_1_end, range_2_start, range_2_end = get_cleaning_ranges(cleaning_ids)
    if range_1_start >= range_2_start and range_1_start <= range_2_end:
        return True
    elif range_1_end >= range_2_start and range_1_end <= range_2_end:
        return True
    elif range_2_start >= range_1_start and range_2_start <= range_1_end:
        return True
    elif range_2_end >= range_1_start and range_2_end <= range_1_end:
        return True
    else:
        return False


def main():
    cleaning_schedule = import_cleaning_ids("input.txt")
    fully_overlap_count = 0
    partially_overlap_count = 0
    for cleaning_ids in cleaning_schedule:
        cleaning_ids = cleaning_ids.strip()
        if ranges_fully_overlap(cleaning_ids):
            fully_overlap_count += 1
        if ranges_partially_overlap(cleaning_ids):
            partially_overlap_count += 1
    print(f"Number of fully overlapping cleaning schedules: {fully_overlap_count}")
    print(f"Number of partially overlapping cleaning schedules: {partially_overlap_count}")


if __name__ == '__main__':
    main()