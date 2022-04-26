package _2shortDeclarations

import "testing"

func TestShortDeclaration(t *testing.T) {
	exp := 23
	if act := shortDeclaration(); exp != act {
		t.Errorf("shortdeclaration() = %q, want %q", act, exp)
	}

}
