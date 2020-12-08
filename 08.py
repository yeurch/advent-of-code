def run_prog(prog):
    acc = 0
    pc = 0
    executed_instructions = set()

    while not (pc in executed_instructions or pc == len(prog)):
        executed_instructions.add(pc)
        op = prog[pc][0]
        arg = prog[pc][1]
        if op == 'jmp':
            pc += arg
            continue
        elif op == 'acc':
            acc += arg
        pc += 1
    return (pc == len(prog), acc)

def main():
    results = []

    with open('08-input.txt', 'r') as f:
        items = [(i[:3], int(i[4:])) for i in f.read().splitlines()]


    # Part 1
    part1 = run_prog(items)
    results.append(part1[1])

    # Part 2
    for i,x in enumerate(items):
        prog = items.copy()
        prog[i] = ('nop' if x[0] == 'jmp' else ('jmp' if x[0] == 'nop' else x[0]), x[1])
        part2 = run_prog(prog)
        if part2[0]:
            results.append(part2[1])
            break

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    main()