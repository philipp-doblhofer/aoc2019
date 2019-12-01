package main

import "testing"

func TestPart1(t *testing.T) {
	if calculateFuelRequirement_part1(12) != 2 {
		t.Errorf("Error")
	}
	if calculateFuelRequirement_part1(14) != 2 {
		t.Errorf("Error")
	}
	if calculateFuelRequirement_part1(1969) != 654 {
		t.Errorf("Error")
	}
	if calculateFuelRequirement_part1(100756) != 33583 {
		t.Errorf("Error")
	}
}

func TestPart2(t *testing.T) {
	if calculateFuelRequirement_part2(14) != 2 {
		t.Errorf("Error")
	}
	if calculateFuelRequirement_part2(1969) != 966 {
		t.Errorf("Error")
	}
	if calculateFuelRequirement_part2(100756) != 50346 {
		t.Errorf("Error")
	}
}