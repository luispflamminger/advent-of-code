
def procedural(text):
    list = text.split("\n\n")

    res_list = []
    for elf in list:
        
        sum = 0
        elf = elf.split("\n")
        for f in elf:
            sum += int(f)
        res_list.append(sum)

    res_list.sort(reverse=True)
    print("Solution Part 1: ", res_list[0])
    print("Solution Part 2: ", res_list[:3])

def functional(text):
    # Ugliest shit i've ever seen
    x = list(map( lambda x : sum(map( lambda x : int(x), x.split("\n"))), text.split("\n\n")))
    x.sort(reverse=True)

    print("Solution Part 1: ", x[0])
    print("Solution Part 2: ", sum(x[:3]))


if __name__ == "__main__":
    with open("./input/day1.txt", "r") as f:
        text = f.read()

    procedural(text)
    functional(text)
