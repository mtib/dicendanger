package main

import (
	"fmt"
	"github.com/mtib/dicendanger/dice"
	"math/rand"
	"strconv"
	"strings"
)

var (
	help = `Befehle:
	w <Würfel, ...> : würfeln und Summe berechnen
	r <Würfel, ...> : Alias von 'w'
	list            : oder 'le' Liste von Gegnertypen
	new <Gegnertyp> : erstellt Gegner im Index last ('ne', 'neu')
	raws            : Liste der Rohdaten zu Gegnertypen
	raw <Gegnertyp> : Rohdaten zu Gegnertyp
	last            : Infos zum aktiven Gegner
	last att        : Aktiven Gegner angreifen lassen
	last def        : Aktiven Gegner verteidigen lassen

Würfel:
	z.B.: 1d20, 1w20, 2D10, -1W4, ...
	"[+-]?([0-9]*)[dwDW]([0-9]*)(([+-])([0-9]*)){0,1}"`
)

var last Enemy

func cmdLoop(stinput func(string) string, itinput func(string) int64) {
CMDLOOP:
	for {
		cmd := strings.Split(stinput(""), " ")
		switch cmd[0] {
		case "w", "r":
			dice.CmdW(cmd[1:])
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
		case "ne", "new", "neu":
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
		case "ce", "create":
			fmt.Println("Erstelle Gegner: <name> <hd> <atd> <atb> <rkb>")
			name := stinput("Name")
			hitd := dice.CompileMultiple(stinput("Lebens Würfel"))
			attd := dice.CompileMultiple(stinput("Angriff Würfel"))
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
		case "help", "?":
			fmt.Println(help)
		} // Cmd Switch
	} // For
}
