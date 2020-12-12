import time

def part_1(steps):
    heading = 0 # 0=east, 1=south, 2=west, 3=north
    vectors = [(1, 0), (0, -1), (-1, 0), (0, 1)]
    x_pos, y_pos = 0, 0
    for step in steps:
        cmd = step[0]
        if cmd == 'N': y_pos += step[1]
        elif cmd == 'E': x_pos += step[1]
        elif cmd == 'S': y_pos -= step[1]
        elif cmd == 'W': x_pos -= step[1]
        elif cmd == 'R': heading = (heading + (step[1] // 90)) % 4
        elif cmd == 'L': heading = (heading + ((360 - step[1]) // 90)) % 4
        elif cmd == 'F':
            x_pos += step[1] * vectors[heading][0]
            y_pos += step[1] * vectors[heading][1]
        else: raise Exception(f'Unexpected command: {cmd}')
    return abs(x_pos) + abs(y_pos)

def part_2(steps):
    x_pos, y_pos = 0, 0
    waypoint_x, waypoint_y = 10, 1
    for step in steps:
        cmd = step[0]
        theta = 0
        if cmd == 'R':
            theta = step[1] // 90
        elif cmd == 'L':
            theta = (360 - step[1]) // 90
        if cmd == 'N': waypoint_y += step[1]
        elif cmd == 'E': waypoint_x += step[1]
        elif cmd == 'S': waypoint_y -= step[1]
        elif cmd == 'W': waypoint_x -= step[1]
        elif cmd == 'F':
            x_pos += step[1] * waypoint_x
            y_pos += step[1] * waypoint_y
        elif theta > 0:
            for i in range(theta):
                temp = waypoint_y
                waypoint_y = 0 - waypoint_x
                waypoint_x = temp
        else: raise Exception(f'Unexpected command: {cmd}')
        #print(f'ship=(); waypoint=()')
    return abs(x_pos) + abs(y_pos)

def main():
    results = []

    with open('12-input.txt', 'r') as f:
        items = [(i[0], int(i[1:])) for i in f.read().splitlines()]

    results.append(part_1(items))
    results.append(part_2(items))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time()
    main()
    print("--- Executed in %s seconds ---" % (time.time() - start_time))
