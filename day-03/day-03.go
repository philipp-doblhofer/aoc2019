package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
	"math"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan_distance(x1 int, y1 int, x2 int, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

type Point struct {
    x int
    y int
}

func Task3() string {
	originX := 0
	originY := 0
	
	minDistance := math.MaxInt64
	stepsDone := math.MaxInt64
	
	mapdata := make(map[Point] int)
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the input for part 1 (double empty line to end):\n")
	
	line := 0
	for {
		scanner.Scan()
		text := scanner.Text()
		if(len(text) != 0) {
			line++
			steps := 0
			curX := originX
			curY := originY
			for _, data := range strings.Split(text, ",") {
				if data != "" {
					dirX := 0
					dirY := 0
					val := toInt(data[1:])
					
					if(string(data[0]) == "L") {
						dirX = -1;
						dirY = 0;
					} else if(string(data[0]) == "R") {
						dirX = 1;
						dirY = 0;
					} else if(string(data[0]) == "U") {
						dirX = 0;
						dirY = +1;
					} else if(string(data[0]) == "D") {
						dirX = 0;
						dirY = -1;
					}
					for ; val > 0; val-- {
						curX += dirX
						curY += dirY
						steps++
						
						if(line == 2 && mapdata[Point{curX, curY}] != 0) {
							// Intersection
							dist := manhattan_distance(originX, originY, curX, curY)
							if dist != 0 && (steps+mapdata[Point{curX, curY}]) < stepsDone {
								stepsDone = (steps+mapdata[Point{curX, curY}])
							}
							if dist != 0 && dist < minDistance {
								minDistance = dist
							}
						}
						
						if (line == 1) {
							mapdata[Point{curX, curY}] = steps
						}
					}
					
				}
			}
		} else {
			break
		}
		
	}
	
	return strconv.Itoa(stepsDone) + " | Distance : " + strconv.Itoa(minDistance)
}

func toInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil { panic(err) }
	return i
}

func main() {
	fmt.Println(Task3())
}