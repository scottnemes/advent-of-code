# load list of moves / desired outcomes
def import_strategy_guide(file):
    with open(file) as f:
        lines = f.readlines()
    return lines


# play a round by part 1 rules, where value 2 is the move you play
def play_rock_paper_scissors_part_1(your_move, their_move):
    # Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
    # score = 
    # shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
    # +
    # score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
    move_factor = {
        "rock": 1,
        "paper": 2,
        "scissors": 3
    }
    if your_move == their_move:
        win_factor = 3
    elif your_move == "rock" and their_move == "scissors":
        win_factor = 6
    elif your_move == "scissors" and their_move == "paper":
        win_factor = 6
    elif your_move == "paper" and their_move == "rock":
        win_factor = 6
    else:
        win_factor = 0
    return win_factor + move_factor[your_move]


# play a round by part 2 rules, where value 2 is the desired outcome
def play_rock_paper_scissors_part_2(desired_result, their_move):
    # Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
    # score = 
    # shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
    # +
    # score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
    move_factor = {
        "rock": 1,
        "paper": 2,
        "scissors": 3
    }
    if desired_result == "draw":
        win_factor = 3
        your_move = their_move
    elif desired_result == "win":
        win_factor = 6
        if their_move  == "scissors":
            your_move = "rock"
        if their_move == "paper":
            your_move = "scissors"
        if their_move == "rock":
            your_move = "paper"
    elif desired_result == "lose":
        win_factor = 0
        if their_move  == "scissors":
            your_move = "paper"
        if their_move == "paper":
            your_move = "rock"
        if their_move == "rock":
            your_move = "scissors"
    return win_factor + move_factor[your_move]


def main():
    moves = {
        "A": "rock",
        "X": "rock",
        "B": "paper",
        "Y": "paper",
        "C": "scissors",
        "Z": "scissors"
    }
    results = {
        "X": "lose",
        "Y": "draw",
        "Z": "win"
    }

    total_score_part_1 = 0
    total_score_part_2 = 0

    strategy_guide = import_strategy_guide("input.txt")
    # part 1
    for round in strategy_guide:
        round = round.strip()
        their_move = moves[round[0]]
        your_move = moves[round[2]]
        total_score_part_1 += play_rock_paper_scissors_part_1(your_move, their_move)
    
    # part 2
    for round in strategy_guide:
        round = round.strip()
        their_move = moves[round[0]]
        desired_result = results[round[2]]
        total_score_part_2 += play_rock_paper_scissors_part_2(desired_result, their_move)
    
    print(f"Total score (part 1): {total_score_part_1}")
    print(f"Total score (part 2): {total_score_part_2}")


if __name__ == '__main__':
    main()