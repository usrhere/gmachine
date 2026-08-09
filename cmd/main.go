package main

import (
	"gmachine"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Pass input file as argument.")
	}
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	m := gmachine.New()
	m.RunProgram(f)
}

