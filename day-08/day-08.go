package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"math"
)

func Part1() int {
	input, _ := ioutil.ReadFile("input.txt")

	width := 25
	height := 6
	
	cur_x := 0
	cur_y := 0
	cur_layer := 0
	
	layers := make(map[int]map[int]map[int] byte)
	
	for i := range input {
		b := input[i]
		
		layer, ok := layers[cur_layer]
		if !ok {
			layer = make(map[int]map[int] byte)
			layers[cur_layer] = layer
		}
		
		_, ok = layers[cur_layer][cur_x]
		if !ok {
			xval := make(map[int] byte)
			layers[cur_layer][cur_x] = xval
		}
		
		layers[cur_layer][cur_x][cur_y] = b
		
		cur_x = (cur_x + 1) % width
		if cur_x == 0 {
			cur_y = (cur_y + 1) % height
		}
		if cur_x == 0 && cur_y == 0 {
			cur_layer++
		}
	}
	
	bestLayer_0 := math.MaxInt16
	bestLayer_result := 0
	
	for cur_layer--; cur_layer >= 0;  cur_layer-- {
		zeros := 0
		ones := 0
		twos := 0
		
		for cur_x = 0; cur_x < width; cur_x++ {
			for cur_y = 0; cur_y < height; cur_y++ {
				switch(layers[cur_layer][cur_x][cur_y]) {
					case '0':
						zeros++
					case '1':
						ones++
					case '2':
						twos++
				}
			}
		}
		
		if zeros < bestLayer_0  {
			bestLayer_0 = zeros
			bestLayer_result = ones*twos
		}
	}
	
	return bestLayer_result
}

func Part2() {
	input, _ := ioutil.ReadFile("input.txt")

	width := 25
	height := 6
	image := [25][6]byte{}
	
	cur_x := 0
	cur_y := 0
	cur_layer := 0
	
	layers := make(map[int]map[int]map[int] byte)
	
	for i := range input {
		b := input[i]
		
		layer, ok := layers[cur_layer]
		if !ok {
			layer = make(map[int]map[int] byte)
			layers[cur_layer] = layer
		}
		
		_, ok = layers[cur_layer][cur_x]
		if !ok {
			xval := make(map[int] byte)
			layers[cur_layer][cur_x] = xval
		}
		
		layers[cur_layer][cur_x][cur_y] = b
		
		cur_x = (cur_x + 1) % width
		if cur_x == 0 {
			cur_y = (cur_y + 1) % height
		}
		if cur_x == 0 && cur_y == 0 {
			cur_layer++
		}
	}
	
	
	for cur_layer--; cur_layer >= 0;  cur_layer-- {
		for cur_x = 0; cur_x < width; cur_x++ {
			for cur_y = 0; cur_y < height; cur_y++ {
				switch(layers[cur_layer][cur_x][cur_y]) {
					case '0':
						image[cur_x][cur_y] = ' ';
					case '1':
						image[cur_x][cur_y] = '#';
						
				}
			}
		}
	}
	
	for cur_y = 0; cur_y < height; cur_y++ {
		for cur_x = 0; cur_x < width; cur_x++ {
			fmt.Printf("%c", image[cur_x][cur_y]);
		}
		fmt.Printf("\n");
	}
	
}



func toInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil { panic(err) }
	return i;
}

func main() {
	fmt.Printf("Output Part 1: %d\n", Part1())
	fmt.Println("Part 2:")
	Part2()
}