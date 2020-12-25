import time

def main():
    results = []

    with open('22-input.txt', 'r') as f:
        input = [i.strip() for i in f.read().splitlines()]
    
    hands = ([], [])
    p = 0
    for i in [x for x in input if not x.startswith('Player')]:
        if i == '':
            p = 1
            continue
        hands[p].append(int(i))

    while len(hands[0]) > 0 and len(hands[1]) > 0:
        cards = (hands[0].pop(0), hands[1].pop(0))
        winner = 0 if cards[0] > cards[1] else 1
        hands[winner].append(cards[winner])
        hands[winner].append(cards[1-winner])
    
    winner_hand = hands[0] if len(hands[1]) == 0 else hands[1]
    num_cards = len(winner_hand)
    result = 0
    for i in range(num_cards, 0, -1):
        result += winner_hand.pop(0) * i
    results.append(result)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))