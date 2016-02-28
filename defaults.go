package main

// Name: HD ATD ATB ARB
var rawEnemies = RawEnemyList{
	"Bandit": rawEnemyGen("2d6+1 s1d4", "1d4+1", 3, 10),
	"Baer":   rawEnemyGen("2d6+6", "1d6", 4, 8),
	"Golem":  rawEnemyGen("3d10+2", "1d6+3", 6, 4),
	"Soldat": rawEnemyGen("3d6", "2d4", 4, 7),
	"Pixie":  rawEnemyGen("2d6", "1d4", 6, 14),
	"Magier": rawEnemyGen("3d6 1d10", "1d8", 6, 8),
	"Imp":    rawEnemyGen("2d6+3", "1d4 1d4 s1d4", 4, 8),
	"Goblin": rawEnemyGen("3d4", "1d4", 3, 8),
}
