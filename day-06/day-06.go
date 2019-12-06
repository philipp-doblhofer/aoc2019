package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func Part1() {
	orbits := make(map[string] string)
	lines := []string{}
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 1 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			lines = append(lines, text)
			data := strings.Split(text, ")")
			orbits[data[1]] = data[0]
		} else {
			break
		}
		
	}
	
	totalOrbits := 0
	for o := range orbits {
		for {
			parent, ok := orbits[o]
			if ok {
				o = parent
				totalOrbits++
			} else {
				break
			}
		}
	}
	
	fmt.Printf("Total Orbits: %d\n", totalOrbits);
}

func Part2() {
	orbits := make(map[string] string)
	lines := []string{}
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 2 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			lines = append(lines, text)
			data := strings.Split(text, ")")
			orbits[data[1]] = data[0]
		} else {
			break
		}
		
	}
	
	distanceToYou := make(map[string] int)
	o := orbits["YOU"]
	distance := 0
	ok := false
	for {
		distanceToYou[o] = distance
		o, ok = orbits[o]
		if ok {
			distance++
		} else {
			break
		}
	}
	
	o = orbits["SAN"]
	distance = 0
	for {
		pathValue, ok := distanceToYou[o]
		if ok {
			distance += pathValue
			break
		}
		
		o, ok = orbits[o]
		distance++
	}
	
	fmt.Printf("Distance YOU<->SAN: %d\n", distance);
}

func toInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil { panic(err) }
	return i;
}

func main() {
	Part1()
	Part2()
}