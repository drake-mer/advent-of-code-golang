package year_2019

import (
	"fmt"
	"strconv"
)


func recurse_fuel(input_mass int) int {
	if input_mass < 3 {
		return 0
	}
	required_fuel := input_mass / 3
	required_fuel = required_fuel - 2
	if required_fuel < 0 {
		return 0
	}
	return required_fuel + recurse_fuel(required_fuel)
}


func day1(data [] string){
	var output int = 0
	for _, mass := range data {
			if mass == "" {
				continue
			}
			x, err := strconv.Atoi(mass)
			if err != nil {
				fmt.Printf("Error handling for string conversion of %s: %s\n", mass, err)
			}
			x = x / 3
			x = x - 2
			output += x
	}
	fmt.Printf("part1: %d\n", output)
	output = 0
	for _, mass := range data {
		if mass == "" {
			continue
		}
		x, err := strconv.Atoi(mass)
		if err != nil {
			fmt.Printf("Problem dude")
		}
		output += recurse_fuel(x)
	}
	fmt.Printf("part2: %d\n", output)
}
