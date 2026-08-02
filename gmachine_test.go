package gmachine_test

import (
	"testing"

	"gmachine"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantP gmachine.Word = 0
	var wantA gmachine.Word = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantMemValue gmachine.Word = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHALT(t *testing.T) {
	g := gmachine.New()
	g.RunProgram([]gmachine.Word{gmachine.OpHALT})
	if g.P != 1 {
		t.Errorf("want P == 1, got %d", g.P)
	}
}

func TestNOOP(t *testing.T) {
	g := gmachine.New()
	g.RunProgram([]gmachine.Word{
		gmachine.OpNOOP,
		gmachine.OpHALT,
	})
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}

func TestINCA(t *testing.T) {
	g := gmachine.New()
	g.RunProgram([]gmachine.Word{gmachine.OpINCA})
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestDECA(t *testing.T) {
	g := gmachine.New()
	g.A = 2
	g.RunProgram([]gmachine.Word{gmachine.OpDECA})
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestSubtraction(t *testing.T) {
	var operand, wantA gmachine.Word
	operand = 3
	wantA = 1
	g := gmachine.New()
	g.RunProgram([]gmachine.Word{
		gmachine.OpSETA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if g.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
	operand = 4
	wantA = 2
	g.P = 0
	g.RunProgram([]gmachine.Word{
		gmachine.OpSETA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if g.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
	operand = 100
	wantA = 98
	g.P = 0
	g.RunProgram([]gmachine.Word{
		gmachine.OpSETA,
		operand,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	if g.A != wantA {
		t.Errorf("want A == %d, got %d", wantA, g.A)
	}
}

func TestSETA(t *testing.T) {
	var operand gmachine.Word = 99
	g := gmachine.New()
	g.RunProgram([]gmachine.Word{
		gmachine.OpSETA,
		operand,
	})
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
	if g.A != operand {
		t.Errorf("want A == %d, got %d", operand, g.A)
	}

}
