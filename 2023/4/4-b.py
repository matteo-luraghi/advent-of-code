cards = {}
for i in range(1, 224):
    cards[str(i)] = 1
id = 0
while True:
    try:
        line = input("")
    except EOFError:
        break
    

    id = line.split(":")[0].split(" ")[-1]
    print("Working on", id)
    winners = line.split(":")[1].split("|")[0].split(" ")
    used = line.split(":")[1].split("|")[1].split(" ")
    
    points = 0
    for num in used:
        if num != '':
            if num in winners:
                points += 1
    for i in range(1, points+1):
        cards[str(int(id)+i)] += 1*int(cards[id])


print(cards)
sum = 0
for key in range(1, int(id)+1):
    sum += cards[str(key)]
print(sum)
