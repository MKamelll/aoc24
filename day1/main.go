package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//test := "test.txt"
	input := "input.txt"

	content, err := os.ReadFile(input)

	if err != nil {
		log.Fatal("Couldn't read file, " + err.Error())
	}

	lines := strings.Split(string(content), "\n")

	var first_list []int
	var second_list []int

	for _, elm := range lines {
		numbers := strings.Split(elm, " ")
		var numbers_filtered []string
		for _, str := range numbers {
			if str != "" {
				numbers_filtered = append(numbers_filtered, str)
			}
		}
		num1, err := strconv.Atoi(string(numbers_filtered[0]))
		if err != nil {
			log.Fatal(err.Error())
		}

		num2, err := strconv.Atoi(string(numbers_filtered[len(numbers_filtered)-1]))

		if err != nil {
			log.Fatal(err.Error())
		}

		first_list = append(first_list, num1)
		second_list = append(second_list, num2)
	}

	sort.Ints(first_list)
	sort.Ints(second_list)

	var distances []int

	for i, _ := range first_list {
		dis := first_list[i] - second_list[i]
		if dis < 0 {
			distances = append(distances, dis*-1)
		} else {
			distances = append(distances, dis)
		}
	}

	var result int

	for _, elm := range distances {
		result += elm
	}

	var similarities []int

	for _, e1 := range first_list {
		var count int
		for _, e2 := range second_list {
			if e2 == e1 {
				count += 1
			}
		}
		similarities = append(similarities, e1*count)
	}

	var result2 int

	for _, elm := range similarities {
		result2 += elm
	}

	fmt.Println(result2)

}
