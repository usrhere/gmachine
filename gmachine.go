// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

type Word uint64

const (
	DefaultMemSize = 1
	OpHALT         = 0
	OpNOOP         = 1
	OpINCA         = 2
	OpDECA         = 3
	OpSETA         = 4
)

type Machine struct {
	P      Word // Program counter
	A      Word // Accumulator
	Memory []Word
}

func (m *Machine) Run() {
instructions:
	for i := 0; i < len(m.Memory); i++ {
		switch m.Memory[i] {
		case 0:
			break instructions
		case 2:
			m.A++
		case 3:
			m.A--
		case 4:
			i++
			m.A = m.Memory[i]
		}
	}
	m.P = Word(len(m.Memory))
}

func (m *Machine) RunProgram(program []Word) {
	m.Memory = program
	m.Run()
}

func New() Machine {
	m := Machine{
		P:      0,
		Memory: []Word{0},
	}
	return m
}
