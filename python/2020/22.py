import time

def state_string(hands):
    hand0 = ','.join([str(x) for x in hands[0]])
    hand1 = ','.join([str(x) for x in hands[1]])
    return ';'.join([hand0, hand1])

def run_game(hands, part=1, recursed=False):
    previous_game_states = set()

    while len(hands[0]) > 0 and len(hands[1]) > 0:
        if part == 2:
            state = state_string(hands)
            if state in previous_game_states:
                winner = 0
                break
            else:
                previous_game_states.add(state)
            cards = (hands[0].pop(0), hands[1].pop(0))
            if cards[0] <= len(hands[0]) and cards[1] <= len(hands[1]):
                (winner,_) = run_game((hands[0][0:cards[0]],hands[1][0:cards[1]]), part, True)
            else:
                winner = 0 if cards[0] > cards[1] else 1
        else:
            cards = (hands[0].pop(0), hands[1].pop(0))
            winner = 0 if cards[0] > cards[1] else 1
        hands[winner].append(cards[winner])
        hands[winner].append(cards[1-winner])

    
    winner_hand = hands[winner]
    num_cards = len(winner_hand)
    result = 0
    if not recursed:
        for i in range(num_cards, 0, -1):
            result += winner_hand.pop(0) * i
    return (winner, result)

def parse_hands(input):
    hands = ([], [])
    p = 0
    for i in [x for x in input if not x.startswith('Player')]:
        if i == '':
            p = 1
            continue
        hands[p].append(int(i))
    return hands

def main():
    results = []

    with open('22-input.txt', 'r') as f:
        input = [i.strip() for i in f.read().splitlines()]
    
    results.append(run_game(parse_hands(input))[1])
    results.append(run_game(parse_hands(input), 2)[1])

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))