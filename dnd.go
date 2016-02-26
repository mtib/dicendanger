package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	regs         = "([0-9]*)[dwDW]([0-9]*)(([+-])([0-9]*)){0,1}"
	wurfregex, _ = regexp.Compile(regs)
	debug        = flag.Bool("debug", false, "Enable Debug Mode")
	inpt         = "2d5 3d10"
)

var aktiveEnemies Encounter

// RawEnemyList enthaelt alle Rohdaten fuer exportierbare Gegner
type RawEnemyList map[string]RawEnemy

func rawEnemyGen(hd, ad string, at, ar int64) RawEnemy {
	return RawEnemy{CompileMultiple(hd), CompileMultiple(ad), at, ar}
}

func newEnemy(name string) Enemy {
	re, ok := rawEnemies[name]
	if !ok {
		return Enemy{}
	}
	return Enemy{name, re.Hitdice.RollAll(), &re.Attdice, &re.Attack, &re.Armor}
}

// CompileMultiple uebersetzt einen String durch regex in Wuerfel Array
func CompileMultiple(d string) WurfelClass {
	data := wurfregex.FindAllStringSubmatch(d, -1)
	wstrings := strings.Split(d, " ")
	res := make([]Wurfel, len(data))
	dprint(data)
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

func cwo(d string) Wurfel {
	return CompileOnce(d)
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

func dprint(s interface{}) {
	if *debug {
		fmt.Println(s)
	}
}

func rollString(d string) int64 {
	return CompileMultiple(d).RollAll()
}

var last Enemy

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	stdin := bufio.NewScanner(os.Stdin)
	var stinput = func(desc string) string {
		fmt.Printf("%s> ", desc)
		stdin.Scan()
		return strings.TrimSpace(stdin.Text())
	}
	var itinput = func(desc string) int64 {
		val, _ := strconv.ParseInt(stinput(desc), 10, 64)
		return val
	}
CMDLOOP:
	for {
		cmd := strings.Split(stinput(""), " ")
		switch cmd[0] {
		case "w", "r":
			cmdW(cmd[1:])
			break
		case "export":
			if len(cmd) >= 2 {
				rawEnemies.jsonexport(cmd[1])
			} else {
				rawEnemies.jsonexport("dnd.json")
			}
		case "import":
			if len(cmd) >= 2 {
				rawEnemies.jsonimport(cmd[1])
			} else {
				rawEnemies.jsonimport("dnd.json")
			}
		case "raws":
			for k, v := range rawEnemies {
				fmt.Printf("%s: %v\n", k, v)
			}
		case "ne":
			if len(cmd) < 2 {
				break
			}
			_, ok := rawEnemies[cmd[1]]
			if ok {
				last = newEnemy(cmd[1])
				fmt.Println(last)
			} else {
				fmt.Println(cmd[1], "existiert nicht")
			}
		case "ce":
			fmt.Println("Erstelle Gegner: <name> <hd> <atd> <atb> <rkb>")
			name := stinput("Name")
			hitd := CompileMultiple(stinput("Lebens Würfel"))
			attd := CompileMultiple(stinput("Angriff Würfel"))
			attb := itinput("Angriff Bonus")
			rklb := itinput("Rüstungsbonus")
			rawEnemies[name] = RawEnemy{Hitdice: hitd, Attdice: attd, Attack: attb, Armor: rklb}
			last = Enemy{Name: name, Health: hitd.RollAll(), Armor: &rklb, Attack: &attb, Attdice: &attd}
		case "le", "list":
			fmt.Println("Gegner:")
			for k := range rawEnemies {
				fmt.Println("\t", k)
			}
			break
		case "q", "quit":
			break CMDLOOP
		case "last":
			if len(cmd) == 1 {
				cmd = []string{"last", "info"}
			}
			switch cmd[1] {
			case "att", "a", "attack":
				var attbon, attwf int64
				attbon = rand.Int63n(20) + 1 + *last.Attack
				if len(cmd) >= 3 {
					attwf, _ = strconv.ParseInt(cmd[2], 10, 64)
				} else {
					attwf = itinput("Rüstungsbonus")
				}
				if attbon > attwf {
					fmt.Printf("Gegner trifft\nSchaden: %v\n", last.Attdice.RollAll())
				} else {
					fmt.Println("Gegner trifft nicht!")
				}
			case "def", "d", "defend":
				var enhit int64
				if len(cmd) < 3 || cmd[2] == "" {
					enhit = itinput("Angriffswurf")
				} else {
					enhit, _ = strconv.ParseInt(cmd[2], 10, 64)
				}
				if enhit > *last.Armor {
					fmt.Println("Angriff erfolgreich!")
					dmg := itinput("Schaden")
					last.Health -= dmg
					fmt.Println(last.Name, last.Health, "TP")
					if last.Health <= 0 {
						fmt.Println(last.Name, "ist tot!")
					}
				} else {
					fmt.Println("Angriff fehlgeschlagen")
				}
			case "info", "i":
				fmt.Println(last)
			} // Last Switch
		case "raw":
			if len(cmd) < 2 {
				break
			}
			fmt.Println(rawEnemies[cmd[1]])
			break
		} // Cmd Switch
	} // For
} // Main

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

func cmdW(cmdarr []string) {
	w := WurfelClass(CompileMultiple(strings.Join(cmdarr, " ")))
	fmt.Println("Würfel:", w)
	fmt.Printf("Ergebnis: %d\n", w.RollAll())
}
