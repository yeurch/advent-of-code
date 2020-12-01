with open('01-input.txt', 'r') as f:
    items = [int(i) for i in f.read().splitlines()]

items.sort()

done = False
for i in range(len(items)-1):
    for j in range(i+1, len(items)):
        sum = items[i] + items[j]

        if sum == 2020:
            print(f'{items[i]} * {items[j]} = {items[i] * items[j]}')
            done = True
            break
        elif sum > 2020:
            break
    if done:
        break

done = False
for i in range(len(items)-2):
    for j in range(i+1, len(items)-1):
        for k in range(j+1, len(items)):
            sum = items[i] + items[j] + items[k]

            if sum == 2020:
                print(f'{items[i]} * {items[j]} * {items[k]} = {items[i] * items[j] * items[k]}')
                done = True
                break
            elif sum > 2020:
                break
        if done:
            break
    if done:
        break

