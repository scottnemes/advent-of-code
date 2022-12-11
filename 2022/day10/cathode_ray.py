# addx V takes two cycles to complete. After two cycles, the X register is increased by the value V. (V can be negative.)
# noop takes one cycle to complete. It has no other effect.


# load list of cpu instructions
def import_cpu_instructions(file):
    lines = []
    with open(file) as f:
        for line in f.readlines():
            lines.append(line.strip())
    return lines


# build out a list of cycles with each instruction in the right cycle
# leaves empty lists in the cycles where no instructions are happening
def create_program(cpu_instructions, total_cycles):
    cpu_instructions_by_cycle = [[] for i in range(total_cycles)]
    cycle = 0
    for instruction in cpu_instructions:
        instruction_split = instruction.split()
        command = instruction_split[0]
        if command != "noop":
            value = instruction_split[1]
            cycle += 2
            cpu_instructions_by_cycle[cycle] = [command, value]
        else:
            cycle += 1
            cpu_instructions_by_cycle[cycle] = [command, 0]
    return cpu_instructions_by_cycle


# update the given pixel in the correct row if it syncs up with the cycle
# the sprite is three wide; x is the center, so we subtract one for the left and add one for the right
def update_crt(crt_screen, crt_row, cycle, x):
    cycle = cycle - (40 * crt_row)
    for pixel in [x-1, x, x+1]:
        if pixel == cycle:
            try:
                crt_screen[crt_row][pixel] = "#"
            except:
                continue
    return crt_screen


def main():
    cpu_instructions = import_cpu_instructions("input.txt")
    crt_cols = 40
    crt_rows = 6
    crt_screen = [[" " for i in range(crt_cols)] for j in range(crt_rows)]
    # adds up all the required cycles so we know how long to make the program list
    total_cycles = 0
    for line in cpu_instructions:
        if line[0] == "addx":
            total_cycles += 3
        else:
            total_cycles += 2
    cpu_instructions_by_cycle = create_program(cpu_instructions, total_cycles)
    # default starting value
    x = 1
    cycle = 0
    signal_total = 0
    for cycle in range(total_cycles):
        signal_strength = cycle * x
        if cycle in [20,60,100,140,180,220]:
            signal_total += signal_strength
            print(f"cycle: {cycle}, x: {x}, signal_strength: {x * cycle }")
        instruction = cpu_instructions_by_cycle[cycle]
        if instruction:
            command = instruction[0]
            if command == "addx":
                value = instruction[1]
                x += int(value)
        if cycle >= 0 and cycle < 40:
            crt_row = 0
            crt_screen = update_crt(crt_screen, crt_row, cycle, x)
        elif cycle >= 40 and cycle < 80:
            crt_row = 1
            crt_screen = update_crt(crt_screen, crt_row, cycle, x)
        elif cycle >= 80 and cycle < 120:
            crt_row = 2
            crt_screen = update_crt(crt_screen, crt_row, cycle, x)
        elif cycle >= 120 and cycle < 160:
            crt_row = 3
            crt_screen = update_crt(crt_screen, crt_row, cycle, x)
        elif cycle >= 160 and cycle < 200:
            crt_row = 4
            crt_screen = update_crt(crt_screen, crt_row, cycle, x)
        elif cycle >= 200 and cycle < 240:
            crt_row = 5
            crt_screen = update_crt(crt_screen, crt_row, cycle, x)
    print(f"Signal strength total (part 1): {signal_total}")
    for line in crt_screen:
        print("".join(line))
      

if __name__ == '__main__':
    main()