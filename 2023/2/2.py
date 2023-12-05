res = 0
while True:
    try:
        line = input("")
    except EOFError:
        print(res)
        break

    id = line.split(":")[0].split(" ")[1]
    sets = line.split(":")[1].split(";")
    
    no = 0
    for set in sets:
        info = set.split(",")
        for el in info:
            el = el.split(" ")
            match el[2]:
                case 'red':
                    if int(el[1]) > 12:
                        no = 1
                case 'green':
                    if int(el[1]) > 13:
                        no = 1
                case 'blue':
                    if int(el[1]) > 14:
                        no = 1
    if no == 0:
        res += int(id)
