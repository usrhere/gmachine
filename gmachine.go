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
	for {
		fmt.Printf("A: %-3d (0x%-3x) | B: %-3d (0x%-3x) | PC: %-3d", m.A, m.A, m.B, m.B, m.PC)
		m.Step()
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

// TODO: read binary file and execute it step by step
// print out the values of the registries, user should press enter after each and then execute next instruction
