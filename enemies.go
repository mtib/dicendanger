package main

import (
	"fmt"
	"github.com/mtib/dicendanger/dice"
)

// Encounter sammelt Gegner eines Encounters
type Encounter []Enemy

// RawEnemy wird benutzt um einen "lebendigen" Gegner zu generieren
type RawEnemy struct {
	Hitdice dice.WurfelClass
	Attdice dice.WurfelClass
	Attack  int64
	Armor   int64
}

// Enemy ist der "lebendige" Gegner
type Enemy struct {
	Name          string
	Health        int64
	Attdice       *dice.WurfelClass
	Attack, Armor *int64
}

var aktiveEnemies Encounter

// RawEnemyList enthaelt alle Rohdaten fuer exportierbare Gegner
type RawEnemyList map[string]RawEnemy

func rawEnemyGen(hd, ad string, at, ar int64) RawEnemy {
	return RawEnemy{dice.CompileMultiple(hd), dice.CompileMultiple(ad), at, ar}
}

func newEnemy(name string) Enemy {
	re, ok := rawEnemies[name]
	if !ok {
		return Enemy{}
	}
	return Enemy{name, re.Hitdice.RollAll(), &re.Attdice, &re.Attack, &re.Armor}
}

func (e Enemy) String() string {
	return fmt.Sprintf("%s:\n\tTP: %d\n\tATD: %v\n\tATB: %v\n\tRKL: %v", e.Name, e.Health, *e.Attdice, *e.Attack, *e.Armor)
}

func (re RawEnemy) String() string {
	return fmt.Sprintf("[%d:%d]TP, [%d:%d]ATD, [%d]ATB, [%d]RKL", re.Hitdice.Min(), re.Hitdice.Max(), re.Attdice.Min(), re.Attdice.Max(), re.Attack, re.Armor)
}
