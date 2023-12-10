package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func smaller(a rune, b rune) bool {
	lables := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	return slices.Index(lables, a) > slices.Index(lables, b)
}

func order(hands []string) []string {
	return hands
}

func points(table map[string]int, hands []string, startRank int, startRes int) (int, int) {
	rank := startRank
	res := startRes
	for _, hand := range hands {
		res += table[hand] * rank
		rank--
	}
	return res, rank
}

func main() {
	file, err := os.Open("7-a-test.txt")
	if err != nil {
		println(err)
	}
	defer file.Close()

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

	fiveOfKind = order(fiveOfKind)
	fourOfKind = order(fourOfKind)
	fullHouse = order(fullHouse)
	threeOfKind = order(threeOfKind)
	pair = order(pair)
	high = order(high)

	res, rank := points(table, fiveOfKind, len(table), 0)
	res, rank = points(table, fourOfKind, rank, res)
	res, rank = points(table, fullHouse, rank, res)
	res, rank = points(table, threeOfKind, rank, res)
	res, rank = points(table, pair, rank, res)
	res, rank = points(table, high, rank, res)

	println(res)
}
