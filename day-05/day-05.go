package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func codeExecutor(memory []int) []int {	
	for ptr := 0; ptr < len(memory); ptr++ {
		opcode := memory[ptr] % 100
		mode_parameter_1 := 0
		mode_parameter_2 := 0
		
		if memory[ptr] >= 100 {
			if ((memory[ptr] % 1000) - (memory[ptr] % 100)) == 100 {
				mode_parameter_1 = 1
			}
			if ((memory[ptr] % 10000) - (memory[ptr] % 1000)) == 1000 {
				mode_parameter_2 = 1
			}
		} 
	
		switch opcode {
			case 1:
				// Add
					
				parameter_1 := memory[ptr+1]
				if mode_parameter_1 == 0 {
					parameter_1 = memory[parameter_1]
				}
				parameter_2 := memory[ptr+2]
				if mode_parameter_2 == 0 {
					parameter_2 = memory[parameter_2]
				}
				
				memory[memory[ptr+3]] = parameter_1 + parameter_2
				ptr = ptr + 3
			case 2:
				// Multiply
				parameter_1 := memory[ptr+1]
				if mode_parameter_1 == 0 {
					parameter_1 = memory[parameter_1]
				}
				parameter_2 := memory[ptr+2]
				if mode_parameter_2 == 0 {
					parameter_2 = memory[parameter_2]
				}
				
				memory[memory[ptr+3]] = parameter_1 * parameter_2
				ptr = ptr + 3
			case 3:
				// Read
				
				scanner := bufio.NewScanner(os.Stdin)
				for {
					fmt.Print("Input: ")
					scanner.Scan()
					text := scanner.Text()
					if(len(text) != 0) {
						memory[memory[ptr+1]] = toInt(text)
						break
					}
		
				}
				ptr = ptr + 1
			case 4:
				// Write
				
				fmt.Printf("(%d)\n", memory[memory[ptr+1]])
				ptr = ptr + 1
			case 5:
				// Jump-if-true
				
				parameter_1 := memory[ptr+1]
				if mode_parameter_1 == 0 {
					parameter_1 = memory[parameter_1]
				}
				parameter_2 := memory[ptr+2]
				if mode_parameter_2 == 0 {
					parameter_2 = memory[parameter_2]
				}
				
				if(parameter_1 != 0) {
					ptr = parameter_2 - 1 // minus 1, cause the for loop is adding 1
					continue
				}
				ptr = ptr + 2
			case 6:
				// Jump-if-false
				
				parameter_1 := memory[ptr+1]
				if mode_parameter_1 == 0 {
					parameter_1 = memory[parameter_1]
				}
				parameter_2 := memory[ptr+2]
				if mode_parameter_2 == 0 {
					parameter_2 = memory[parameter_2]
				}
				
				if(parameter_1 == 0) {
					ptr = parameter_2 - 1 // minus 1, cause the for loop is adding 1
					continue
				}
				ptr = ptr + 2
			case 7:
				// less-than
				
				parameter_1 := memory[ptr+1]
				if mode_parameter_1 == 0 {
					parameter_1 = memory[parameter_1]
				}
				parameter_2 := memory[ptr+2]
				if mode_parameter_2 == 0 {
					parameter_2 = memory[parameter_2]
				}
				
				memory[memory[ptr+3]] = 0
				if parameter_1 < parameter_2 {
					memory[memory[ptr+3]] = 1
				}
				ptr = ptr + 3
			case 8:
				// equals
				
				parameter_1 := memory[ptr+1]
				if mode_parameter_1 == 0 {
					parameter_1 = memory[parameter_1]
				}
				parameter_2 := memory[ptr+2]
				if mode_parameter_2 == 0 {
					parameter_2 = memory[parameter_2]
				}
				
				memory[memory[ptr+3]] = 0
				if parameter_1 == parameter_2 {
					memory[memory[ptr+3]] = 1
				}
				ptr = ptr + 3
			case 99:
				// Halt
				return memory
		}
    }
	
	return memory
}

func Computer() {
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
	
	fmt.Println("Executing code ...");
	
	codeExecutor(memory)
}

func toInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil { panic(err) }
	return i;
}

func main() {
	Computer()
}