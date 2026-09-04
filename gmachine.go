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
	OpJMP          = 0xf7
)

type Machine struct {
	PC          uint16 // Program counter
	A           byte   // Accumulator A
	B           byte   // Accumulator B
	Memory      [65536]byte
	Debug       bool
	Interactive bool
	Monitor     bool
	Assemble    bool
}

func instructionLength(opcode byte) int {
	switch opcode {
	case OpLDA, OpLDB:
		return 2
	case OpJMP:
		return 3
	default:
		return 1
	}
}
func convertLEBytesToUint16(low, high byte) uint16 {
	return uint16(high)*256 + uint16(low)
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
	case OpJMP:
		source = "jmp"
	}
	instructionLength := instructionLength(objects[0])
	switch instructionLength {
	case 2:
		source += fmt.Sprintf(" %d", objects[1])
	case 3:
		source += fmt.Sprintf(" %d", convertLEBytesToUint16(objects[1], objects[2]))
	}
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
	case OpJMP:
		low := m.Memory[m.PC+1]
		high := m.Memory[m.PC+2]
		m.PC = uint16(high)*256 + uint16(low)
		return false
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
		case OpJMP:
			next = fmt.Sprintf("jmp, %d", uint16(m.Memory[m.PC+1])+uint16(m.Memory[m.PC+2])*256)
		}
		fmt.Printf("A: %-3d (0x%-3x) | B: %-3d (0x%-3x) | PC: %-3d | Next instruction: %s", m.A, m.A, m.B, m.B, m.PC, next)
		if m.Step() {
			fmt.Println("\nHalted")
		}
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}
}
func (m *Machine) monitor() {
	var address uint16
	fmt.Println("Memory dump")
	fmt.Print("Enter address (or press enter to start from 0): ")
	fmt.Scanln(&address)
	for i := address; int(i)+16 < len(m.Memory); i += 16 {
		fmt.Printf("%04X: % x", i, m.Memory[i:i+16])
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}

}

func (m *Machine) Run() {
	if m.Debug {
		m.debug()
	} else if m.Monitor {
		m.monitor()
	} else if m.Assemble {
		Assemble(m.Memory[:])
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
		case strings.HasPrefix(opcodes[i], "jmp "):
			objects = append(objects, OpJMP)
			number, err := strconv.Atoi(strings.Split(opcodes[i], " ")[1])
			if err != nil {
				panic(err)
			}
			high := number / 256
			low := number % 256
			objects = append(objects, byte(low), byte(high))
		}
	}
	// fmt.Println(objects)
	return objects
}

// TODO: try to write conditional loop
