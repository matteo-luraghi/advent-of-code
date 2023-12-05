sum = 0
while True:
    try:
        line = input("")
    except EOFError:
        break
    
    points = 0

    winners = line.split(":")[1].split("|")[0].split(" ")
    used = line.split(":")[1].split("|")[1].split(" ")
    
    for num in used:
        if num != '':
            if num in winners:
                if points == 0:
                    points = 1
                else:
                    points *= 2

    sum += points

print(sum)
