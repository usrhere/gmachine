// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

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
	PC     byte // Program counter
	A      byte // Accumulator
	Memory []byte
}

func (m *Machine) Run() {
instructions:
	for i := 0; i < len(m.Memory); i++ {
		switch m.Memory[i] {
		case OpHALT:
			break instructions
		case OpINCA:
			m.A++
		case OpDECA:
			m.A--
		case OpLDA:
			i++
			m.A = m.Memory[i]
		}
	}
	m.PC = byte(len(m.Memory))
}

func (m *Machine) RunProgram(program []byte) {
	m.Memory = program
	m.Run()
}

func New() Machine {
	m := Machine{
		PC:     0,
		Memory: []byte{0},
	}
	return m
}

// TODO: write some code that reads a binary file and executes it
// E.g: 0x11 99 => converted to binary
