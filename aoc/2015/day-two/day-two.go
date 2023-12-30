package advent2015DayTwo

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func minSize(w int, h int, l int) int {

	a := l * w
	b := w * h
	c := h * l

	if a <= b {
		if a <= c {
			return a
		} else {
			return c
		}
	} else {
		if b <= c {
			return b
		} else {
			return c
		}
	}
}

func PartOne() int {
	fmt.Println("Hello DayTwo")
	data, err := os.ReadFile("./advent-of-code/2015/day-two/input-day-two.txt")
	if err != nil {
		log.Fatal(err)
	}

	boxSizes := strings.Split(string(data), "\n")

	result := 0
	for i := 0; i < len(boxSizes); i++ {
		dim := strings.Split(boxSizes[i], "x")

		w, wErr := strconv.Atoi(dim[0])
		if wErr != nil {
			log.Fatal(err)
		}

		h, hErr := strconv.Atoi(dim[1])
		if hErr != nil {
			log.Fatal(err)
		}

		l, lErr := strconv.Atoi(dim[2])
		if lErr != nil {
			log.Fatal(err)
		}

		minValue := minSize(w, h, l)
		boxSurface := (2 * l * w) + (2 * w * h) + (2 * h * l)
		result += (boxSurface + minValue)

	}
	fmt.Println(result)
	return result
}
