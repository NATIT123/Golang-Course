package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	job  string
	year int
}

func (p Person) CaculateAge() {
	fmt.Println(time.Now().Year() - p.year)
}

func (p Person) ValidJob() bool {
	isValid := p.year / len(p.name)
	if isValid == 0 {
		fmt.Println("Invalid job")
		return false
	}
	fmt.Println("Valid job")
	return true
}

func main() {
	p := Person{name: "Tu", job: "Software Engineer", year: 2003}
	fmt.Println(p)
	p.CaculateAge()
	p.ValidJob()
	ex2("Hello")

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(readFile("a.txt"))
}
