// Package r8 emulates a simple CPU called the R8.
package r8

const OpNOP = 1

type CPU struct {
	PC     uint16
	Memory [65536]byte
}

func New() *CPU {
	return &CPU{}
}

func (cpu *CPU) Step() {
	// Over to you to implement `Step`!
}
