package main

import (
	"fmt"
	"github.com/drake-mer/advent-of-code-golang/cli"
	"github.com/drake-mer/advent-of-code-golang/years"
)


func main() {
	fmt.Printf("hello\n")
	cli.CliHello()
	years.RunProgram(cli.YEAR, cli.DAY)
}
