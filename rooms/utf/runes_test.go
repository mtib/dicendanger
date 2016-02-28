package utf

import (
	"fmt"
	"testing"
)

func TestRunes(t *testing.T) {
	for index := 0x2500; index <= 0x257f; index++ {
		fmt.Printf("%c", rune(index))
	}
	fmt.Print("\n")
}
