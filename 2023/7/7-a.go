package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isSmaller(a rune, b rune) bool {
	lables := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	return slices.Index(lables, a) > slices.Index(lables, b)
}

func strongestHand(a string, b string) int {
	x := []rune(a)
	y := []rune(b)
	i := 0
	for i < 5 {
		if x[i] != y[i] {
			if isSmaller(x[i], y[i]) {
				return 1
			}
			return -1
		}
		i++
	}
	return 0
}

func points(table map[string]int, hands []string, startRank int, startRes int) (int, int) {
	println("RES PRE:", startRes, "RANK PRE", startRank)
	rank := startRank
	res := startRes
	for _, hand := range hands {
		res += table[hand] * rank
		rank--
	}
	println("RES:", res, "RANK:", rank)
	return res, rank
}

func printHands(hands []string) {
	for _, hand := range hands {
		println(hand)
	}
}

func main() {
	file, err := os.Open("7-a.txt")
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
		} else {
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
		}

		bid, err := strconv.Atoi(splitted[1])
		if err != nil {
			println(err)
		}
		table[hand] = bid
	}

	slices.SortFunc(fiveOfKind, strongestHand)
	slices.SortFunc(fourOfKind, strongestHand)
	slices.SortFunc(fullHouse, strongestHand)
	slices.SortFunc(threeOfKind, strongestHand)
	slices.SortFunc(pair, strongestHand)
	slices.SortFunc(high, strongestHand)

	println("FIVE:", len(fiveOfKind))
	printHands(fiveOfKind)
	println("FOUR:", len(fourOfKind))
	printHands(fourOfKind)
	println("FULL:", len(fullHouse))
	printHands(fullHouse)
	println("THREE:", len(threeOfKind))
	printHands(threeOfKind)
	println("PAIR:", len(pair))
	printHands(pair)
	println("HIGH:", len(high))
	printHands(high)

	res, rank := points(table, fiveOfKind, len(table), 0)
	res, rank = points(table, fourOfKind, rank, res)
	res, rank = points(table, fullHouse, rank, res)
	res, rank = points(table, threeOfKind, rank, res)
	res, rank = points(table, pair, rank, res)
	res, rank = points(table, high, rank, res)

	println("LEN TABLE:", len(table))
	println("RES:", res)
}
