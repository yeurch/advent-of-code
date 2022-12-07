import time

def exec_prog(p, max_iter=-1):
    pc, n, delta_pc, done = 0, 0, 0, False
    prog_len = len(p)
    while max_iter < 0 or n < max_iter:
        if p[pc] == 1:
            p[p[pc+3]] = p[p[pc+1]] + p[p[pc+2]]
            delta_pc = 4
        elif p[pc] == 2:
            p[p[pc+3]] = p[p[pc+1]] * p[p[pc+2]]
            delta_pc = 4
        elif p[pc] == 99:
            done = True
            break
        else:
            raise Exception(f'Unexpected op code {p[pc]} at pos {pc}')
        pc += delta_pc
        n += 1
        if pc >= prog_len:
            raise Exception(f'Instruction pointer out of bounds. pc={pc}, prog size={len(p)}')
    if not done:
        raise Exception(f'No result obtained after {n} iterations')
    return p[0]

def main():
    results = []

    with open('02-input.txt', 'r') as f:
        items = [int(i) for i in f.read().split(',')]

    # Part 1
    prog = items[:]
    prog[1] = 12
    prog[2] = 2
    results.append(exec_prog(prog))

    # Part 2
    success = False
    for i in range(100):
        for j in range(100):
            prog = items[:]
            prog[1] = i
            prog[2] = j
            result = exec_prog(prog)
            if result == 19690720:
                success = True
                break
        if success:
            break
    results.append(100*i + j)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in {0:.3f} milliseconds ---".format((time.time() - start_time)*1000))