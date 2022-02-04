package main

import "fmt"

func test(prices []int) []int {

	output := make([]int, len(prices))
	for i, j := range prices {
		output[i] = i + j
	}

	return output
}

func main() {

	prices := []int{12, 24, 56}
	prices[2] = 11

	b := test(prices)

	fmt.Println(b)

}
