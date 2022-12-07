def main():
    results = []

    with open('06-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    # Question 1
    sum = 0
    questions = set({})
    for n,i in enumerate(items):
        for c in i:
            questions.add(c)
        if i == '' or n == len(items)-1:
            sum += len(questions)
            questions = set({})
    results.append(sum)

    # Question 2
    sum = 0
    questions = set([chr(i) for i in range(97, 123)])
    for n,i in enumerate(items):
        person = set([c for c in i])
        questions = questions.intersection(person) if len(person) > 0 else questions
        if i == '' or n == len(items)-1:
            sum += len(questions)
            questions = set([chr(i) for i in range(97, 123)])
    results.append(sum)

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    main()