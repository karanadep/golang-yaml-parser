package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printperson(per Person) {
	fmt.Println(per)
}

func main() {

	var per1 Person
	var per2 Person

	per1.name = "Karan"
	per1.age = 111
	per1.salary = 10000000
	per1.job = "Scientist"

	per2.name = "Karan"
	per2.age = 111
	per2.salary = 10000000
	per2.job = "Scientist"

	printperson(per1)
	printperson(per2)
}
