# load list of crate moves
def import_moves(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


# put moves into an easier to parse list format
def format_moves(moves):
    # move 3 from 9 to 4
    formatted_moves = []
    for move in moves:
        move = move.strip()
        split_move = move.split(" ")
        formatted_moves.append([int(split_move[1]),int(split_move[3]),int(split_move[5])])
    return formatted_moves


# move crates one by one
def move_crate_9000(crate_piles, source, destination):
    # decrement by one to make indexing easier
    source -= 1
    destination -= 1
    crate_piles[destination].append(crate_piles[source][-1])
    crate_piles[source] = crate_piles[source][:-1]


# move crates multiple at a time, keeping the existing order
def move_crate_9001(crate_piles, count, source, destination):
    # decrement by one to make indexing easier
    source -= 1
    destination -= 1
    for crate in crate_piles[source][count*-1:]:
        crate_piles[destination].append(crate)
    crate_piles[source] = crate_piles[source][:count*-1]


def main():
    crate_piles_9000 = [
        ["F", "D", "B", "Z", "T", "J", "R", "N"],
        ["R", "S", "N", "J", "H"],
        ["C", "R", "N", "J", "G", "Z", "F", "Q"],
        ["F", "V", "N", "G", "R", "T", "Q"],
        ["L", "T", "Q", "F"],
        ["Q", "C", "W", "Z", "B", "R", "G", "N"],
        ["F", "C", "L", "S", "N", "H", "M"],
        ["D", "N", "Q", "M", "T", "J"],
        ["P", "G", "S"]
    ]
    crate_piles_9001 = [
        ["F", "D", "B", "Z", "T", "J", "R", "N"],
        ["R", "S", "N", "J", "H"],
        ["C", "R", "N", "J", "G", "Z", "F", "Q"],
        ["F", "V", "N", "G", "R", "T", "Q"],
        ["L", "T", "Q", "F"],
        ["Q", "C", "W", "Z", "B", "R", "G", "N"],
        ["F", "C", "L", "S", "N", "H", "M"],
        ["D", "N", "Q", "M", "T", "J"],
        ["P", "G", "S"]
    ]    

    moves = import_moves("input.txt")
    formatted_moves = format_moves(moves)
    # for part 1, we move crates one by one
    # so loop through the move from source to dest N times
    for move in formatted_moves:
        for i in range(move[0]):
            move_crate_9000(crate_piles_9000, move[1], move[2])

    # get the last item (crate) in the list to find the top create
    top_crates_9000 = ""
    for crate_stack in crate_piles_9000:
        top_crates_9000 = f"{top_crates_9000}{crate_stack[-1]}"
    
    print(f"Top crate for each stack (CrateMover 9000 - part1): {top_crates_9000}")

    # for part 2 we can move multiple crates at once, so make a single call per move
    for move in formatted_moves:
        move_crate_9001(crate_piles_9001, move[0], move[1], move[2])

    # get the last item (crate) in the list to find the top create
    top_crates_9001 = ""
    for crate_stack in crate_piles_9001:
        top_crates_9001 = f"{top_crates_9001}{crate_stack[-1]}"
    
    print(f"Top crate for each stack (CrateMover 9001 - part2): {top_crates_9001}")


if __name__ == '__main__':
    main()