package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	partOne := flag.Bool("o", false, "Run part one")
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Printf(err.Error())
		panic("Where's the file nate")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	var fn func(i int) int

	if *partOne {
		fn = partOneCalc
	} else {
		fn = partTwoCalc
	}

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Printf(err.Error())
			panic("Something is wrong with input")
		}

		count += fn(mass)
	}

	fmt.Printf("%d\n", count)
}

func partOneCalc(mass int) int {
	return getFuel(mass)
}

func partTwoCalc(mass int) int {
	counter := 0

	for mass > 0 {
		mass = getFuel(mass)

		if mass > 0 {
			counter += mass
		}
	}

	return counter
}

func getFuel(mass int) int {
	return (mass / 3) - 2
}
