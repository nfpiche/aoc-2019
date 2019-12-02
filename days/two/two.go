package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	needle := flag.Int("n", 0, "What are we looking for")
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		fmt.Printf(err.Error())
		panic("Where's the file nate")
	}

	flag.Parse()
	numbers := getNumbers(string(input))
	var r int

	fmt.Println(*needle)
	if *needle == 0 {
		r = partOneCalc(numbers)
	} else {
		r = partTwoCalc(numbers, *needle)
	}

	fmt.Println(r)
}

func getNumbers(input string) []int {
	asStrings := strings.Split(input, ",")
	numbers := []int{}

	for _, s := range asStrings {
		n, err := strconv.Atoi(strings.TrimSuffix(s, "\n"))

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, n)
	}

	return numbers
}

func partOneCalc(numbers []int) int {
	i := 0

	for {
		v := numbers[i]
		a := numbers[i+1]
		b := numbers[i+2]
		c := numbers[i+3]

		if v == 1 {
			numbers[c] = numbers[a] + numbers[b]
		} else if v == 2 {
			numbers[c] = numbers[a] * numbers[b]
		} else if v == 99 {
			break
		} else {
			fmt.Println("Bad op code ", v)
			panic("Bad op code")
		}

		i += 4
	}

	return numbers[0]
}

func partTwoCalc(numbers []int, needle int) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			clone := make([]int, len(numbers))
			copy(clone, numbers)
			clone[1] = noun
			clone[2] = verb
			r := partOneCalc(clone)

			if r == needle {
				return 100*noun + verb
			}
		}
	}

	panic("no bueno")
}
