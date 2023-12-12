package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func update(array []int64, scanner *bufio.Scanner) []int64 {
	var used []int64
	for scanner.Scan() {
		var i int64 = 0
		text := scanner.Text()
		if text == "" {
			break
		}
		numsStr := strings.Split(text, " ")
		var nums []int64
		for _, el := range numsStr {
			if el != "" {
				num, err := strconv.Atoi(el)
				if err != nil {
					println(err)
				}
				nums = append(nums, int64(num))
			}
		}
		dist := nums[2]
		for i < dist {
			data := nums[1] + i
			if slices.Contains(array, data) {
				idx := slices.Index(array, data)
				array = append(array[:idx], array[idx+1:]...) //delete the element at the index idx
				used = append(used, int64(nums[0]+i))
			}
			i++
		}
	}
	for _, el := range array {
		used = append(used, el)
	}
	return used
}

func main() {
	file, err := os.Open("5-a.txt")
	if err != nil {
		println(err)
	}
	defer file.Close()

	var seeds []int64

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seedsStr := strings.Split(strings.Split(scanner.Text(), ":")[1], " ")
	for _, el := range seedsStr {
		if el != "" {
			seed, err := strconv.Atoi(el)
			if err != nil {
				println(err)
			}
			seeds = append(seeds, int64(seed))
		}
	}

	for i := range seeds {
		fmt.Printf("%d ", seeds[i])
	}
	fmt.Printf("\n")

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" && strings.Contains(text, "map") {
			seeds = update(seeds, scanner)
			for _, seed := range seeds {
				fmt.Printf("%d ", seed)
			}
			fmt.Printf("\n")
		}
	}

	fmt.Println(slices.Min(seeds))
}
