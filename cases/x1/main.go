package main

import (
	"fmt"
	"math"
)

func main() {
	scale := 8
	v := int64(123456)

	sig := false
	if v < 0 {
		v = -v
		sig = true
	}
	s := fmt.Sprintf(fmt.Sprintf("%%d.%%0%dd", scale), v/int64(math.Pow10(scale)), v%int64(math.Pow10(scale)))
	if sig {
		s = fmt.Sprintf("-%s", s)
	}
	fmt.Println(s)
}
