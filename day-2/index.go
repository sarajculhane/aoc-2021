package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseFile() []string {
	b, err := ioutil.ReadFile("day-2/input.txt")

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	arr := strings.Split(str, "\n")

	return arr[0 : len(arr)-1]
}

func getPosition(arr []string) int {
	dirMap := make(map[string]int)

	dirMap["depth"] = 0
	dirMap["horizontal"] = 0

	for _, val := range arr {
		direction := strings.Split(val, " ")
		num, err := strconv.Atoi(direction[1])
		if err != nil {
			panic(err)
		}
		fmt.Println(num, val, dirMap)
		switch direction[0] {
		case "forward":
			dirMap["horizontal"] += num
		case "down":
			dirMap["depth"] += num
		case "up":
			dirMap["depth"] -= num
		}
	}

	return dirMap["horizontal"] * dirMap["depth"]
}

func main() {
	data := parseFile()
	getPosition(data)
}
