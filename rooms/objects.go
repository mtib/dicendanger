package rooms

import (
	"unicode/utf8"
)

type Object []string

func (o Object) NumRunes() int {
	var sum int
	for _, v := range o {
		sum += utf8.RuneCountInString(v)
	}
	return sum
}

func (o Object) Rows() int {
	return len(o)
}

func (o Object) Columns() int {
	var res int
	for _, v := range o {
		if curr := utf8.RuneCountInString(v); curr > res {
			res = curr
		}
	}
	return res
}
