package main

import "fmt"

func main() {
	prices := [3]int{12, 24, 56}
	prices[2] = 11

	fmt.Println(prices)
}
