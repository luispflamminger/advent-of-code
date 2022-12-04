with open("./input/day4.prod", "r") as f:
    input = f.read()

data = list(map(lambda pair : list(map(lambda range : list(map( lambda id: int(id), range.split("-"))), pair.split(","))), input.split("\n")))
# Resulting data structure: [[[1, 2], [2,3]], [[5, 7], [8, 29]]]


counter = 0
for pair in data:
    # Pairs, where one contains the other
    if pair[0][0] <= pair[1][0] and pair[0][1] >= pair[1][1]  or pair[0][0] >= pair[1][0] and pair[0][1] <= pair[1][1]:
        counter += 1

print("Part 1: ", counter)

counter = 0
for pair in data:
    # Pairs, where they overlap
    if pair[0][0] <= pair[1][0] <= pair[0][1] or pair[0][0] <= pair[1][1] <= pair[0][1] or pair[1][0] <= pair[0][0] <= pair[1][1] or pair[1][0] <= pair[0][1] <= pair[1][1]:
        counter += 1

print("Part 2: ", counter)


