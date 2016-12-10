package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	puzzleInput := "R2, L3, R2, R4, L2, L1, R2, R4, R1, L4, L5, R5, R5, R2, R2, R1, L2, L3, L2, L1, R3, L5, R187, R1, R4, L1, R5, L3, L4, R50, L4, R2, R70, L3, L2, R4, R3, R194, L3, L4, L4, L3, L4, R4, R5, L1, L5, L4, R1, L2, R4, L5, L3, R4, L5, L5, R5, R3, R5, L2, L4, R4, L1, R3, R1, L1, L2, R2, R2, L3, R3, R2, R5, R2, R5, L3, R2, L5, R1, R2, R2, L4, L5, L1, L4, R4, R3, R1, R2, L1, L2, R4, R5, L2, R3, L4, L5, L5, L4, R4, L2, R1, R1, L2, L3, L2, R2, L4, R3, R2, L1, L3, L2, L4, L4, R2, L3, L3, R2, L4, L3, R4, R3, L2, L1, L4, R4, R2, L4, L4, L5, L1, R2, L5, L2, L3, R2, L2"
	fmt.Println(ProcessInput(puzzleInput))
}

type Instruction struct {
	direction string
	distance  int
}

var N, S, E, W int
var Pointing string

func reset() {
	N = 0
	S = 0
	E = 0
	W = 0
	Pointing = "N"
}

func ProcessInput(input string) float64 {
	reset()
	instructions := getInstructions(input)
	for _, inst := range instructions {
		switch inst.direction {
		case "R":
			turnRight()
		case "L":
			turnLeft()
		}
		updatePointers(inst.distance)
	}
	return math.Abs(float64(N-S)) + math.Abs(float64(E-W))
}

func getInstructions(input string) []Instruction {
	split := strings.Split(input, ",")
	n := len(split)
	instructions := make([]Instruction, n, n)
	for i, item := range split {
		item = strings.TrimSpace(item)
		dir := item[0:1]
		dist, err := strconv.Atoi(item[1:len(item)])
		if err != nil {
			continue
		}
		instruction := Instruction{dir, dist}
		instructions[i] = instruction
	}
	return instructions
}

func turnRight() {
	switch Pointing {
	case "N":
		Pointing = "E"
	case "E":
		Pointing = "S"
	case "S":
		Pointing = "W"
	case "W":
		Pointing = "N"
	}
}

func turnLeft() {
	switch Pointing {
	case "N":
		Pointing = "W"
	case "E":
		Pointing = "N"
	case "S":
		Pointing = "E"
	case "W":
		Pointing = "S"
	}
}

func updatePointers(distance int) {
	switch Pointing {
	case "N":
		N += distance
	case "E":
		E += distance
	case "S":
		S += distance
	case "W":
		W += distance
	}
}
