class AOCGrid():
    def __init__(self, data):
        self.__data = data
        self.width = len(data[0])
        self.height = len(data)

    def from_chars(data):
        return AOCGrid([list(row) for row in data.strip().split('\n')])

    def get_cell(self, x, y):
        if x >= self.width or y >= self.height or x < 0 or y < 0:
            return None
        return self.__data[y][x]

    def set_cell(self, x, y, val):
        if x >= self.width or y >= self.height or x < 0 or y < 0:
            raise ValueError("Provided inputs were out of bounds")
        self.__data[y][x] = val

    def find(self, c):
        for y in range(self.height):
            for x in range(self.width):
                if self.get_cell(x, y) == c:
                    return (x, y)
        return None

    def find_all(self, c):
        result = []
        for y in range(self.height):
            for x in range(self.width):
                if self.get_cell(x, y) == c:
                    result.append((x, y))
        return result

    def count(self, c):
        return len(self.find_all(c))

    def get_ray(self, start_x, start_y, delta_x, delta_y, length):
        result = []
        for i in range(length):
            val = self.get_cell(start_x + i*delta_x, start_y + i*delta_y)
            if val == None:
                return None
            result.append(val)
        return result

    def get_omni_rays(self, start_x, start_y, length):
        result = []
        for dy in range(-1, 2):
            for dx in range(-1, 2):
                if dx == 0 and dy == 0:
                    continue
                ray = self.get_ray(start_x, start_y, dx, dy, length)
                if ray:
                    result.append(ray)
        return result
