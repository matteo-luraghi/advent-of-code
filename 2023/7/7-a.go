package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("7-a-test.txt")
	if err != nil {
		println(err)
	}
	defer file.Close()

	res := 0

	table := make(map[string]int)
	var fiveOfKind []string
	var fourOfKind []string
	var fullHouse []string
	var threeOfKind []string
	var pair []string
	var high []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), " ")
		hand := splitted[0]

		var used []rune
		var points []int

		for _, label := range hand {
			if !slices.Contains(used, label) {
				used = append(used, label)
				count := strings.Count(hand, string(label))
				points = append(points, count)
			}
		}

		if slices.Contains(points, 3) && slices.Contains(points, 2) {
			fullHouse = append(fullHouse, hand)
		}

		value := slices.Max(points)

		switch value {
		case 5:
			fiveOfKind = append(fiveOfKind, hand)
		case 4:
			fourOfKind = append(fourOfKind, hand)
		case 3:
			threeOfKind = append(threeOfKind, hand)
		case 2:
			pair = append(pair, hand)
		default:
			high = append(high, hand)
		}

		bid, err := strconv.Atoi(splitted[1])
		if err != nil {
			println(err)
		}
		table[hand] = bid
	}
	println(res)
}
