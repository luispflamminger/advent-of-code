priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

def prio(char):
   return priorities.find(char) + 1

with open("./input/day3.txt", "r") as f:
    input = f.read()

rucksacks = input.split("\n")

compartments = list(map(lambda x : [x[:int(len(x)/2)], x[int(len(x)/2):]], rucksacks))

prio_sum = 0
for r in compartments:
    for item in r[0]:
        if item in r[1]:
            prio_sum += prio(item)
            break

print("Part 1: ", prio_sum)

res = sum(map( lambda x : prio(set(filter( lambda y : y in x[1] , x[0])).pop()) , map( lambda x : [x[:int(len(x)/2)], x[int(len(x)/2):]], input.split('\n'))))

print("Part 1 (completely functional): ", res)


groups = []

for i in range(0, len(rucksacks), 3):
    groups.append(rucksacks[i:i+3])

prio_sum = 0
for group in groups:
    for item in group[0]:
        if item in group[1] and item in group[2]:
            prio_sum += prio(item)
            break

print("Part 2: ", prio_sum)