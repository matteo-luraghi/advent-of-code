from enum import Enum
sum = 0
line = input("")

#spelled(0).name == "one", spelled["one"].value == 0
class spelled(Enum):
    one = 1
    two = 2
    three = 3
    four = 4
    five = 5
    six = 6
    seven = 7
    eight = 8
    nine = 9

def findR(line, numbers):
    for i in range(0, len(line)):
            for spelNum in spelled:
                if spelNum.name in line[:i]:
                    numbers.append(spelNum.value)
                    return numbers
            if line[i].isdigit():
                numbers.append(line[i])
                return numbers 


def findB(line, numbers):
    i = len(line) -1
    while i >= 0:
        for spelNum in spelled:
                if spelNum.name in line[i+1:]:
                    numbers.append(spelNum.value)
                    return numbers
        if line[i].isdigit():
            numbers.append(line[i])
            return numbers
        i -=1
    
while line != None:
    numbers = []
    numbers = findR(line, numbers)
    numbers = findB(line, numbers)
    numbers = [int(x) for x in numbers]

    if len(numbers) == 1:
        sum += 10*numbers[0] + numbers[0]
    else:
        sum += 10*numbers[0] + numbers[1]

    try: 
        line = input("")
    except EOFError:
        print(sum)
        break
