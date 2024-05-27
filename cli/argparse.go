package cli

import (
	"flag"
	"fmt"
	"time"
)

var YEAR int
var DAY int
var DOWNLOAD bool

func CliHello() {
	fmt.Println("Parsing argument parameters")
	now := time.Now()
	default_year := now.Year()
	default_day := now.Day()
	default_download := false
	flag.IntVar(&YEAR, "y", default_year, "The year of the exercise")
	flag.IntVar(&DAY, "d", default_day, "The day of the exercise")
	flag.BoolVar(&DOWNLOAD, "dl", default_download, "Force downloading on HTTP" )
	flag.Parse()
	fmt.Printf("Day %d, Year %d\n", DAY, YEAR)
}
