package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNum() int {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString(' ')
	if err != nil {
		fmt.Println(err)
	}
	line, err = reader.ReadString('\n')
	dataArr := strings.Split(line, " ")
	dataStr := ""
	for _, el := range dataArr {
		dataStr += el
	}
	dataStr = strings.Replace(dataStr, "\n", "", -1)
	dataInt, err := strconv.Atoi(dataStr)
	if err != nil {
		fmt.Println(err)
	}
	return dataInt
}

func main() {
	time := getNum()
	dist := getNum()
	speed := 1
	count := 0
	for speed < time {
		if speed*(time-speed) > dist {
			count++
		}
		speed++
	}
	println(count)
}
