package advent

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() (int, []byte) {
	file, err := os.Open("./advent-of-code/2015/day-one/input-day-one.txt")
	check(err)

	data := make([]byte, 1024*12)
	count, err := file.Read(data)
	check(err)

	fmt.Println(count)
	return count, data
}

func DayOne() {
	fmt.Println("Advent of Code | 2015 | Day One")

	count, data := readInput()
	// fmt.Println(string(data))
	floor := 0

	for i := 0; i < count; i++ {
		dir := string(data[i])
		if dir == "(" {
			floor += 1
		}
		if dir == ")" {
			floor -= 1
		}
	}

	fmt.Println(floor)
}

func PartTwo() {
	fmt.Println("Advent of Code | 2015 | Day One Part Two")
	count, data := readInput()

	floor := 0

	for i := 0; i < count; i++ {

		if floor == -1 {
			fmt.Printf("Step: %d | Floor: %d", i, floor)
			fmt.Println(" ")
		}

		dir := string(data[i])
		if dir == "(" {
			floor += 1
		}
		if dir == ")" {
			floor -= 1
		}
	}
}
