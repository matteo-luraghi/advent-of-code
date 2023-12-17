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
	var tmp []int64
	var used []int64
	for scanner.Scan() {
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
		for _, el := range array {
			scarto := el - nums[1]
			if scarto >= 0 && scarto < dist {
				tmp = append(tmp, el)
				used = append(used, int64(nums[0]+scarto))
			}
		}

	}
	for _, el := range array {
		if !slices.Contains(tmp, el) {
			used = append(used, el)
		}
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
	var initSeeds []int64

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seedsStr := strings.Split(strings.Split(scanner.Text(), ":")[1], " ")
	for _, el := range seedsStr {
		if el != "" {
			seed, err := strconv.Atoi(el)
			if err != nil {
				println(err)
			}
			initSeeds = append(initSeeds, int64(seed))
		}
	}

	j := 0
	for j < len(initSeeds) {
		if j%2 == 0 {
			var k int64
			for k < initSeeds[j+1] {
				seeds = append(seeds, initSeeds[j]+k)
				k++
			}
		}
		j++
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
