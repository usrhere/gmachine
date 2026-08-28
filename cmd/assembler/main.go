package main

import (
	"flag"
	"os"

	"gmachine"
)


func main() {
	var (
		f []byte
		err error
	)
	file := flag.String("i", "", "input, must be assembly file")
	output := flag.String("o", "", "output file")
	flag.Parse()
	if *file != "" {
		f, err = os.ReadFile(*file)
		if err != nil {
			panic(err)
		}
	}
	objects := gmachine.Assemble(f)
	os.WriteFile(*output, objects, 0644)
}

//func main() {
//	var (
//		f   []byte
//		err error
//	)
//	file := flag.String("i", "", "input, must be assembly file")
//	output := flag.String("o", "", "output file")
//	flag.Parse()
//	if *file != "" {
//		f, err = os.ReadFile(*file)
//		if err != nil {
//			panic(err)
//		}
//	}
//	var instructions []byte
//	opcodes := strings.Split(string(f), "\n")
//	for i := 0; i < len(opcodes); i++ {
//		fmt.Println(opcodes[i])
//		switch {
//		case opcodes[i] == "halt":
//			instructions = append(instructions, gmachine.OpHALT)
//		case opcodes[i] == "inc a":
//			instructions = append(instructions, gmachine.OpINCA)
//		case opcodes[i] == "dec a":
//			instructions = append(instructions, gmachine.OpDECA)
//		case strings.HasPrefix(opcodes[i], "lda "):
//			instructions = append(instructions, gmachine.OpLDA)
//			number, err := strconv.Atoi(strings.Split(opcodes[i], " ")[1])
//			if err != nil {
//				panic(err)
//			}
//			instructions = append(instructions, byte(number))
//		case opcodes[i] == "inc b":
//			instructions = append(instructions, gmachine.OpINCB)
//		case opcodes[i] == "dec b":
//			instructions = append(instructions, gmachine.OpDECB)
//		case strings.HasPrefix(opcodes[i], "ldb "):
//			instructions = append(instructions, gmachine.OpLDB)
//			number, err := strconv.Atoi(strings.Split(opcodes[i], " ")[1])
//			if err != nil {
//				panic(err)
//			}
//			instructions = append(instructions, byte(number))
//		}
//	}
//	fmt.Println(instructions)
//	os.WriteFile(*output, instructions, 0644)
//}
