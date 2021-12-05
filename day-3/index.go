package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseFile() []string {
	b, err := ioutil.ReadFile("day-3/input.txt")

	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	arr := strings.Fields(str)

	return arr
}

func toBinaryNumber(val string) int64 {
	output, err := strconv.ParseInt(val, 2, 64)

	if err != nil {
		fmt.Println(err)
	}
	return output
}

func getGammaAndEpsilon(arr []string) int64 {

	bitLen := len(strings.Split(arr[0], ""))
	bitArr := make([]int, bitLen)
	gamma, epsilon := "", ""

	for _, bit := range arr {
		bits := strings.Split(bit, "")

		for i, num := range bits {

			if num == "0" {
				bitArr[i]++
			} else {
				bitArr[i]--
			}
		}
	}
	for _, num := range bitArr {
		if num < 0 {
			gamma += "1"
			epsilon += "0"

		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	return toBinaryNumber(gamma) * toBinaryNumber(epsilon)

}

func main() {
	data := parseFile()
	getGammaAndEpsilon(data)
}
