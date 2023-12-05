sum = 0
line = input("")
while line != None:
    numbers = [int(word) for word in line if word.isdigit()]

    if len(numbers) == 1:
        sum += 10*numbers[0] + numbers[0]
    else:
        sum += 10*numbers[0] + numbers[-1]

    try: 
        line = input("")
    except EOFError:
        print(sum)
        quit()
