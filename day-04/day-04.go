package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func isValidPassword(t string) bool {
	if len(t) != 6 {
		return false;
	}
	
	hasDoubleDigit := false
	lastDigit := '0'
	
	for _, char := range t {
		if char < lastDigit {
			return false;
		}
		if char == lastDigit {
			hasDoubleDigit = true
		}
		lastDigit = char
	}
	
	return hasDoubleDigit;
}

func isValidPassword_Part2(t string) bool {
	if len(t) != 6 {
		return false;
	}
	
	hasDoubleDigit := false
	digitGroupStarted := 1 
	lastDigit := '0'
	
	for _, char := range t {
		if char < lastDigit {
			return false;
		}
		if char == lastDigit {
			digitGroupStarted++
		} else {
			if digitGroupStarted == 2 {
				hasDoubleDigit = true
			}
			digitGroupStarted = 1
		}
		lastDigit = char
	}
	if digitGroupStarted == 2 {
		hasDoubleDigit = true
	}
	
	return hasDoubleDigit;
}

func Part1() string {
	min := 0
	max := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 1 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			if min == 0 {
				min = toInt(text)
			} else {
				max = toInt(text)
			}
		} else {
			break
		}
		
	}
	cnt := 0
	for i:=min; i<=max; i++ {
		if isValidPassword(strconv.Itoa(i)) {
			cnt++
		}
	}
	return strconv.Itoa(cnt)
}

func Part2() string {
	min := 0
	max := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 2 (double empty line to end):\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			if min == 0 {
				min = toInt(text)
			} else {
				max = toInt(text)
			}
		} else {
			break
		}
		
	}
	cnt := 0
	for i:=min; i<=max; i++ {
		if isValidPassword_Part2(strconv.Itoa(i)) {
			cnt++
		}
	}
	return strconv.Itoa(cnt)
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