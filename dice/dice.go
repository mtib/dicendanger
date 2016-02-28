package dice

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Regs ist der Regex fuer Wuerfel
	Regs = "([0-9]*)[dwDW]([0-9]*)(([+-])([0-9]*)){0,1}"
	// Wurfregex ist der compilierte Regex fuer Wuerfel
	Wurfregex, _ = regexp.Compile(Regs)
)

// CompileMultiple uebersetzt einen String durch regex in Wuerfel Array
func CompileMultiple(d string) WurfelClass {
	data := Wurfregex.FindAllStringSubmatch(d, -1)
	wstrings := strings.Split(d, " ")
	res := make([]Wurfel, len(data))
	for k, v := range data {
		num, _ := strconv.ParseInt(v[1], 10, 64)
		typ, _ := strconv.ParseInt(v[2], 10, 64)
		mvar, _ := strconv.ParseInt(v[5], 10, 64)
		var sub bool
		if string(wstrings[k][0]) == "s" || string(wstrings[k][0]) == "-" {
			sub = true
		} else {
			sub = false
		}
		res[k] = Wurfel{v[0], num, typ, v[3], v[4], mvar, sub}
	}
	return res
}

// CompileOnce uebersetzt einen String durch regex in einen Wuerfel
func CompileOnce(d string) Wurfel {
	return CompileMultiple(d)[0]
}

// Cwo wandelt einen String in einen! Wuerfel um
func Cwo(d string) Wurfel {
	return CompileOnce(d)
}

// RollString wuerfelt einen String, nachdem er ihn in
// mehrere Wuerfel compiliert.
func RollString(d string) int64 {
	return CompileMultiple(d).RollAll()
}

func (w Wurfel) String() string {
	var pref string
	if w.Sub {
		pref = "-"
	} else {
		pref = "+"
	}
	if w.Mvar == 0 {
		return fmt.Sprintf("%s%vD%v", pref, w.Num, w.Typ)
	}
	return fmt.Sprintf("%s%vD%v%s%v", pref, w.Num, w.Typ, w.Mtyp, w.Mvar)
}

// Roll wuerfelt den Wuerfel
func (w Wurfel) Roll() int64 {
	var ans int64
	for i := int64(0); i < w.Num; i++ {
		ans = ans + rand.Int63n(w.Typ) + 1
	}
	switch w.Mtyp {
	case "+":
		ans = ans + w.Mvar
		break
	case "-":
		ans = ans + w.Mvar
		break
	}
	if w.Sub {
		ans *= -1
	}
	return ans
}

// RollAll wuerfelt alle Wuerfel der Klasse
func (wa WurfelClass) RollAll() int64 {
	var sum int64
	for _, v := range wa {
		roll := v.Roll()
		fmt.Printf("%v --> %d\n", v, roll)
		sum += roll
	}
	return sum
}

func (wa WurfelClass) String() string {
	return fmt.Sprintf("%v [%d:%d]", []Wurfel(wa), wa.Min(), wa.Max())
}

// MinMax gibt ein Array zurueck mit dem Mindest- und
// Maximalwert des Wuerfelwurfes.
func (wa WurfelClass) MinMax() [2]int64 {
	var min, max int64
	for _, v := range wa {
		min += SubMult(v) * (v.Num + ModAdd(v))
		if v.Sub {
			max--
		} else {
			max += v.Num*v.Typ + ModAdd(v)
		}
	}
	return [2]int64{min, max}
}

// Min gibt den kleinstmoeglichen Wert zurueck
func (wa WurfelClass) Min() int64 {
	return wa.MinMax()[0]
}

// Max gibt den groesstmoeglichen Wert zurueck
func (wa WurfelClass) Max() int64 {
	return wa.MinMax()[1]
}

// ModAdd gibt den Modifikator nach Applikation des Vorzeichens
// zurueck.
func ModAdd(v Wurfel) int64 {
	switch v.Mtyp {
	case "-":
		return -1 * v.Mvar
	case "+":
		return v.Mvar
	}
	return 0
}

// SubMult gibt 1 zurueck wenn der Wurf positiv ist, und
// -1 wenn der Wurf negativ ist.
func SubMult(v Wurfel) int64 {
	if v.Sub {
		return -1
	}
	return 1
}

// CmdW wandelt ein string[] zu einem Wuerfel[] um.
func CmdW(cmdarr []string) {
	w := WurfelClass(CompileMultiple(strings.Join(cmdarr, " ")))
	fmt.Println("Würfel:", w)
	fmt.Printf("Ergebnis: %d\n", w.RollAll())
}
