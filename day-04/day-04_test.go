package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	if isValidPassword("111111") != true {
		t.Errorf("Error")
	}
	if isValidPassword("223450") == true {
		t.Errorf("Error")
	}
	if isValidPassword("123789") == true {
		t.Errorf("Error")
	}	
}

func TestPart2(t *testing.T) {
	if isValidPassword_Part2("112233") != true {
		t.Errorf("Error")
	}
	if isValidPassword_Part2("123444") == true {
		t.Errorf("Error")
	}
	if isValidPassword_Part2("111122") != true {
		t.Errorf("Error")
	}	
}