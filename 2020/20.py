import math
import time

class Tile:
    def __init__(self, id, data, image):
        self.id = id
        self.data = data
        self.image = image
        self.calc_edges()
        self.pos = (None, None)
        self.matches = None

    def __str__(self):
        return '\n'.join(self.data)

    def calc_edges(self):
        self.edges = []
        self.edges.append(self.data[0])
        self.edges.append(''.join([x[-1] for x in self.data]))
        self.edges.append(self.data[-1][::-1])
        self.edges.append(''.join([x[0] for x in self.data][::-1]))

    def calc_matches(self):
        result = [[], [], [], []]
        matches = 0
        my_edges_reverse = [x[::-1] for x in self.edges]
        for is_flip,my_edges in enumerate([self.edges, my_edges_reverse]):
            for _,(id,i) in enumerate(self.image.tiles.items()):
                if id == self.id:
                    continue
                for e in i.edges:
                    if e in my_edges:
                        j = my_edges.index(e)
                        result[j].append((id, i.edges.index(e)))
                        matches += 1
        self.matches = result
        self.num_matching_edges = len([x for x in result if len(x) > 0])
        return result
    
    def rotate(self):
        new_data = []
        for i in range(len(self.data)):
            new_data.append(''.join([x[i] for x in self.data[::-1]]))
        self.data = new_data
        self.calc_edges()
        self.calc_matches()
    
    def flip(self):
        self.data = [x[::-1] for x in self.data]
        self.calc_edges()
        self.calc_matches()
    
    def set_pos(self, x, y):
        self.pos = (x, y)
        self.image.record_tile_pos(self)

class Image:
    def __init__(self, items):
        self.tiles = dict()
        self.tile_positions = dict()
        for i in range((len(items)+1) // 12):
            tile_id = int(items[i*12][5:-1])
            tile_data = items[i*12+1:i*12+11]
            tile = Tile(tile_id, tile_data, self)
            self.tiles[tile_id] = tile
    
    def __str__(self):
        image_size = int(math.sqrt(len(self.tiles)))
        tile_size = len(self.get_tile_at(0, 0).data)
        result = ''
        for line in range(image_size * tile_size):
            if line % tile_size == 0:
                if line > 0:
                    result += '\n'
                tiles = []
                for x in range(image_size):
                    tiles.append(self.get_tile_at(x, line // tile_size))
            for x in range(image_size):
                result += f'{tiles[x].data[line % tile_size]} '
            result += '\n'
        return result
    
    def get_tile(self, id):
        return self.tiles[id]

    def calc_matches(self):
        for k in self.tiles:
            self.tiles[k].calc_matches()

    def record_tile_pos(self, tile):
        self.tile_positions[tile.pos] = tile.id

    def get_tile_at(self, x, y):
        return self.get_tile(self.tile_positions[(x, y)])

    def get_image_data(self):
        result = []
        image_size = int(math.sqrt(len(self.tiles)))
        tile_size = len(self.get_tile_at(0, 0).data) - 2
        for line in range(image_size * tile_size):
            line_str = ''
            if line % tile_size == 0:
                tiles = []
                for x in range(image_size):
                    tiles.append(self.get_tile_at(x, line // tile_size))
            for x in range(image_size):
                line_str += f'{tiles[x].data[(line % tile_size) + 1][1:-1]}'
            result.append(line_str)
        return result
    
    def part1(self):
        self.calc_matches()
        result = 1
        corner_pieces = [k for k in self.tiles if self.tiles[k].num_matching_edges == 2]
        if len(corner_pieces) != 4:
            raise Exception('Difficulty detecting corners')
        for x in corner_pieces:
            result *= x
        return result
    
    def part2(self):
        image_size = int(math.sqrt(len(self.tiles)))
        for pos in range(image_size*image_size):
            xpos = pos % image_size
            ypos = pos // image_size
            if pos == 0:
                # First tile, just pick our first corner
                tile = [self.get_tile(k) for k in self.tiles if self.tiles[k].num_matching_edges == 2][0]
                while True:
                    if (len(tile.matches[1]) + len(tile.matches[2])) == 2:
                        break
                    tile.rotate()
            else:
                if xpos > 0:
                    prev_tile = tile
                    tile = self.get_tile(tile.matches[1][0][0])
                    link_dir = (1,3)
                else:
                    prev_tile = self.get_tile_at(0, ypos - 1)
                    tile = self.get_tile(prev_tile.matches[2][0][0])
                    link_dir = (2,0)
                for x in range(8):
                    if prev_tile.edges[link_dir[0]] == tile.edges[link_dir[1]][::-1]:
                        break
                    tile.rotate()
                    if x == 3: tile.flip()
            tile.set_pos(xpos, ypos)

        sea_monster = ((18,0),(0,1),(5,1),(6,1),(11,1),(12,1),(17,1),(18,1),(19,1),(1,2),(4,2),(7,2),(10,2),(13,2),(16,2))
        raw_image = self.get_image_data()

        for orientations in range(8):
            for x in range(len(raw_image) - 20):
                for y in range(len(raw_image) - 3):
                    is_monster = True
                    for m in range(len(sea_monster)):
                        if raw_image[y + sea_monster[m][1]][x + sea_monster[m][0]] == '.':
                            is_monster = False
                            break
                    if is_monster:
                        for m in range(len(sea_monster)):
                            s = list(raw_image[y + sea_monster[m][1]])
                            s[x + sea_monster[m][0]] = 'O'
                            raw_image[y + sea_monster[m][1]] = ''.join(s)

            # rotate
            new_data = []
            for i in range(len(raw_image)):
                new_data.append(''.join([x[i] for x in raw_image[::-1]]))
            raw_image = new_data

            # flip
            if orientations == 3 or orientations == 7: raw_image = [x[::-1] for x in raw_image]
        
        result = 0
        for y in raw_image:
            result += y.count('#')

        return result


def main():
    results = []

    with open('20-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    image = Image(items)

    results.append(image.part1())
    results.append(image.part2())

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))