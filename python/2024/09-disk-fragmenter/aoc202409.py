import sys, os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))
import pathlib

def parse_data(puzzle_input):
    return puzzle_input

def expand_blocks(data):
    result = []
    file_index = 0
    is_file = True
    for n in [int(c) for c in data]:
        result += [file_index if is_file else -1] * n
        file_index += 1 if is_file else 0
        is_file = not is_file
    return result

def calc_checksum(blocks):
    result = 0
    for i in range(len(blocks)):
        if blocks[i] == -1:
            continue
        result += i * (0 if blocks[i] == -1 else blocks[i])
    return result

def part1(data):
    blocks = expand_blocks(data)
    for i in range(len(blocks)-1, -1, -1):
        if blocks[i] == -1:
            continue
        first_gap = blocks.index(-1)
        if first_gap >= i:
            break
        blocks[first_gap] = blocks[i]
        blocks[i] = -1
    return calc_checksum(blocks)

def fat_to_blocks(fat, num_blocks):
    disk = [-1] * num_blocks
    for i in range(len(fat)):
        file_index, file_offset, file_size = fat[i]
        for b in range(file_size):
            disk[file_offset + b] = file_index
    return disk

def part2(data):
    num_files = (len(data)+1) // 2
    offset = 0
    fat = []
    gaps = []
    for i in range(num_files):
        file_length = int(data[2*i])
        gap_length = int(data[2*i + 1]) if i < num_files - 1 else 0
        fat.append([i, offset, file_length]) # (file_index, offset, length)
        if gap_length > 0:
            gaps.append([offset + file_length, gap_length]) # (offset, length)
        offset += file_length + gap_length

    for i in range(len(fat)-1, -1, -1): # iterate over files in reverse order
        file_index, file_offset, file_size = fat[i]
        for gap in gaps:
            if gap[1] >= file_size and gap[0] < file_offset:
                # we found a gap for this file to go in
                fat[i][1] = gap[0]
                gap[0] += file_size
                gap[1] -= file_size
                break

    # We've defragmented, so now we need a representation of our disk
    disk = fat_to_blocks(fat, offset)
    return calc_checksum(disk)

def solve(puzzle_input):
    data = parse_data(puzzle_input)
    yield part1(data)
    yield part2(data)

if __name__ == "__main__":
    for path in sys.argv[1:]:
            print(f"\n{path}:")
            solutions = solve(puzzle_input=pathlib.Path(path).read_text().rstrip())
            print("\n".join(str(solution) for solution in solutions))
