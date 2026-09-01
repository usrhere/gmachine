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
	disassemble := flag.Bool("s", false, "disassemble given file")
	monitor := flag.Bool("m", false, "run in monitoring mode")
	flag.Parse()
	if *file != "" {
		f, err = os.ReadFile(*file)
		if err != nil {
			panic(err)
		}
	}
	if *disassemble {
		gmachine.DisassembleProgram(f)
		return
	}
	m := gmachine.New()
	if *debug {
		m.Debug = true
	}
	if *interactive {
		m.Interactive = true
	}
	if *monitor {
		m.Monitor = true
	}
	m.RunProgram(f)
}
