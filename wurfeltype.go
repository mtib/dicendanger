package main

// Wurfel umfasst alles was man braucht um Wuerfel zu simulieren
type Wurfel struct {
	Desc     string
	Num, Typ int64
	Mod      string
	Mtyp     string
	Mvar     int64
	Sub      bool
}

// WurfelClass kombiniert Wuerfel
type WurfelClass []Wurfel

// RawEnemy wird benutzt um einen "lebendigen" Gegner zu generieren
type RawEnemy struct {
	Hitdice WurfelClass
	Attdice WurfelClass
	Attack  int64
	Armor   int64
}

// Enemy ist der "lebendige" Gegner
type Enemy struct {
	Name          string
	Health        int64
	Attdice       *WurfelClass
	Attack, Armor *int64
}

// Encounter sammelt Gegner eines Encounters
type Encounter []Enemy
