package main

import (
	"fmt"
)

func (e Enemy) String() string {
	return fmt.Sprintf("%s:\n\tTP: %d\n\tATD: %v\n\tATB: %v\n\tRKL: %v", e.Name, e.Health, *e.Attdice, *e.Attack, *e.Armor)
}

func (wa WurfelClass) String() string {
	return fmt.Sprintf("%v [%d:%d]", []Wurfel(wa), wa.min(), wa.max())
}

func (wa WurfelClass) minmax() [2]int64 {
	var min, max int64
	for _, v := range wa {
		min += subMult(v) * (v.Num + modAdd(v))
		if v.Sub {
			max--
		} else {
			max += v.Num*v.Typ + modAdd(v)
		}
	}
	return [2]int64{min, max}
}

func (wa WurfelClass) min() int64 {
	return wa.minmax()[0]
}

func (wa WurfelClass) max() int64 {
	return wa.minmax()[1]
}

func modAdd(v Wurfel) int64 {
	switch v.Mtyp {
	case "-":
		return -1 * v.Mvar
	case "+":
		return v.Mvar
	}
	return 0
}

func subMult(v Wurfel) int64 {
	if v.Sub {
		return -1
	}
	return 1
}

func (re RawEnemy) String() string {
	return fmt.Sprintf("[%d:%d]TP, [%d:%d]ATD, [%d]ATB, [%d]RKL", re.Hitdice.min(), re.Hitdice.max(), re.Attdice.min(), re.Attdice.max(), re.Attack, re.Armor)
}
