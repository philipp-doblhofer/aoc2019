package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestPart1(t *testing.T) {
	if !reflect.DeepEqual(codeExecutor([]int64{1,9,10,3,2,3,11,0,99,30,40,50}), []int64{3500,9,10,70, 2,3,11,0, 99, 30,40,50})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{1,0,0,0,99}), []int64{2,0,0,0,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{2,3,0,3,99}), []int64{2,3,0,6,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{2,4,4,5,99,0}), []int64{2,4,4,5,99,9801})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{1,1,1,4,99,5,6,0,99}), []int64{30,1,1,4,2,5,6,0,99})  {
		t.Errorf("Error")
	}
}

func TestPart2(t *testing.T) {
	// Test equals(position mode)
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,8,9,10,9,0,0,99,-1,8}), []int64{0,0,8,9,10,9,0,0,99,0,8})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,8,9,10,9,0,0,99,7,8}), []int64{0,0,8,9,10,9,0,0,99,0,8})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,8,9,10,9,0,0,99,9,8}), []int64{0,0,8,9,10,9,0,0,99,0,8})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,8,9,10,9,0,0,99,8,8}), []int64{0,0,8,9,10,9,0,0,99,1,8})  {
		t.Errorf("Error")
	}
	// Test equals(immidiate mode)
	
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1108,-1,8,3,0,0,99}), []int64{0,0,1108,0,8,3,0,0,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1108,7,8,3,0,0,99}), []int64{0,0,1108,0,8,3,0,0,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1108,9,8,3,0,0,99}), []int64{0,0,1108,0,8,3,0,0,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1108,8,8,3,0,0,99}), []int64{0,0,1108,1,8,3,0,0,99})  {
		t.Errorf("Error")
	}
	// Test less-than(position mode)
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,7,9,10,9,0,0,99,-1,8}), []int64{0,0,7,9,10,9,0,0,99,1,8})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,7,9,10,9,0,0,99,7,8}), []int64{0,0,7,9,10,9,0,0,99,1,8})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,7,9,10,9,0,0,99,9,8}), []int64{0,0,7,9,10,9,0,0,99,0,8})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,7,9,10,9,0,0,99,8,8}), []int64{0,0,7,9,10,9,0,0,99,0,8})  {
		t.Errorf("Error")
	}
	// Test less-than(immidiate mode)
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1107,-1,8,3,0,0,99}), []int64{0,0,1107,1,8,3,0,0,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1107,7,8,3,0,0,99}), []int64{0,0,1107,1,8,3,0,0,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1107,9,8,3,0,0,99}), []int64{0,0,1107,0,8,3,0,0,99})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1107,8,8,3,0,0,99}), []int64{0,0,1107,0,8,3,0,0,99})  {
		t.Errorf("Error")
	}
	
	// Test jump(position mode)
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,6,12,15,1,13,14,13,0,0,99,0,0,1,9}), []int64{0,0,6,12,15,1,13,14,13,0,0,99,0,0,1,9})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,6,12,15,1,13,14,13,0,0,99,1,0,1,9}), []int64{0,0,6,12,15,1,13,14,13,0,0,99,1,1,1,9})  {
		t.Errorf("Error")
	}
	
	// Test jump(immidiate mode)
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1105,0,9,1101,0,0,12,0,0,99,0}), []int64{0,0,1105,0,9,1101,0,0,12,0,0,99,0})  {
		t.Errorf("Error")
	}
	if !reflect.DeepEqual(codeExecutor([]int64{0,0,1105,1,9,1101,0,0,12,0,0,99,1}), []int64{0,0,1105,1,9,1101,0,0,12,0,0,99,1})  {
		t.Errorf("Error")
	}
}