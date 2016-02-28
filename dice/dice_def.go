package dice

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
