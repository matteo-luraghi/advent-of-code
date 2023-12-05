res = 0
while True:
    try:
        line = input("")
    except EOFError:
        print(res)
        break

    id = line.split(":")[0].split(" ")[1]
    sets = line.split(":")[1].split(";")
    
    reds = 0
    greens = 0
    blues = 0

    for set in sets:
        info = set.split(",")
        for el in info:
            el = el.split(" ")
            match el[2]:
                case 'red':
                    if int(el[1]) > reds:
                        reds = int(el[1])
                case 'green':
                    if int(el[1]) > greens:
                        greens = int(el[1])
                case 'blue':
                    if int(el[1]) > blues:
                        blues = int(el[1])
    res += greens*blues*reds
