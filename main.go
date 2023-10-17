package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const Version = "0.1.0"

var usageStr = `
Gonorrhea a command line interface for insert a virus.

Basic Commands (Beginner):
	version	 		Print the version information

Usage:
	gonorrhea [flags] [options]
`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func printVersion() {
	fmt.Printf("version: %s\n", Version)
	os.Exit(0)
}

func main() {

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		usage()
	}

	for _, v := range args {
		if v == "version" {
			printVersion()
		}

	}
}
