package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extractNumbers(input string) []int {
	fields := strings.Fields(input)
	numbers := make([]int, len(fields))
	for i, numStr := range fields {
		num := 0
		fmt.Sscanf(numStr, "%d", &num)
		numbers[i] = num
	}
	return numbers
}

func main() {
	file, err := os.Open("4-a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points int
	sum := 0
	for scanner.Scan() {
		points = 0

		splitByDots := strings.Split(scanner.Text(), ":")
		data := splitByDots[1]

		splitByPipe := strings.Split(data, "|")
		beforePipe := splitByPipe[0]
		afterPipe := splitByPipe[1]

		winners := extractNumbers(beforePipe)
		used := extractNumbers(afterPipe)

		fmt.Println(winners)
		fmt.Println(used)

		for _, use := range used {
			for _, winner := range winners {
				if use == winner {
					if points == 0 {
						points++
					} else {
						points *= 2
					}
				}
			}
		}
		fmt.Println("POINTS:", points)
		sum += points
	}
	fmt.Println("SUM:", sum)
}
