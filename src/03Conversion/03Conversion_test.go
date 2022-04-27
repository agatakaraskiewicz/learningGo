package _3Conversion

import (
	"fmt"
	"testing"
)

func TestConversion(t *testing.T) {
	expx := 4
	expMilky := 2
	actx, actMilky, actKnoppers := conversion()

	fmt.Println(actKnoppers)
	y := fmt.Sprintf("Knoppers type is %T", actKnoppers)

	if y != "Knoppers type is int" {
		t.Errorf("It was wrong: %q", y)
	}

	if expx != actx {
		t.Errorf("Conversions() = %q, want %q", actx, expx)
	}

	if expMilky != actMilky {
		t.Errorf("Conversion()2 = %q, want %q", actMilky, expMilky)
	}
}
