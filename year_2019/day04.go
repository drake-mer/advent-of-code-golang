package year_2019

import (
	"fmt"
	"strconv"
	"strings"
)

func SmallGroup(data string) bool {
	var accumulator []int
	group_size := 1
	for pos, _ := range(data) {
		if pos == 5 {
			break
		}
		a := data[pos]
		b := data[pos + 1]
		if a == b {
			group_size ++
			if pos == 4 {
				accumulator = append(accumulator, group_size)
			}
		} else {
			accumulator = append(accumulator, group_size)
			group_size = 1
		}
	}
	for _, value := range(accumulator) {
		if value == 2 {
			return true
		}
	}
	return false
}

func ValidPassword(data string) bool {
	valid := false
	for pos, _ := range data {
		if pos == 5 {
			break
		}
		if data[pos] > data[pos + 1] {
			return false
		}
	}
	for pos, _ := range data {
		if pos == 5 {
			break
		}
		if data[pos] == data[pos + 1] {
			valid = true
		}
	}
	return valid
}


func day04(data [] string) (int, int) {
	data = strings.Split(data[0], "-")
	a, err1 := strconv.Atoi(data[0])
	b, err2 := strconv.Atoi(data[1])
	if err1 != nil || err2 != nil {
		fmt.Printf("Could not convert number, %e %e\n", err1, err2)
	}
	fmt.Printf("looking for passwords between %d and %d\n", a, b)
	total := 0
	for k := a; k<=b ; k++ {
		password := fmt.Sprintf("%d", k)
		if ValidPassword(password) {
			total++
		}
	}
	total2 := 0
	for k := a; k<=b ; k++ {
		password := fmt.Sprintf("%d", k)
		if ValidPassword(password) && SmallGroup(password) {
			fmt.Printf("valid password %s\n", password)
			total2++
		} else {
			fmt.Printf("non valid password %s\n", password)
		}
	}
	return total, total2
}