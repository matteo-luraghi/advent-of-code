seeds = []
used = []

def add():
    line = input("")
    try:
        while line != "" and line != " ":
            nums = line.split(" ")
            dist = int(nums[2])
            for i in range(0, dist):
                data = int(nums[1])+i
                if data in seeds:
                    seeds.remove(data)
                    used.append(int(nums[0])+i)            
            line = input()
    except EOFError:
        pass

while True:
    try:
        line = input("")
    except EOFError:
        break

    if "seeds:" in line:
        ourSeeds = line.split(":")[1].split(" ")
        for seed in ourSeeds:
            if seed != "":
                seeds.append(int(seed))
        print(seeds)
    if "map" in line:
        add()

    for el in used:
        seeds.append(el)
    used = []

print("MIN", min(seeds))
