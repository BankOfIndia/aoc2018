package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func getData () []int {
	var array []int
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan () {
		value, _ := strconv.Atoi(input.Text())
		array = append(array, value)
	}
	return array
}

func partOne (list []int) int {
	end_sum := 0
	for _, x := range(list) {
		end_sum += x
	}
	return end_sum
}

func partTwo (list []int) int {
	a := 0
	freqs := make(map[int]string)
	for i := 0; i < len(list); i++ {
		freqs[a] = ""
		a+=list[i]
		if _, ok := freqs[a]; ok {
			return a
		}
		if i == len(list)-1 {
			i = -1
		}
	}
	return 0
}
func main () {
	data := getData()
	fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}