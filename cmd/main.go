package main

import (
//	"flag"
	"gmachine"
	"os"
	"fmt"
)

//func main() {
//	var data []byte
//	var err error
//	debug := flag.Bool("d", false, "enable debug mode")
//	file := flag.String("f", "", "file to process")
//	interactive := flag.Bool("i", false, "enable interactive mode")
//	assemble := flag.Bool("a", false, "assemble given file")
//	disassemble := flag.Bool("s", false, "disassemble given file")
//	monitor := flag.Bool("m", false, "run in monitoring mode")
//	flag.Parse()
//	if *file != "" {
//		data, err = os.ReadFile(*file)
//		if err != nil {
//			panic(err)
//		}
//	}
//	if *disassemble {
//		fmt.Printf("%s", gmachine.DisassembleProgram(data))
//		return
//	}
//	if *assemble {
//		objects := gmachine.Assemble(data)
//		err := os.WriteFile("output.bin", objects, 0o0600)
//		if err != nil {
//			panic(err)
//		}
//		return
//	}
//	m := gmachine.New()
//	if *debug {
//		m.Debug = true
//	}
//	if *interactive {
//		m.Interactive = true
//	}
//	if *monitor {
//		m.Monitor = true
//	}
//	m.RunProgram(data)
//}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("help")
		return
	}
	switch os.Args[1] {
	case "mon":
		if len(os.Args) < 3 {
			fmt.Println("mon help")
		}
		m := gmachine.New()
		m.Debug = true
		data, err := os.ReadFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		m.RunProgram(data)
	}
}
