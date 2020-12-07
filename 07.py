import re

def parse_rules(items):
    rules = {}
    bag_regex = re.compile(r'(\d+) (.+) bags?')

    for rule in items:
        keyval = rule.split(' bags contain ')
        key = keyval[0]
        contents = keyval[1][:-1].split(', ')
        if 'no other bags' in contents: contents.remove('no other bags')
        bag_rules = []
        for inner_bag in contents:
            m = bag_regex.fullmatch(inner_bag)
            qty = int(m.group(1))
            colour = m.group(2)
            bag_rules.append((qty, colour))
        rules[key] = bag_rules
    return rules

def bags_inside(colour, rules):
    result = 1 # Don't forget this bag!
    for rule in rules[colour]:
        result += rule[0] * bags_inside(rule[1], rules)
    return result

def main():
    results = []

    with open('07-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    rules = parse_rules(items)

    # Part 1
    my_bag = 'shiny gold'
    containers = {my_bag}
    old_len = 0
    while len(containers) != old_len:
        old_len = len(containers)
        for _, (colour,contents) in enumerate(rules.items()):
            if len([x for x in contents if x[1] in containers]) > 0:
                containers.add(colour)
    results.append(len(containers) - 1) # Don't forget to exclude my bag

    # Part 2
    results.append(bags_inside(my_bag, rules) - 1) # Don't forget to exclude my bag

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    main()