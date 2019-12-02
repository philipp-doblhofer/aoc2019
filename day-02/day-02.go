package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func codeExecutor_part1(memory []int) []int {	
	for ptr := 0; ptr < len(memory); ptr++ {
		switch memory[ptr] {
			case 1:
				// Add
				memory[memory[ptr+3]] = memory[memory[ptr+1]] + memory[memory[ptr+2]]
				ptr = ptr + 3
			case 2:
				// Multiply
				memory[memory[ptr+3]] = memory[memory[ptr+1]] * memory[memory[ptr+2]]
				ptr = ptr + 3
			case 99:
				// Halt
				return memory
		}
    }
	
	return memory
}

func Part1() string {
	var memory []int
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 1 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			for _, data := range strings.Split(text, ",") {
				if data != "" {
					memory = append(memory, toInt(data))
				}
			}
		} else {
			break
		}
		
	}
	
	// Part 1 -> replace values!
	memory[1] = 12
	memory[2] = 02
	
	return strconv.Itoa(codeExecutor_part1(memory)[0])
}

func Part2() string {
	var memory []int
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 2 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			for _, data := range strings.Split(text, ",") {
				if data != "" {
					memory = append(memory, toInt(data))
				}
			}
		} else {
			break
		}
	}
	
	// Part 2 -> replace values!
	for verb := 0; verb < len(memory); verb++ {
		for noun := 0; noun < len(memory); noun++ {
			// Reset Memory
			dat := make([]int, len(memory))
			copy(dat, memory)
			
			dat[1] = noun
			dat[2] = verb
	
			dat = codeExecutor_part1(dat)
			noun := dat[1]
			verb := dat[2]
	
			if dat[0] == 19690720 {
				return strconv.Itoa(100*noun + verb)
			}
		}
	}
	
	return "Nothing found"
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