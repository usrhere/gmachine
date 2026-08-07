package gmachine_test

import (
	"testing"

	"gmachine"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantP byte = 0
	var wantA byte = 0
	if wantP != g.PC {
		t.Errorf("want initial P value %d, got %d", wantP, g.PC)
	}
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantMemValue byte = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHALT(t *testing.T) {
	g := gmachine.New()
	g.RunProgram([]byte{gmachine.OpHALT})
	if g.PC != 1 {
		t.Errorf("want P == 1, got %d", g.PC)
	}
}

func TestNOOP(t *testing.T) {
	g := gmachine.New()
	g.RunProgram([]byte{
		gmachine.OpNOOP,
		gmachine.OpHALT,
	})
	if g.PC != 2 {
		t.Errorf("want P == 2, got %d", g.PC)
	}
}

func TestINCA(t *testing.T) {
	g := gmachine.New()
	g.RunProgram([]byte{gmachine.OpINCA})
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestDECA(t *testing.T) {
	g := gmachine.New()
	g.A = 2
	g.RunProgram([]byte{gmachine.OpDECA})
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestSubtraction(t *testing.T) {
	var operand, wantA byte
	operand = 3
	wantA = 1
	g := gmachine.New()
	g.RunProgram([]byte{
		gmachine.OpLDA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if g.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
	operand = 4
	wantA = 2
	g.PC = 0
	g.RunProgram([]byte{
		gmachine.OpLDA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if g.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
	operand = 100
	wantA = 98
	g.PC = 0
	g.RunProgram([]byte{
		gmachine.OpLDA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if g.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
}

func TestSETA(t *testing.T) {
	var operand byte = 99
	g := gmachine.New()
	g.RunProgram([]byte{
		gmachine.OpLDA,
		operand,
	})
	if g.PC != 2 {
		t.Errorf("want P == 2, got %d", g.PC)
	}
	if g.A != operand {
		t.Errorf("want A == %d, got %d", operand, g.A)
	}

}

func TestB(t *testing.T) {
	var operand byte = 99
	g := gmachine.New()
	g.RunProgram([]byte{
		gmachine.OpLDB,
		operand,
		gmachine.OpINCB,
		gmachine.OpINCB,
		gmachine.OpDECB,
	})
	if g.B != 100 {
		t.Errorf("want B == %d, got %d", 100, g.B)
	}
}
