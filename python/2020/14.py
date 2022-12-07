import re
import time

set_regex = re.compile(r'mem\[(\d+)\] = (\d+)')

def binToDec(b):
    result, i = 0, 0
    while len(b) > 0: 
        digit = int(b[-1])
        b = b[:-1]
        result = result + digit * pow(2, i) 
        i += 1
    return result

def execute_prog(program, part):
    memory = dict()
    for instruction in program:
        set_match = set_regex.match(instruction)
        if not set_match:
            mask = instruction[7:]
        elif part == 1:
            addr = set_match.group(1)
            value = bin(int(set_match.group(2)))[2:].zfill(36)
            memory[addr] = ''.join([value[i] if mask[i] == 'X' else mask[i] for i in range(36)])
        else: # part 2
            addr = bin(int(set_match.group(1)))[2:].zfill(36)
            value = bin(int(set_match.group(2)))[2:].zfill(36)
            addresses_stack = [''.join([addr[i] if mask[i] == '0' else mask[i] for i in range(36)])]
            if len(addresses_stack) != 1: raise Exception("boom")
            while len(addresses_stack) > 0:
                a = addresses_stack.pop()
                if not 'X' in a:
                    memory[a] = value
                else:
                    addresses_stack.append(a.replace('X', '0', 1))
                    addresses_stack.append(a.replace('X', '1', 1))
    return sum([binToDec(v) for _,(k,v) in enumerate(memory.items())])

def main():
    results = []

    with open('14-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    results.append(execute_prog(items, 1))
    results.append(execute_prog(items, 2))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in {0:.3f} milliseconds ---".format((time.time() - start_time)*1000))