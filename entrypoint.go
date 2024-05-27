package main

import (
	"github.com/drake-mer/advent-of-code-golang/cli"
	"github.com/drake-mer/advent-of-code-golang/years"
)


func main() {
	cli.CliHello()
	years.RunProgram(cli.YEAR, cli.DAY, cli.DOWNLOAD)
}
