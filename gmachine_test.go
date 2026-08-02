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
	g.Run()
	if g.P != 1 {
		t.Errorf("want P == 1, got %d", g.P)
	}
}

func TestNOOP(t *testing.T) {
	g := gmachine.New()
	g.Memory[0] = gmachine.OpNOOP
	g.Memory = append(g.Memory, gmachine.OpHALT)
	g.Run()
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}

func TestINCA(t *testing.T) {
	g := gmachine.New()
	g.Memory[0] = gmachine.OpINCA
	g.Run()
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestDECA(t *testing.T) {
	g := gmachine.New()
	g.A = 2
	g.Memory[0] = gmachine.OpDECA
	g.Run()
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestSubtraction(t *testing.T) {
	g := gmachine.New()
	g.A = 3
	g.Memory[0] = gmachine.OpDECA
	g.Memory = append(g.Memory, gmachine.OpDECA)
	g.Run()
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}
