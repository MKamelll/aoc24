package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func all_value[T comparable](arr []T, value T) bool {
	for _, elm := range arr {
		if elm != value {
			return false
		}
	}

	return true
}

func is_line_safe(line []int) bool {
	var is_increased bool
	var is_decreased bool
	var is_decreased_or_increased bool = true
	var diffs []int
	if line[0] < line[1] {
		is_increased = true
	} else if line[0] > line[1] {
		is_decreased = true
	}
	for i, _ := range line {
		if i > 0 {
			a := line[i-1]
			b := line[i]

			if a < b && is_decreased {
				is_decreased_or_increased = false
			} else if a > b && is_increased {
				is_decreased_or_increased = false
			} else if a == b {
				is_decreased_or_increased = false
			}

			dif := a - b

			diffs = append(diffs, int(math.Abs(float64(dif))))
		}
	}

	var is_diffs bool = true

	for _, elm := range diffs {
		if elm < 1 || elm > 3 {
			is_diffs = false
			break
		}
	}

	return is_decreased_or_increased && is_diffs
}

func part1(content string) {

	lines := strings.Split(content, "\n")
	var result []bool

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		var numbers_int []int

		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			numbers_int = append(numbers_int, num)
		}

		is_safe := is_line_safe(numbers_int)
		result = append(result, is_safe)

	}

	var count int

	for _, elm := range result {
		if elm == true {
			count += 1
		}
	}

	fmt.Println(count)

}

func part2(content string) {
	lines := strings.Split(content, "\n")
	var result []bool

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		var numbers_int []int

		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			numbers_int = append(numbers_int, num)
		}

		if is_line_safe(numbers_int) {
			result = append(result, true)
			continue
		}

		var mutations [][]int

		for i, _ := range numbers_int {
			if i < len(numbers_int) {
				numbers_int_copy := append([]int{}, numbers_int...)
				mutations = append(mutations, append(numbers_int_copy[:i], numbers_int_copy[i+1:]...))
			}
		}

		for _, mut := range mutations {
			if is_line_safe(mut) {
				result = append(result, true)
				break
			}
		}

	}

	fmt.Println(len(result))

}

func main() {

	//content, err := os.ReadFile("test.txt")
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Couldn't read file, " + err.Error())
	}

	//part1(string(content))
	part2(string(content))
}
