package main

import (
	"fmt"
	"bufio"
	"strconv"
	"math"
	"os"
)

func calculateFuelRequirement_part1(mass int) int {
	return int(math.Floor(float64(mass)/3.0))-2
}


func calculateFuelRequirement_part2(mass int) int {
	// Also calculate fuel requirements for transporting the fuel too (fuel also has some mass)
	fuel := int(math.Floor(float64(mass)/3.0))-2
	if fuel < 0 {
		return 0
	} else {
		return fuel + calculateFuelRequirement_part2(fuel)
	}
}

func Part1() string {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 1 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			sum += calculateFuelRequirement_part1(toInt(text))
		} else {
			break
		}
		
	}
	return strconv.Itoa(sum)
}

func Part2() string {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 2 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			sum += calculateFuelRequirement_part2(toInt(text))
		} else {
			break
		}
		
	}
	return strconv.Itoa(sum)
}

func toInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil { panic(err) }
	return i;
}

func main() {
	fmt.Println("Part 1: " + Part1())
	fmt.Println("Part 2: " + Part2())
}