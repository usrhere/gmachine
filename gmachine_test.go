package gmachine_test

import (
	"os"
	"slices"
	"testing"

	"gmachine"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Parallel()
	m := gmachine.New()
	var wantP uint16 = 0
	var wantA byte = 0
	if wantP != m.PC {
		t.Errorf("want initial P value %d, got %d", wantP, m.PC)
	}
	if wantA != m.A {
		t.Errorf("want initial A value %d, got %d", wantA, m.A)
	}
	var wantMemValue byte = 0
	gotMemValue := m.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHALT(t *testing.T) {
	m := gmachine.New()
	m.RunProgram([]byte{gmachine.OpHALT})
	if m.PC != 1 {
		t.Errorf("want P == 1, got %d", m.PC)
	}
}

func TestNOOP(t *testing.T) {
	m := gmachine.New()
	m.RunProgram([]byte{
		gmachine.OpNOOP,
		gmachine.OpHALT,
	})
	if m.PC != 2 {
		t.Errorf("want P == 2, got %d", m.PC)
	}
}

func TestINCA(t *testing.T) {
	m := gmachine.New()
	m.RunProgram([]byte{gmachine.OpINCA})
	if m.A != 1 {
		t.Errorf("want A == 1, got %d", m.A)
	}
}

func TestDECA(t *testing.T) {
	m := gmachine.New()
	m.A = 2
	m.RunProgram([]byte{gmachine.OpDECA})
	if m.A != 1 {
		t.Errorf("want A == 1, got %d", m.A)
	}
}

func TestSubtraction(t *testing.T) {
	var operand, wantA byte
	operand = 3
	wantA = 1
	m := gmachine.New()
	m.RunProgram([]byte{
		gmachine.OpLDA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if m.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, m.A)
	}
	operand = 4
	wantA = 2
	m.PC = 0
	m.RunProgram([]byte{
		gmachine.OpLDA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if m.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, m.A)
	}
	operand = 100
	wantA = 98
	m.PC = 0
	m.RunProgram([]byte{
		gmachine.OpLDA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if m.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, m.A)
	}
}

func TestSETA(t *testing.T) {
	var operand byte = 99
	m := gmachine.New()
	m.RunProgram([]byte{
		gmachine.OpLDA,
		operand,
		gmachine.OpHALT,
	})
	if m.PC != 3 {
		t.Errorf("want P == 2, got %d", m.PC)
	}
	if m.A != operand {
		t.Errorf("want A == %d, got %d", operand, m.A)
	}

}

func TestB(t *testing.T) {
	var operand byte = 99
	m := gmachine.New()
	m.RunProgram([]byte{
		gmachine.OpLDB,
		operand,
		gmachine.OpINCB,
		gmachine.OpINCB,
		gmachine.OpDECB,
	})
	if m.B != 100 {
		t.Errorf("want B == %d, got %d", 100, m.B)
	}
}

func TestReadInstructionsFromFile(t *testing.T) {
	m := gmachine.New()
	f, err := os.ReadFile("test/f1.bin")
	if err != nil {
		t.Error("Error reading file")
	}
	m.RunProgram(f)
	wantA := 81
	if m.A != byte(wantA) {
		t.Errorf("want A == %d, got %d", wantA, m.A)
	}
}

func TestPCIsIncrementedForEachInstruction(t *testing.T) {
	m := gmachine.New()
	program := []byte{
		gmachine.OpNOOP,
		gmachine.OpNOOP,
	}
	m.LoadToMemory(program)
	m.Step()
	wantPC := 1
	if m.PC != uint16(wantPC) {
		t.Errorf("want PC == %d, got %d", wantPC, m.PC)
	}
}

func TestAssemblerAndDisassemblerAreConsistent(t *testing.T) {
	assembly := `lda 50
inc a
inc a
dec a
ldb 100
dec b
dec b
inc b
halt
`
	objects := gmachine.Assemble([]byte(assembly))
	source := gmachine.DisassembleProgram(objects)
	if !slices.Equal([]byte(assembly), source) {
		t.Errorf("slices are not equal, want %v , got %s", assembly, source)
		t.Error(cmp.Diff(assembly, string(source)))
	}
}
