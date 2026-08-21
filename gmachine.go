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

func translateInstructionToString(opcode []byte) (string, []byte) {
	var instruction string
	switch opcode[0] {
	case OpHALT:
		instruction = "halt"
	case OpINCA:
		instruction = "inc a"
	case OpDECA:
		instruction = "dec a"
	case OpLDA:
		instruction = "ld a"
	case OpINCB:
		instruction = "inc b"
	case OpDECB:
		instruction = "dec b"
	case OpLDB:
	}
	instructionLength := instructionLength(opcode[0])
	if instructionLength != 1 {
		instruction += strconv.Itoa(instructionLength)
	}
	opcode = opcode[instructionLength:]
	return instruction, opcode
}

func DisassembleProgram(program []byte) {
	var instruction string
	for range program {
		instruction, program = translateInstructionToString(program)
		fmt.Printf("%s\n", instruction)
		if len(program) == 0 {
			break
		}
	}
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

// TODO: Write assembly code and produce a binary file. Without operands in the beginning. It should take assembly file as input and produce the binary.
