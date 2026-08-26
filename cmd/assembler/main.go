package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DefaultMemSize = 1
	OpHALT         = 0
	OpNOOP         = 1
	OpINCA         = 0x30
	OpDECA         = 0x40
	OpLDA          = 0x10
	OpINCB         = 0x31
	OpDECB         = 0x41
	OpLDB          = 0x11
)

func main() {
	var (
		f   []byte
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
	var instructions []byte
	opcodes := strings.Split(string(f), "\n")
	for i := 0; i < len(opcodes); i++ {
		fmt.Println(opcodes[i])
		switch {
		case opcodes[i] == "halt":
			instructions = append(instructions, OpHALT)
		case opcodes[i] == "inc a":
			instructions = append(instructions, OpINCA)
		case opcodes[i] == "dec a":
			instructions = append(instructions, OpDECA)
		case strings.HasPrefix(opcodes[i], "lda "):
			instructions = append(instructions, OpLDA)
			number, err := strconv.Atoi(strings.Split(opcodes[i], " ")[1])
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, byte(number))
		case opcodes[i] == "inc b":
			instructions = append(instructions, OpINCB)
		case opcodes[i] == "dec b":
			instructions = append(instructions, OpDECB)
		case strings.HasPrefix(opcodes[i], "ldb "):
			instructions = append(instructions, OpLDB)
			number, err := strconv.Atoi(strings.Split(opcodes[i], " ")[1])
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, byte(number))
		}
	}
	fmt.Println(instructions)
	os.WriteFile(*output, instructions, 0644)
}
