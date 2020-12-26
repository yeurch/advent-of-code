import time

directions = [{ # even y
    'ne': (0,1),
    'e': (1,0),
    'se': (0,-1),
    'sw': (-1,-1),
    'w': (-1,0),
    'nw': (-1, 1)
},{ # odd y
    'ne': (1,1),
    'e': (1,0),
    'se': (1,-1),
    'sw': (0,-1),
    'w': (-1,0),
    'nw': (0, 1)
}]

def get_neighbours(tile):
    dirs = directions[tile[1] % 2]
    result = []
    for _,(k,v) in enumerate(dirs.items()):
        result.append((tile[0]+v[0], tile[1]+v[1]))
    return result

def main():
    results = []

    with open('24-input.txt', 'r') as f:
        input = [i.strip() for i in f.read().splitlines()]
    
    black_tiles = set()

    # Part 1
    for route in input:
        moves = []
        i = 0
        while i < len(route):
            if route[i] == 'e' or route[i] == 'w':
                moves.append(route[i])
            else:
                moves.append(route[i] + route[i+1])
                i += 1
            i += 1
        
        x, y = 0, 0
        for move in moves:
            x += directions[y%2][move][0]
            y += directions[y%2][move][1]
        adding = True if (x,y) not in black_tiles else False
        if adding:
            black_tiles.add((x,y))
        else:
            black_tiles.remove((x,y))
    
    results.append(len(black_tiles))

    # Part 2
    for d in range(100):
        white_tiles = {}
        tiles_to_flip = set()
        for black_tile in black_tiles:
            neighbours = get_neighbours(black_tile)
            black_neighbours = 0
            for neighbour in neighbours:
                if neighbour in black_tiles:
                    black_neighbours += 1
                else:
                    if neighbour in white_tiles:
                        white_tiles[neighbour] += 1
                    else:
                        white_tiles[neighbour] = 1
            if black_neighbours == 0 or black_neighbours > 2:
                tiles_to_flip.add(black_tile)
        for _,(k,v) in enumerate(white_tiles.items()):
            if v == 2:
                tiles_to_flip.add(k)
        for tile in tiles_to_flip:
            if tile in black_tiles:
                black_tiles.remove(tile)
            else:
                black_tiles.add(tile)
        #print(f'Day {d}: {len(black_tiles)} black tiles')
    results.append(len(black_tiles))
            
    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))