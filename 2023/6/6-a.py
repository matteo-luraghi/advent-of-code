times = input("").split(":")[1].split(" ")
distances = input("").split(":")[1].split(" ")

times = [int(x) for x in times if x != ""]
distances = [int(x) for x in distances if x != ""]

res= 1
for i in range(0, len(times)):
    count = 0
    for speed in range(times[i]):
        if speed* (times[i]-speed) > distances[i]:
            count += 1
    res *= count
print(res)
