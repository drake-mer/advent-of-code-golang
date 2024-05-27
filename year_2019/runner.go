package year_2019

import (
	"fmt"
	"github.com/drake-mer/advent-of-code-golang/cli"
)


func RunDay(day int, download bool){
	fmt.Printf("day %d\n", day)
	data := cli.GetInputWithCaching(day, 2019, download)
	var part1 int = 0
	var part2 int = 0
	switch day {
	case 1:
		part1, part2 = day01(data)
	case 2:
		part1, part2 = day02(data)
	case 3:
		part1, part2 = day03(data)
	case 4:
		part1, part2 = day04(data)
	default:
		panic("not implemented")
	}

	fmt.Printf("part1: %d\npart2: %d\n", part1, part2)
}
