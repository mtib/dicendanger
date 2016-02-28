package rooms

import (
	"fmt"
	"testing"
)

func TestHorizontals(t *testing.T) {
	var wallFArray map[string]func(int) Object
	wallFArray = map[string]func(int) Object{
		"Light":  HorizontalLightWall,
		"Heavy":  HorizontalHeavyWall,
		"Double": HorizontalDoubleWall,
	}
	for i := 0; i <= 10; i++ {
		for fn, fu := range wallFArray {
			wall := fu(i)
			if wall.NumRunes()/wall.Rows() != i {
				t.Fail()
			}
			if wall.Columns() != i {
				t.Fail()
			}
			fmt.Printf("%2d[%s]: %s\n", i, fn, wall)
		}
	}
}

func TestVerticals(t *testing.T) {
	var wallFArray map[string]func(int) Object
	wallFArray = map[string]func(int) Object{
		"Light":  VerticalLightWall,
		"Heavy":  VerticalHeavyWall,
		"Double": VerticalDoubleWall,
	}
	for i := 0; i <= 10; i++ {
		for fn, fu := range wallFArray {
			wall := fu(i)
			if wall.Rows() != i {
				t.Fail()
			}
			if len(wall) != i {
				t.Fail()
			}
			fmt.Printf("%2d[%s]: %s\n", i, fn, wall)
		}
	}
}
