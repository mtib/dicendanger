package rooms

import (
	"github.com/mtib/dicendanger/rooms/utf"
)

func HorizontalLightWall(length int) Object {
	return hwall(length, utf.LightHorizontal)
}
func HorizontalHeavyWall(length int) Object {
	return hwall(length, utf.HeavyHorizontal)
}
func HorizontalDoubleWall(length int) Object {
	return hwall(length, utf.DoubleHorizontal)
}

func hwall(length int, wall rune) Object {
	if length <= 0 {
		return Object{""}
	}
	var res string
	for i := 0; i < length; i++ {
		res += string(wall)
	}
	return Object{res}
}

func VerticalLightWall(length int) Object {
	return vwall(length, utf.LightVertical)
}
func VerticalHeavyWall(length int) Object {
	return vwall(length, utf.HeavyVertical)
}
func VerticalDoubleWall(length int) Object {
	return vwall(length, utf.DoubleVertical)
}

func vwall(length int, wall rune) Object {
	if length <= 0 {
		return Object{}
	}
	var vw Object
	for i := 0; i < length; i++ {
		vw = append(vw, string(wall))
	}
	return vw
}
