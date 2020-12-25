import time      

def main():
    results = []

    with open('25-input.txt', 'r') as f:
        public_keys = [int(i.strip()) for i in f.read().splitlines()]

    loop_sizes = []
    for pk in public_keys:
        subject_number = 7
        value = 1
        loop_size = 0
        while value != pk:
            value = (value * subject_number) % 20201227
            loop_size += 1
        loop_sizes.append(loop_size)
        print(f'pk {pk} has loop size {loop_size}')
        
    subject_number = public_keys[1]
    value = 1
    for i in range(loop_sizes[0]):
        value = (value * subject_number) % 20201227
    results.append(value)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))