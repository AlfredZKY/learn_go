package other

import (
	"fmt"
	"testing"
)

func TestUnicode(t *testing.T) {
	str := "Go 爱好者 "
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
}
