import sys


def get_input(file):
    with open(file) as f:
        lines = f.readlines()
    return lines
  

def main():
    safe_actions = get_input("input.txt")
    position = 50
    count = 0
    min = 0
    max = 99
    for action in safe_actions:
        direction = action[0]
        amount = int(action[1:]) % (max + 1)
        #print(f"dir: {direction}, amount: {amount}, pos: {position}")
        if direction == "L":
            position -= amount
            #print(f"first move pos: {position}, amount: {amount}")
            if position < min:
                over = abs(position)
                position = max - over + 1
        elif direction == "R":
            position += amount
            if position >= max:
                over = abs(position - max)
                position = min + over - 1
        else:
            print("wrong")
        #print(f"pos: {position}")
        if position == 0:
            count += 1
    print(f"Part 1 Count: {count}")

    position = 50
    count = 0

    for action in safe_actions:
        direction = action[0]
        amount = int(action[1:]) % (max + 1)
        count += int(action[1:]) // (max + 1)
        if direction == "L":
            position -= amount
            if position < min:
                over = abs(position)
                position = max - over + 1
                count += 1
        elif direction == "R":
            position += amount
            if position >= max:
                over = abs(position - max)
                position = min + over - 1
                count += 1
        else:
            print("wrong")
    print(f"Part 2 Count: {count}")



if __name__ == '__main__':
    main()
