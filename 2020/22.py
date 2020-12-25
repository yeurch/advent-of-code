import time
from queue import SimpleQueue

def main():
    results = []

    with open('22-input.txt', 'r') as f:
        input = [i.strip() for i in f.read().splitlines()]
    
    hands = (SimpleQueue(), SimpleQueue())
    p = 0
    for i in [x for x in input if not x.startswith('Player')]:
        if i == '':
            p = 1
            continue
        hands[p].put(int(i))
    print(hands)

    while not (hands[0].empty() or hands[1].empty()):
        cards = (hands[0].get(), hands[1].get())
        winner = 0 if cards[0] > cards[1] else 1
        hands[winner].put(cards[winner])
        hands[winner].put(cards[1-winner])
    
    winner_hand = hands[0] if hands[1].empty() else hands[1]
    num_cards = winner_hand.qsize()
    result = 0
    for i in range(num_cards, 0, -1):
        result += winner_hand.get() * i
    results.append(result)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))