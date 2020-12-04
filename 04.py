from collections import *
from functools import lru_cache, reduce
import heapq
import itertools
import math
import random
import sys
import re

def process_batch(items, rules):
    valid = 0
    fields_present = set({})
    for item in items:
        if item == '':
            valid += 1 if len(set(rules.keys()).difference(fields_present)) == 0 else 0
            fields_present = set({})
            continue
        fields = item.split(' ')
        for field in fields:
            key,value = field.split(':', 2)
            if key in rules and rules[key](value):
                fields_present.add(key)

    # Handle the last record, with no separating blank line
    valid += 1 if len(set(rules.keys()).difference(fields_present)) == 0 else 0
    return valid

def main():
    results = []

    with open('04-input.txt', 'r') as f:
        items = [i.strip() for i in f.read().splitlines()]

    rules = {
        'byr': lambda x: True,
        'iyr': lambda x: True,
        'eyr': lambda x: True,
        'hgt': lambda x: True,
        'hcl': lambda x: True,
        'ecl': lambda x: True,
        'pid': lambda x: True
    }
    results.append(process_batch(items, rules))

#byr (Birth Year) - four digits; at least 1920 and at most 2002.
#iyr (Issue Year) - four digits; at least 2010 and at most 2020.
#eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
#hgt (Height) - a number followed by either cm or in:
  #If cm, the number must be at least 150 and at most 193.
  #If in, the number must be at least 59 and at most 76.
#hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
#ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
#pid (Passport ID) - a nine-digit number, including leading zeroes.
#cid (Country ID) - ignored, missing or not.
#Your job is to count the passports where all required fields 

    rules = {
        'byr': lambda x: x.isnumeric() and 1920 <= int(x) <= 2002,
        'iyr': lambda x: x.isnumeric() and 2010 <= int(x) <= 2020,
        'eyr': lambda x: x.isnumeric() and 2020 <= int(x) <= 2030,
        'hgt': lambda x: x[:-2].isnumeric() and (150 <= int(x[:-2]) <= 193 if x[-2:] == 'cm' else 59 <= int(x[:-2]) <= 76),
        'hcl': lambda x: re.fullmatch(r'\#[0-9a-f]{6}', x),
        'ecl': lambda x: x in ['amb','blu','brn','gry','grn','hzl','oth'],
        'pid': lambda x: re.fullmatch(r'\d{9}', x)
    }
    results.append(process_batch(items, rules))

    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    main()