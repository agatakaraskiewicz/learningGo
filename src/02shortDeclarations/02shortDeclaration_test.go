package _2shortDeclarations

import "testing"

func TestShortDeclaration(t *testing.T) {
	expx := 23
	expy := 114
	actx, acty := shortDeclaration()

	if expx != actx {
		t.Errorf("shortdeclaration() = %q, want %q", actx, expx)
	}

	if expy != acty {
		t.Errorf("shortdeclaration()2 = %q, want %q", acty, expy)
	}
}
