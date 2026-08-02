// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

type Word uint64

const (
	DefaultMemSize = 1
	OpHALT         = 0
	OpNOOP         = 1
	OpINCA         = 2
	OpDECA         = 3
)

type Machine struct {
	P      Word // Program counter
	A      Word // Accumulator
	Memory []Word
}

func (m *Machine) Run() {
instructions:
	for _, v := range m.Memory {
		m.P++
		switch v {
		case 0:
			break instructions
		case 2:
			m.A++
		case 3:
			m.A--
		}
	}
}

func New() Machine {
	m := Machine{
		P:      0,
		Memory: []Word{0},
	}
	return m
}
