package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	sum := 0
	left := -1
	right := -1
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		left = -1
		right = -1
		// fmt.Println(scanner.Text())
		for _, char := range scanner.Text() {
			if unicode.IsDigit(char) {
				buf := make([]byte, 1)
				_ = utf8.EncodeRune(buf, char)
				num, err := strconv.Atoi(string(buf))
				if err != nil {
					log.Fatal(err)
				}
				if left == -1 {
					left = num
				}
				right = num
			}
		}
		sum += left*10 + right
	}
	fmt.Println("Res:", sum)
}
