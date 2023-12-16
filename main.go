package main

import (
	"fmt"

	advent2015DayOne "github.com/LeonAndrade/go-dojo/advent-of-code/2015/day-one"
	advent2015DayTwo "github.com/LeonAndrade/go-dojo/advent-of-code/2015/day-two"
	"github.com/LeonAndrade/go-dojo/fundamentals"
)

func DayOne() {
	advent2015DayOne.PartOne()
	advent2015DayOne.PartTwo()
}

func DayTwo() {
	advent2015DayTwo.PartOne()
}

func main() {
	fmt.Println("=== Go Dojo ===")
	fmt.Println(" ")

	fundamentals.BasicDataStructures()

	DayTwo()
}
