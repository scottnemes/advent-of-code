"""
--- Day 5: Supply Stacks ---

--- Part One ---
The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.
The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.
The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.
They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 
In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3
Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3
Finally, one crate is moved from stack 1 to stack 2:

        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3
The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.
After the rearrangement procedure completes, what crate ends up on top of each stack?
"""

"""
[N]     [Q]         [N]            
[R]     [F] [Q]     [G] [M]        
[J]     [Z] [T]     [R] [H] [J]    
[T] [H] [G] [R]     [B] [N] [T]    
[Z] [J] [J] [G] [F] [Z] [S] [M]    
[B] [N] [N] [N] [Q] [W] [L] [Q] [S]
[D] [S] [R] [V] [T] [C] [C] [N] [G]
[F] [R] [C] [F] [L] [Q] [F] [D] [P]
 1   2   3   4   5   6   7   8   9 


--- Part Two ---
As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.

Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

The CrateMover 9001 is notable for many new and exciting features: air conditioning, leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.

Again considering the example above, the crates begin in the same configuration:

    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 
Moving a single crate from stack 2 to stack 1 behaves the same as before:

[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 
However, the action of moving three crates from stack 1 to stack 3 means that those three moved crates stay in the same order, resulting in this new configuration:

        [D]
        [N]
    [C] [Z]
    [M] [P]
 1   2   3
Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

        [D]
        [N]
[C]     [Z]
[M]     [P]
 1   2   3
Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

        [D]
        [N]
        [Z]
[M] [C] [P]
 1   2   3
In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.

Before the rearrangement process finishes, update your simulation so that the Elves know where they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?
"""


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