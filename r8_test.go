package r8_test

import (
	"testing"

	"github.com/bitfield/go-r8"
)

func TestNewInitialisesCPU(t *testing.T) {
	t.Parallel()
	cpu := r8.New()
	if cpu.PC != 0 {
		t.Errorf("after New, want pc == 0, got %d", cpu.PC)
	}
	got := cpu.Memory[0]
	if got != 0 {
		t.Errorf("after New, want Memory[0] == 0, got %d", got)
	}
}

// Uncomment this test once the previous test passes!
// func TestNopInstructionIncrementsPC(t *testing.T) {
// 	t.Parallel()
// 	cpu := r8.New()
// 	cpu.Memory[0] = r8.OpNOP
// 	cpu.Step()
// 	if cpu.PC != 1 {
// 		t.Errorf("want pc == 1, got %d", cpu.PC)
// 	}
// }
