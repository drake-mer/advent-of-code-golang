
package years

import (
	"github.com/drake-mer/advent-of-code-golang/year_2019"
	"github.com/drake-mer/advent-of-code-golang/year_2023"
)


func RunProgram(year int, day int) {
	switch year {
	case 2019:
		year_2019.RunDay(day)
	case 2020:
		panic("not implemented")
	case 2021:
		panic("not implemented")
	case 2022:
		panic("not implemented")
	case 2023:
		year_2023.RunDay(day)
	default:
		panic("not implemented")
	}
}