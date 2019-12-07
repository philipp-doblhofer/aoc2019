package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func codeExecutor_day7(memory []int, input chan int, output chan int, finished chan bool) {
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
				memory[memory[ptr+1]] = <- input
				ptr = ptr + 1
			case 4:
				// Write
				
				output <- memory[memory[ptr+1]]
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
				finished <- true
				return
		}
    }
}

func getThrust() int {
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
	
	
	orig := []int{0, 1, 2, 3, 4}
	ret := 0
	
    for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
    	phase := getPerm(orig, p)
    	
		dat1 := make([]int, len(memory))
		dat2 := make([]int, len(memory))
		dat3 := make([]int, len(memory))
		dat4 := make([]int, len(memory))
		dat5 := make([]int, len(memory))
		inpA := make(chan int)
		inpB := make(chan int)
		inpC := make(chan int)
		inpD := make(chan int)
		inpE := make(chan int)
		oE := make(chan int)
	
		// A
		copy(dat1, memory)
		go codeExecutor_day7(dat1, inpA, inpB, make(chan bool))
		// B
		copy(dat2, memory)
		go codeExecutor_day7(dat2, inpB, inpC, make(chan bool))
		// C
		copy(dat3, memory)
		go codeExecutor_day7(dat3, inpC, inpD, make(chan bool))
		// D
		copy(dat4, memory)
		go codeExecutor_day7(dat4, inpD, inpE, make(chan bool))
		// E
		copy(dat5, memory)
		go codeExecutor_day7(dat5, inpE, oE, make(chan bool))
	
		inpA <- phase[0]
		inpB <- phase[1]
		inpC <- phase[2]
		inpD <- phase[3]
		inpE <- phase[4]
	
		inpA <- 0
		output := <- oE
		if output > ret {
			ret = output
		}
	}
	return ret;
}


func getThrust2() int {
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
	
	fmt.Println("Executing code ...");
	
	
	orig := []int{5, 6, 7, 8, 9}
	ret := 0
	
    for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
    	phase := getPerm(orig, p)
    	
		dat1 := make([]int, len(memory))
		dat2 := make([]int, len(memory))
		dat3 := make([]int, len(memory))
		dat4 := make([]int, len(memory))
		dat5 := make([]int, len(memory))
		inpA := make(chan int)
		inpB := make(chan int)
		inpC := make(chan int)
		inpD := make(chan int)
		inpE := make(chan int)
		
		finished := make(chan bool)
	
		// A
		copy(dat1, memory)
		go codeExecutor_day7(dat1, inpA, inpB, finished)
		// B
		copy(dat2, memory)
		go codeExecutor_day7(dat2, inpB, inpC, finished)
		// C
		copy(dat3, memory)
		go codeExecutor_day7(dat3, inpC, inpD, finished)
		// D
		copy(dat4, memory)
		go codeExecutor_day7(dat4, inpD, inpE, finished)
		// E
		copy(dat5, memory)
		go codeExecutor_day7(dat5, inpE, inpA, finished)
	
		inpA <- phase[0]
		inpB <- phase[1]
		inpC <- phase[2]
		inpD <- phase[3]
		inpE <- phase[4]
	
		inpA <- 0
		
		for i := 0; i<4; i++ {
			<- finished
		}
		
		output := <- inpA
		if output > ret {
			ret = output
		}
	}
	return ret;
}

func toInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil { panic(err) }
	return i;
}

// https://stackoverflow.com/a/30230552/14660
func nextPerm(p []int) {
    for i := len(p) - 1; i >= 0; i-- {
        if i == 0 || p[i] < len(p)-i-1 {
            p[i]++
            return
        }
        p[i] = 0
    }
}

func getPerm(orig, p []int) []int {
    result := append([]int{}, orig...)
    for i, v := range p {
        result[i], result[i+v] = result[i+v], result[i]
    }
    return result
}

func main() {
	fmt.Printf("Output Part 1: %d\n", getThrust())
	fmt.Printf("Output Part 2: %d\n", getThrust2())
}