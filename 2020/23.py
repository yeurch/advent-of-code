import time

class CircularList:
    def __init__(self, nodes):
        node = Node(value=nodes.pop(0))
        self.__node_dict = {node.value: node}
        self.head = node
        for elem in nodes:
            node.next = Node(value=elem)
            node = node.next
            self.__node_dict[elem] = node
        node.next = self.head

    def advance(self):
        self.head = self.head.next

    def get(self, n):
        return self.__node_dict[n]

    def get_list(self, a, b):
        node = self.head
        for i in range(b):
            if i >= a:
                yield node
            node = node.next

class Node:
    def __init__(self, value):
        self.value = value
        self.next = None

def run_simulation(items, n, part):
    max_item = max(items)
    min_item = min(items)

    data = CircularList(items)
    for i in range(n):
        target = data.head.value
        first_four = list(data.get_list(0, 4))
        first_four_values = [x.value for x in first_four]
        while target in first_four_values:
            target -= 1
            if target < min_item: target = max_item
        target_node = data.get(target)

        # Move three elements
        data.head.next = first_four[-1].next
        first_four[3].next = target_node.next
        target_node.next = first_four[1]

        # Advance the first element
        data.advance()

    node = data.get(1).next

    if part == 1:
        result = []
        while node.value != 1:
            result.append(str(node.value))
            node = node.next

        return ''.join(result)
    else:
        result = 1
        for i in range(2):
            result *= node.value
            node = node.next
        return result

def main():
    results = []

    with open('23-input.txt', 'r') as f:
        input = [i.strip() for i in f.read()]

    items = [int(c) for c in input]
    results.append(run_simulation(items, 100, part=1))

    items = [int(c) for c in input]
    for i in range(len(items)+1, 1000001):
        items.append(i)
    results.append(run_simulation(items, 10000000, part=2))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))