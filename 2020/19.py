from collections import *
from functools import lru_cache, reduce
import heapq
import itertools
import math
import random
import sys
import re
import time

def parse(items):
    i = 0

    rules = dict()
    while True:
        item = items[i]
        i += 1
        if item == '':
            break
        a = item.split(': ')
        key = int(a[0])
        
        if '"' in item:
            # When we have a literal, assume it's the only part of the rule
            rules[key] = a[1].replace('"','')
            continue

        b = a[1].split(' | ')
        values = []
        for c in b:
            value = [int(x) for x in c.split(' ')]
            values.append(value)
        rules[key] = values
    
    messages = items[i:]
    return (rules, messages)

def is_match(x, rules, rule_num=0):
    rule = rules[rule_num]
    if isinstance(rule, list):
        for option in rule:
            consumed = 0
            option_ok = True
            for part_rule in option:
                part_result = is_match(x[consumed:], rules, part_rule)
                if not part_result[0]:
                    option_ok = False
                    break
                consumed += part_result[1]
            if option_ok:
                return (True, consumed)
        return (False, 0)               
    else:
        str_len = len(rule)
        return (True, str_len) if x[:str_len] == rule else (False, 0)

def main():
    results = []

    with open('19-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    rules,messages = parse(items)
    print(rules)
    print(messages)

    # Part 1
    passes = 0
    for msg in messages:
        m = is_match(msg, rules)
        msg_passes = m[0] and m[1] == len(msg)
        print(f'{msg} ... {msg_passes}')
        if msg_passes: passes += 1
    results.append(passes)
    


    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))