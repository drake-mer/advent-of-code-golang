package year_2019

import (
	"fmt"
	"strconv"
	"strings"
)


func runComputer(input_list [] int, noun int, verb int) int {
	integer_list := make([]int, len(input_list))
	copy(integer_list, input_list)
	integer_list[1] = noun
	integer_list[2] = verb
	max_length := len(integer_list)
	must_break := false
	for k := 0; k < max_length; k=k+4 {
		if must_break {
			break
		}
		instruction := integer_list[k]
		operand_1 := integer_list[integer_list[k+1]]
		operand_2 := integer_list[integer_list[k+2]]
		result_storage_location := integer_list[k+3]
		switch instruction {
		case 1:
			res := operand_1 + operand_2
			integer_list[result_storage_location] = res
		case 2:
			res := operand_1 * operand_2
			integer_list[result_storage_location] = res
		case 99:
			must_break = true
		default:
			panic("not supposed to happen")
		}

	}
	return integer_list[0]
}



func day02(data [] string) (int, int) {
	var input_data []string = strings.Split(data[0], ",")
	var integer_list = make([]int, len(input_data))
	for position, input_string := range(input_data) {
		res, err := strconv.Atoi(input_string)
		if err != nil {
			panic(fmt.Sprintf("could not convert string %s into number: %s", input_string, err))
		}
		integer_list[position] = res
	}
	part1 := runComputer(integer_list, 12, 2)
	must_break := false
	var k int
	var l int
	for k = 0; k < 100; k++ {
		for l = 0; l < 100; l++ {
			resp := runComputer(integer_list, k, l)
			if resp == 19690720 {
				must_break = true
				break
			}
		}
		if must_break {
			break
		}
	}
	return part1, (100 * k + l)
}