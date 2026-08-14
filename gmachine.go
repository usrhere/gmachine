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
	PC     uint16 // Program counter
	A      byte   // Accumulator A
	B      byte   // Accumulator B
	Memory [65536]byte
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

func (m *Machine) Run() {
	for {
		halt := m.Step()
		if halt {
			break
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
