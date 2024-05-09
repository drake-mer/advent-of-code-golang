package year_2019

import "fmt"


func RunDay(day int){
	fmt.Printf("day %d\n", day)
	switch day {
	case 1:
		day1()

	default:
		panic("not implemented")
	}
}
