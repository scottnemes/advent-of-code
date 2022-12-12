def import_monkey_data(file):
    lines = []
    with open(file) as f:
        for line in f.readlines():
            if line.strip():
                lines.append(line.strip())
    return lines


# After each monkey inspects an item but before it tests your worry level, divide worry by 3 and round down to int
# On a single monkey's turn, it inspects and throws all of the items it is holding one at a time and in the order listed
# When a monkey throws an item to another monkey, the item goes on the end of the recipient monkey's list.
# If a monkey is holding no items at the start of its turn, its turn ends.

# Count the total number of times each monkey inspects items over 20 rounds:


# Monkey 0:
#   Starting items: 79, 98
#   Operation: new = old * 19
#   Test: divisible by 23
#     If true: throw to monkey 2
#     If false: throw to monkey 3


class Monkey:
    def __init__(self, number):
        self.number = number
        self.items = []
        self.operation = []
        self.test = []
        self.test_true_action = []
        self.test_false_action = []
        self.items_inspected = 0
    def print(self):
        print(f"number: {self.number}, items: {self.items}, operation: {self.operation}, test: {self.test}, test_true_action: {self.test_true_action}, test_false_action: {self.test_false_action}")


def create_monkeys(monkey_data):
    monkeys = []
    mod = 1
    for row in monkey_data:
        if row.split()[0] == "Monkey":
            monkey_number = int(row.split()[1][:-1]) # get number without the ending :
            monkeys.append(Monkey(monkey_number))
    for row in monkey_data:
        if row.split()[0] == "Monkey":
            monkey_number = int(row.split()[1][:-1]) # get number without the ending :
        elif row.split(":")[0] == "Starting items":
            for item in row.split(":")[1].split(","):
                monkeys[monkey_number].items.append(item.strip())
        elif row.split(":")[0] == "Operation":
            monkeys[monkey_number].operation = row.split("=")[1].strip().split()
        elif row.split()[0] == "Test:":
            test = row.split(":")[1].strip().split()
            monkeys[monkey_number].test = [test[0], test[2]]
            mod *= int(test[2])
        elif row.split(":")[0].strip() == "If true":
            target = row.split(":")[1].split()
            monkeys[monkey_number].test_true_action = [target[0],target[3]]
        elif row.split(":")[0].strip() == "If false":
            target = row.split(":")[1].split()
            monkeys[monkey_number].test_false_action = [target[0],target[3]]
    return monkeys, mod


def process_monkey_round(monkeys, monkey_number, mod=1, do_not_worry=True):
    # inspect item
    # perform operation
    # divide by 3
    # perform test
    monkey = monkeys[monkey_number]
    for item in monkey.items:
        if do_not_worry:
            worry_level = int(item)
        else:
            worry_level = int(item) % mod
        operation = monkey.operation
        operation_expression = []
        for step in operation:
            if step == "old":
                step = worry_level
            operation_expression.append(str(step))
        # join expression list into a string, eval it, and then floor divide by 3
        if do_not_worry:
            worry_level = eval(" ".join(operation_expression)) // 3
        else:
            worry_level = eval(" ".join(operation_expression))
        # remove item from the current monkey's item list
        test = monkey.test
        if test[0] == "divisible":
            if do_not_worry:
                result = int(worry_level) % int(test[1])
            else:
                result = int(worry_level) % int(test[1])
        else:
            print("Unknown test")
        if result == 0:
            if monkey.test_true_action[0] == "throw":
                monkeys[int(monkey.test_true_action[1])].items.append(worry_level)
        else:
            if monkey.test_true_action[0] == "throw":
                monkeys[int(monkey.test_false_action[1])].items.append(worry_level)
        monkey.items_inspected += 1
    monkeys[monkey_number].items = []
    return monkeys


def main():
    monkey_data = import_monkey_data("input.txt")

    # part 1
    monkeys, _ = create_monkeys(monkey_data)
    rounds = 20
    for _ in range(rounds):
        for monkey in monkeys:
            monkeys = process_monkey_round(monkeys, monkey.number)  
   
    items_inspected = []
    for monkey in monkeys:
        items_inspected.append(monkey.items_inspected)
    items_inspected.sort()
    monkey_business = items_inspected[len(monkeys)-2] * items_inspected[len(monkeys)-1]
    print(f"Monkey business (part 1): {monkey_business}")

    # part 2
    monkeys, mod = create_monkeys(monkey_data)
    rounds = 10000
    for round in range(rounds):
        for monkey in monkeys:
            monkeys = process_monkey_round(monkeys, monkey.number, mod, False)
   
    items_inspected = []
    for monkey in monkeys:
        items_inspected.append(monkey.items_inspected)
    items_inspected.sort()
    monkey_business = items_inspected[len(monkeys)-2] * items_inspected[len(monkeys)-1]
    print(f"Monkey business (part 2): {monkey_business}")


if __name__ == '__main__':
    main()