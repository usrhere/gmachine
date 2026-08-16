package main

import (
	"flag"
	"gmachine"
	"os"
)

func main() {
	var f []byte
	var err error
	debug := flag.Bool("d", false, "enable debug mode")
	file := flag.String("f", "", "file to process")
	interactive := flag.Bool("i", false, "enable interactive mode")
	flag.Parse()
	if *file != "" {
		f, err = os.ReadFile(*file)
		if err != nil {
			panic(err)
		}
	}
	m := gmachine.New()
	if *debug {
		m.Debug = true
	}
	if *interactive {
		m.Interactive = true
	}
	m.RunProgram(f)
}

