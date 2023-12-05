def mapTo(map):
    line = input("")
    try:
        while line != "":
            nums = line.split(" ")
            dist = int(nums[2])
            for i in range(0, dist):
                key = str(int(nums[1])+i)
                map[key] = int(nums[0])+i
            line = input("")
    except EOFError:
        pass

def fixMap(map,key):
    if key not in map.keys():
        map[key] = key

ourSeeds = []
seedToSoil = {}
soilToFert = {}
fertToWater = {}
waterToLight = {}
lightToTemp = {}
tempToHum = {}
humToLoc = {}

while True:
    try:
        line = input("")
    except EOFError:
        break

    if "seeds:" in line:
        ourSeeds = line.split(":")[1].split(" ")

    match line:
        case "seed-to-soil map:":
            mapTo(seedToSoil)
        case "soil-to-fertilizer map:":
            mapTo(soilToFert)
        case "fertilizer-to-water map:":
            mapTo(fertToWater)
        case "water-to-light map:":
            mapTo(waterToLight)
        case "light-to-temperature map:":
            mapTo(lightToTemp)
        case "temperature-to-humidity map:":
            mapTo(tempToHum)
        case "humidity-to-location map:":
            mapTo(tempToHum)

for seed in seedToSoil.keys():
    if seed != '':
        fixMap(seedToSoil,  seed)
        fixMap(soilToFert,  seedToSoil[seed])
        fixMap(fertToWater, soilToFert[seedToSoil[seed]])
        fixMap(waterToLight, fertToWater[soilToFert[seedToSoil[seed]]])
        fixMap(lightToTemp, waterToLight[fertToWater[soilToFert[seedToSoil[seed]]]])
        fixMap(tempToHum,  lightToTemp[waterToLight[fertToWater[soilToFert[seedToSoil[seed]]]]])
        fixMap(humToLoc,  tempToHum[lightToTemp[waterToLight[fertToWater[soilToFert[seedToSoil[seed]]]]]])

for seed in ourSeeds:
    if seed != '':
        print(seed, tempToHum[lightToTemp[waterToLight[fertToWater[soilToFert[seedToSoil[seed]]]]]])
