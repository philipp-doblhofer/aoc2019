package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func codeExecutor(memory []int64) []int64 {	
	relative_base := int64(0)
	
	getMemoryAt := func(addr int64) *int64 {
		for int64(len(memory)) <= addr {
			memory = append(memory, 0)
		}
		return &memory[addr]
	}
	
	for ptr := int64(0); ptr < int64(len(memory)); ptr++ {
		opcode := memory[ptr] % 100
		
		getParameter := func(parameter int64) *int64 {
			mode := memory[ptr] / 100
			for i:=parameter; i > 1; i-- {
				mode = mode / 10
			}
			mode = mode % 10
			
			if mode == 0 {
				// position mode
				return getMemoryAt(memory[ptr + parameter])
			} else if mode == 1 {
				// immediate mode
				return &memory[ptr + parameter]
			} else if mode == 2 {
				// relative mode
				return getMemoryAt(relative_base + memory[ptr + parameter])
			}
			
			return nil
		}
	
		switch opcode {
			case 1:
				// Add
				*getParameter(3) = *getParameter(1) + *getParameter(2)
				ptr = ptr + 3
			case 2:
				// Multiply
				*getParameter(3) = *getParameter(1) * *getParameter(2)
				ptr = ptr + 3
			case 3:
				// Read
				
				scanner := bufio.NewScanner(os.Stdin)
				for {
					fmt.Print("Input: ")
					scanner.Scan()
					text := scanner.Text()
					if(len(text) != 0) {
						*getParameter(1) = toint64(text)
						break
					}
		
				}
				ptr = ptr + 1
			case 4:
				// Write
				
				fmt.Printf("(%d)\n", *getParameter(1))
				ptr = ptr + 1
			case 5:
				// Jump-if-true
				
				if(*getParameter(1) != 0) {
					ptr = *getParameter(2) - int64(1) // minus 1, cause the for loop is adding 1
					continue
				}
				ptr = ptr + 2
			case 6:
				// Jump-if-false
				
				if(*getParameter(1) == 0) {
					ptr = *getParameter(2) - int64(1) // minus 1, cause the for loop is adding 1
					continue
				}
				ptr = ptr + 2
			case 7:
				// less-than
				
				if *getParameter(1) < *getParameter(2) {
					*getParameter(3) = 1
				} else {
					*getParameter(3) = 0
				}
				ptr = ptr + 3
			case 8:
				// equals
				
				if *getParameter(1) == *getParameter(2) {
					*getParameter(3) = 1
				} else {
					*getParameter(3) = 0
				}
				ptr = ptr + 3
			case 9:
				// adjust relative base
				
				relative_base += *getParameter(1)
				
				ptr = ptr + 1
			case 99:
				// Halt
				return memory
		}
    }
	
	return memory
}

func Computer() {
	var memory []int64
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 1 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			for _, data := range strings.Split(text, ",") {
				if data != "" {
					memory = append(memory, toint64(data))
				}
			}
		} else {
			break
		}
		
	}
	
	fmt.Println("Executing code ...");
	
	codeExecutor(memory)
}

func toint64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil { panic(err) }
	return i;
}

func main() {
	Computer()
}