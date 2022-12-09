# head and tail must always be touching
# if head is two steps directly up, down, left, or right from tail, tail must move one step in that direction
# if head and tail aren't touching and aren't in the same row or column, tail moves one step diagonally
# assume head and tail start in same location


# load in rope movements
def import_rope_movements(file):
    lines = []
    with open(file) as f:
        for line in f.readlines():
            lines.append(line.strip())
    return lines


# move the tail based on the position of the head
# the head may be another tail, but it functions the same
def move_rope_tail(head_x, head_y, tail_x, tail_y, tail_locations, actual_tail):
    x_diff = head_x - tail_x
    y_diff = head_y - tail_y
    if x_diff > 1 and y_diff == 0:
        # tail is 2+ spaces to the left of the head, need to move right one space
        tail_x += 1
    elif x_diff < -1 and y_diff == 0:
        # tail is 2+ spaces to the right of the head, need to move left one space
        tail_x -= 1   
    elif x_diff == 0 and y_diff > 1:
        # tail is 2+ spaces below the head, need to move up one space
        tail_y += 1
    elif x_diff == 0 and y_diff < -1:
        # tail is 2+ spaces above the head, need to move down one space
        tail_y -= 1
    elif x_diff == 1 and y_diff > 1:
        # tail is 2+ spaces below and 1 space to the left of the head, need to move diagonally up 1 space and to the right 1 space
        tail_x += 1
        tail_y += 1
    elif x_diff > 1 and y_diff >= 1:
        # tail is 2+ spaces below and 1+ spaces to the left of the head, need to move diagonally down 1 space and to the left 1 space
        tail_x += 1
        tail_y += 1
    elif x_diff == -1 and y_diff < -1:
        # tail is 2+ spaces above and 1 space to the right of the head, need to move diagonally down 1 space and to the left 1 space
        tail_x -= 1
        tail_y -= 1
    elif x_diff < -1 and y_diff <= -1:
        # tail is 1+ spaces above and 2+ spaces to the right of the head, need to move diagonally down 1 space and to the left 1 space
        tail_x -= 1
        tail_y -= 1
    elif x_diff == -1 and y_diff > 1:
        # tail is 2+ spaces below and 1 space to the right of the head, need to move diagonally up 1 space and to the left 1 space
        tail_x -= 1
        tail_y += 1
    elif x_diff < -1 and y_diff == 1:
        # tail is 1 space below and 2+ spaces to the right of the head, need to move diagonally up 1 space and to the left 1 space
        tail_x -= 1
        tail_y += 1
    elif x_diff == 1 and y_diff < -1:
        # tail is 2+ spaces above and 1 space to the left of the head, need to move diagonally down 1 space and to the right 1 space
        tail_x += 1
        tail_y -= 1
    elif x_diff > 1 and y_diff == -1:
        # tail is 1 space above and 2+ spaces to the left of the head, need to move diagonally down 1 space and to the right 1 space
        tail_x += 1
        tail_y -= 1
    elif x_diff < -1 and y_diff > 1:
        # tail is 2+ spaces below and 2+ spaces to the right of the head, need to move diagonally up 1 space and to the left 1 space
        tail_x -= 1
        tail_y += 1
    elif x_diff > 1 and y_diff < -1:
        # tail is 2+ spaces above and 2+ spaces to the left of the head, need to move diagonally down 1 space and to the right 1 space
        tail_x += 1
        tail_y -= 1
    if actual_tail:
        tail_locations.add(f"{tail_x},{tail_y}")
    return head_x, head_y, tail_x, tail_y, tail_locations


# move the head and then move the (first/only) tal after
def move_rope_head(head_x, head_y, tail_x, tail_y, tail_locations, direction, actual_tail = True):
    grid = [["O" for i in range(6)] for j in range(5)]
    if direction == "U":
        head_y += 1
        head_x, head_y, tail_x, tail_y, tail_locations = move_rope_tail(head_x, head_y, tail_x, tail_y, tail_locations, actual_tail)
    elif direction == "D":
        head_y -= 1
        head_x, head_y, tail_x, tail_y, tail_locations = move_rope_tail(head_x, head_y, tail_x, tail_y, tail_locations, actual_tail)
    elif direction == "L":
        head_x -= 1
        head_x, head_y, tail_x, tail_y, tail_locations = move_rope_tail(head_x, head_y, tail_x, tail_y, tail_locations, actual_tail)
    elif direction == "R":
        head_x += 1
        head_x, head_y, tail_x, tail_y, tail_locations = move_rope_tail(head_x, head_y, tail_x, tail_y, tail_locations, actual_tail)
    else:
        print("Critical failure, party wipe.")
    return head_x, head_y, tail_x, tail_y, tail_locations


# process the rope movements when there is only one tail
# in this case we only need to make a single call to the move_rope_head() function per move
# this will move the head and the single tail
def process_rope_movements_one_tail(rope_movements):
    head_x = 0
    head_y = 0
    tail_x = 0
    tail_y = 0
    tail_locations = set()
    for move in rope_movements:
        move_split = move.split()
        direction = move_split[0]
        count = int(move_split[1])
        for _ in range(count):
            head_x, head_y, tail_x, tail_y, tail_locations = move_rope_head(head_x, head_y, tail_x, tail_y, tail_locations, direction, True)
    return tail_locations


# process the rope movements when there are multiple tails (9 in this specific case)
# we call move_rope_head() function once to move the head and first tail
# and then make subsequent calls directly to the move_rope_tail() function to move the each additional tail
# starting with the 2nd tail, the previous tail serves as the head for the tail to follow
def process_rope_movements_multiple_tails(rope_movements):
    head = [0, 0]
    # tails = [starting_x, starting_y, actual_tail]
    # actual_tail is used so we only track the locations of the very last tail
    tails = [
        [0,0, False],
        [0,0, False],
        [0,0, False],
        [0,0, False],
        [0,0, False],
        [0,0, False],
        [0,0, False],
        [0,0, False],
        [0,0, True]
    ]
    tail_locations = set()

    for move in rope_movements:
        move_split = move.split()
        direction = move_split[0]
        count = int(move_split[1])
        for _ in range(count):
            head[0], head[1], tails[0][0], tails[0][1], tail_locations = move_rope_head(head[0], head[1], tails[0][0], tails[0][1], tail_locations, direction, False)
            # start at tail1 because tail0 tracks the head above
            for i in range(1,9,1):
                tails[i-1][0], tails[i-1][1], tails[i][0], tails[i][1], tail_locations = move_rope_tail(tails[i-1][0], tails[i-1][1], tails[i][0], tails[i][1], tail_locations, tails[i][2])
    return tail_locations


def main():
    rope_movements = import_rope_movements("input.txt")
   
    # part1
    tail_locations = process_rope_movements_one_tail(rope_movements)
    print(f"Number of tail locations visited at least once (part 1): {len(tail_locations)}")

    tail_locations = process_rope_movements_multiple_tails(rope_movements)
    print(f"Number of tail locations visited at least once (part 2): {len(tail_locations)}")

if __name__ == '__main__':
    main()