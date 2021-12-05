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

func getSumThreeMeasurement(data []int) int {
	total, currentSum, prevSum := 0, 0, 0

	for i := 0; i < len(data)-3; i++ {
		currentSum = data[i] + data[i+1] + data[i+2]

		if currentSum > prevSum {
			total += 1
		}
		prevSum = currentSum
	}
	return total
}

func main() {
	data := parseFile()

	getMeasurement(data)
	getSumThreeMeasurement(data)

}
