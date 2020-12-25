import time   

def main():
    results = []

    with open('21-input.txt', 'r') as f:
        input = [i.strip() for i in f.read().splitlines()]

    # A list of all ingredients, duplicated as per input
    ingredient_list = []

    # A map of all allergens to their possible ingredients
    all_allergens = {}

    # A set of known allergenic ingredients
    allergenic_ingredients = set()

    # A set of all ingredients
    all_ingredients = set()

    for i in input:
        parts = i.split(' (contains ')
        ingredients = set([ingredient for ingredient in parts[0].split(' ')])
        ingredient_list.extend(ingredients)
        all_ingredients.update(ingredients)
        allergens = parts[1][:-1].split(', ')
        for allergen in allergens:
            if allergen in all_allergens:
                all_allergens[allergen] = all_allergens[allergen] & ingredients
            else:
                all_allergens[allergen] = ingredients
    
    solved = False
    while not solved:
        solved = True
        for _,(k,v) in enumerate(all_allergens.items()):
            if len(v) > 1:
                solved = False
            else:
                (single_ingredient,) = v
                allergenic_ingredients.add(single_ingredient)
        for k in all_allergens:
            if len(all_allergens[k]) > 1:
                all_allergens[k] = all_allergens[k] - allergenic_ingredients

    results.append(len([i for i in ingredient_list if i not in allergenic_ingredients]))

    print(all_allergens)

    ordered_ingredients = []
    for _,(k,v) in enumerate(sorted(all_allergens.items())):
        (ingredient,) = v
        ordered_ingredients.append(ingredient)
    results.append(','.join(ordered_ingredients))
    


    for i,s in enumerate(results):
        print(f'{i+1}: {s}')

if __name__ == '__main__':
    start_time = time.time_ns()
    main()
    print("--- Executed in {0:.3f} seconds ---".format((time.time_ns() - start_time) / (10 ** 9)))