// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

import (
	"bufio"
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

type Machine struct {
	PC          uint16 // Program counter
	A           byte   // Accumulator A
	B           byte   // Accumulator B
	Memory      [65536]byte
	Debug       bool
	Interactive bool
}

func instructionLength(opcode byte) int {
	if opcode == OpLDA || opcode == OpLDB {
		return 2
	}
	return 1
}

func translateObjectToSource(objects []byte) (string, []byte) {
	var source string
	switch objects[0] {
	case OpHALT:
		source = "halt"
	case OpINCA:
		source = "inc a"
	case OpDECA:
		source = "dec a"
	case OpLDA:
		source = "lda"
	case OpINCB:
		source = "inc b"
	case OpDECB:
		source = "dec b"
	case OpLDB:
		source = "ldb"
	}
	instructionLength := instructionLength(objects[0])
	if instructionLength != 1 {
		source += fmt.Sprintf(" %d", objects[1])
	}
	fmt.Println(source, ", len: ", instructionLength)
	objects = objects[instructionLength:]
	return source, objects
}

func DisassembleProgram(objects []byte) []byte {
	var source []byte
	var instruction string
	for range objects {
		instruction, objects = translateObjectToSource(objects)
		source = append(source, []byte(instruction)...)
		source = append(source, []byte("\n")...)
		//	fmt.Printf("Instruction: %s\n", instruction)
		if len(objects) == 0 {
			break
		}
	}
	return source
}

func (m *Machine) LoadToMemory(program []byte) {
	copy(m.Memory[:], program)
}

func (m *Machine) Step() bool {
	switch m.Memory[m.PC] {
	case OpHALT:
		m.PC++
		return true
	case OpINCA:
		m.A++
	case OpDECA:
		m.A--
	case OpLDA:
		m.A = m.Memory[m.PC+1]
		m.PC++
	case OpINCB:
		m.B++
	case OpDECB:
		m.B--
	case OpLDB:
		m.B = m.Memory[m.PC+1]
		m.PC++
	}
	m.PC++
	return false
}

func (m *Machine) debug() {
	var next string
	for {
		switch m.Memory[m.PC] {
		case OpHALT:
			next = "halt"
		case OpINCA:
			next = "inc a"
		case OpDECA:
			next = "dec a"
		case OpLDA:
			next = fmt.Sprintf("ld a, %d", m.Memory[m.PC+1])
		case OpINCB:
			next = "inc b"
		case OpDECB:
			next = "dec b"
		case OpLDB:
			next = fmt.Sprintf("ld b, %d", m.Memory[m.PC+1])
		}
		fmt.Printf("A: %-3d (0x%-3x) | B: %-3d (0x%-3x) | PC: %-3d | Next instruction: %s", m.A, m.A, m.B, m.B, m.PC, next)
		if m.Step() {
			fmt.Println("\nHalted")
		}
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}
}

func (m *Machine) Run() {
	if m.Debug {
		m.debug()
	} else if m.Interactive {
		var input []byte
		fmt.Printf("Enter Op codes in hex: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			for _, field := range fields {
				b, err := strconv.ParseUint(field, 0, 8)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%d (%s)\n", b, field)
				input = append(input, byte(b))
			}
		}
		m.LoadToMemory(input)
		m.debug()
	} else {
		for {
			halt := m.Step()
			if halt {
				break
			}
		}
	}
}

func (m *Machine) RunProgram(program []byte) {
	m.LoadToMemory(program)
	m.Run()
}

func New() Machine {
	m := Machine{
		PC: 0,
	}
	return m
}

func Assemble(f []byte) []byte {
	var objects []byte
	opcodes := strings.Split(string(f), "\n")
	for i := 0; i < len(opcodes); i++ {
		fmt.Println(opcodes[i])
		switch {
		case opcodes[i] == "halt":
			objects = append(objects, OpHALT)
		case opcodes[i] == "inc a":
			objects = append(objects, OpINCA)
		case opcodes[i] == "dec a":
			objects = append(objects, OpDECA)
		case strings.HasPrefix(opcodes[i], "lda "):
			objects = append(objects, OpLDA)
			number, err := strconv.Atoi(strings.Split(opcodes[i], " ")[1])
			if err != nil {
				panic(err)
			}
			objects = append(objects, byte(number))
		case opcodes[i] == "inc b":
			objects = append(objects, OpINCB)
		case opcodes[i] == "dec b":
			objects = append(objects, OpDECB)
		case strings.HasPrefix(opcodes[i], "ldb "):
			objects = append(objects, OpLDB)
			number, err := strconv.Atoi(strings.Split(opcodes[i], " ")[1])
			if err != nil {
				panic(err)
			}
			objects = append(objects, byte(number))
		}
	}
	fmt.Println(objects)
	return objects
}

// TODO: show content of the memory, user provides the address of the memory, like https://github.com/bitfield/rx82#using-the-monitor
