package year_2019

import (
	"fmt"
	"github.com/drake-mer/advent-of-code-golang/cli"
)


func RunDay(day int){
	fmt.Printf("day %d\n", day)
	data := cli.DownloadInput(day, 2019) 
	switch day {
	case 1:
		day1(data)

	default:
		panic("not implemented")
	}
}
