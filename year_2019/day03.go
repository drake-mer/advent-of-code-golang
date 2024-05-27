package year_2019

import (
	"fmt"
	"strconv"
	"strings"
)

type Coord2D struct {
	x int
	y int
}

type Boolean bool

type WireStep struct {
	steps int
	direction Coord2D
}

func ParseSteps(data string) WireStep{
	direction := data[:1]
	nb_steps, _ := strconv.Atoi(data[1:])
	var increment Coord2D
	switch direction{
	case "U":
		increment = Coord2D{0, 1}
	case "D":
		increment = Coord2D{0, -1}
	case "L":
		increment = Coord2D{-1, 0}
	case "R":
		increment = Coord2D{1, 0}
	default:
		panic(fmt.Sprintf("could not convert %s to a known direction", direction))
	}
	return WireStep{nb_steps, increment}
}


type Vector interface {
	Multiply() Vector
	Add(Vector) Vector
	Scalar(Vector) Vector
}

func (v Coord2D) ScalarProduct(vp Coord2D) int {
	return v.x * vp.x + v.y * vp.y
}

func (v Coord2D) Add(u Coord2D) Coord2D{
	return Coord2D{v.x + u.x, v.y + u.y}
}

func (v Coord2D) Multiply(value int) Coord2D {
	return Coord2D{v.x * value, v.y * value}
}

func abs(value int) int{
	var res int
	if value > 0 {
		res = value
	} else {
		res = (value * (-1))
	}
	return res
}


func FillWireMap(row string)(map[Coord2D]int){
	current_coordinate := Coord2D{0, 0}
	if len(row) == 0 {
		panic("should not happen wtf")
	}
	wire_map := make(map[Coord2D]int)
	// fmt.Printf("working on row %s, length %d\n", row, len(row))
	total_length := 0
	for _, input_data := range(strings.Split(row, ",")){
		step := ParseSteps(input_data)
		for k := 0; k < step.steps; k++ {
			total_length++
			current_coordinate = current_coordinate.Add(step.direction)
			_, exists := wire_map[current_coordinate]
			if !exists {
				wire_map[current_coordinate] = total_length
			}
		}
	}
	fmt.Printf("wire map length %d\n", len(wire_map))
	return wire_map
}


func day03(data [] string) (int, int) {
	map_one := FillWireMap(data[0])
	map_two := FillWireMap(data[1])
	var min_dist int = 100000000
	for key := range map_one {
		_, intersection := map_two[key]
		if intersection {
			distance := abs(key.y) + abs(key.x)
			if distance < min_dist {
				min_dist = distance
			}
		}
	}
	min_dist_2 := 10000000
	for key, l1 := range map_one {
		l2, ok := map_two[key]
		if ok {
			min_dist_2 = min(l1 + l2, min_dist_2)
		}
	}
	return min_dist, min_dist_2
}