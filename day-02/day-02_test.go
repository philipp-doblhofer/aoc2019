package main

import (
	"testing"
	"reflect"
)

func TestPart1(t *testing.T) {
	if !reflect.DeepEqual(codeExecutor_part1([]int{1,9,10,3, 2,3,11,0, 99, 30,40,50}), []int{3500,9,10,70, 2,3,11,0, 99, 30,40,50})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor_part1([]int{1,0,0,0,99}), []int{2,0,0,0,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor_part1([]int{2,3,0,3,99}), []int{2,3,0,6,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor_part1([]int{2,4,4,5,99,0}), []int{2,4,4,5,99,9801})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor_part1([]int{1,1,1,4,99,5,6,0,99}), []int{30,1,1,4,2,5,6,0,99})  {
		t.Errorf("Error")
	}
}

func TestPart2(t *testing.T) {
}