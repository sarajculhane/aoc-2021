package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseFile() []int {
	b, err := ioutil.ReadFile("day-1/input.txt")

	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	arr := strings.Fields(str)
	var numsArr = []int{}

	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numsArr = append(numsArr, j)
	}
	return numsArr
}

func getMeasurement(data []int) int {
	count := 0

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count += 1
		}
	}
	return count
}

func main() {
	data := parseFile()

	getMeasurement(data)
}
