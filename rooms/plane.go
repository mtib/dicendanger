package rooms

import (
	"fmt"
	"github.com/mtib/dicendanger/rooms/utf"
	"strings"
)

type Plane [][]rune

func (p *Plane) AddObject(x, y int, obj Object) {
	for py := y; py < y+len(obj); py++ {
		runelist := []rune(obj[py-y])
		for ix, r := range runelist {
			if !(x+ix >= p.Width() || py >= p.Height()) {
				oldrune := (*p)[py][x+ix]
				(*p)[py][x+ix] = r
				if oldrune != ' ' && oldrune != r {
					(*p)[py][x+ix] = utf.DoublePlus
				}
			}
		}
	}
	p.update()
}

func (p *Plane) update() {
	for x := 0; x < p.Width()-1; x++ {
		// check Down & Right y=0 || y=max
	}
	// Check [max, 0] for Down & Left
	// Check [max, max] for Up & Left
	for y := 1; y < p.Height()-1; y++ {
		for x := 1; x < p.Width()-1; x++ {
			// Check [inner] for Up, Down, Left, Right
		}
	}
}

func (p *Plane) Print() {
	planes := p.String()
	planearr := strings.Split(planes, "\n")
	borderwall := HorizontalDoubleWall(p.Width())[0]
	var bordertopwall string
	for k, v := range []rune(planearr[0]) {
		if v == ' ' || v == utf.DoubleHorizontal {
			bordertopwall += string([]rune(borderwall)[k])
		} else {
			bordertopwall += string(utf.DoubleHorizontalDown)
		}
	}
	fmt.Printf("%c%s%c\n", utf.DoubleDownRight, bordertopwall, utf.DoubleDownLeft)
	for _, k := range planearr {
		left := utf.DoubleVertical
		right := utf.DoubleVertical
		runeline := []rune(k)
		if runeline[0] != ' ' && runeline[0] != utf.DoubleVertical {
			left = utf.DoubleVerticalRight
		}
		if runeline[p.Width()-1] != ' ' && runeline[p.Width()-1] != utf.DoubleVertical {
			right = utf.DoubleVerticalLeft
		}
		fmt.Printf("%c%s%c\n", left, k, right)
	}
	var borderbottomwall string
	for k, v := range []rune(planearr[len(planearr)-1]) {
		if v == ' ' || v == utf.DoubleHorizontal {
			borderbottomwall += string([]rune(borderwall)[k])
		} else {
			borderbottomwall += string(utf.DoubleHorizontalUp)
		}
	}
	fmt.Printf("%c%s%c\n", utf.DoubleUpRight, borderbottomwall, utf.DoubleUpLeft)
}

func (p *Plane) String() string {
	var res string
	for y := 0; y < len(*p); y++ {
		for _, r := range (*p)[y] {
			res += string(r)
		}
		res += "\n"
	}
	return res[:len(res)-1]
}

func (p *Plane) New(width, height int) {
	*p = make([][]rune, height)
	for k := range *p {
		(*p)[k] = make([]rune, width)
		for k2 := range (*p)[k] {
			(*p)[k][k2] = ' '
		}
	}
}

func (p Plane) Width() int {
	return len(p[0])
}

func (p Plane) Height() int {
	return len(p)
}
