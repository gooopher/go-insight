package question

import (
	"fmt"
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		Str  string
		Want string
	}{
		0: {
			Str:  "1duduaudwyqgduywqgi",
			Want: "igqwyudgqywduaudud1",
		},
	}
	for _, tt := range tests {
		str := Reversess(tt.Str)
		if !strings.EqualFold(str, tt.Want) {
			t.Errorf("verify fail")
		}

	}

	fmt.Println()
}
