package main

import (
	"flag"
	"fmt"
)

var (
	version = "dev" // set via -ldflags at build time
)

func main() {
	name := flag.String("name", "World", "name to print")
	showVersion := flag.Bool("version", false, "print version and exit")
	quiet := flag.Bool("quiet", false, "suppress output")

	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	if *quiet {
		return
	}

	fmt.Printf("Hello, %s\n", *name)
}
