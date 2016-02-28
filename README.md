# Dice 'N Danger Toolkit
[![Build Status](https://travis-ci.org/mtib/dicendanger.svg?branch=master)](https://travis-ci.org/mtib/dicendanger)
[![Golang](https://img.shields.io/badge/golang-1.4.1-brightgreen.svg)](https://golang.org/)
[![GoDoc](https://godoc.org/github.com/mtib/dicendanger?status.svg)](https://godoc.org/github.com/mtib/dicendanger)
[![Gitter](https://badges.gitter.im/mtib/dicendanger.svg)](https://gitter.im/mtib/dicendanger?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Dieses Toolkit soll dem technisch versierten Dungeon Master ein wenig Arbeit
abnehmen.

Nachdem die Binärdatei gestartet wurde befindet sich das Programm im interaktiven
Modus. Um die neuste Version zu installieren:

    go get -u github.com/mtib/dicendanger

## Funktionen
### Würfeln
Befehl: ```w``` oder ```r```.

```bash
w dice [dice, ...]
r dice [dice, ...]
<w|r> [-]<N[d|w|D|W]M> [...]
dice = [-]N[dwDW]M = "([0-9]*)[dwDW]([0-9]*)(([+-])([0-9]*)){0,1}"
NdM === NwM === NDM === NWM für N,M := Natürliche Zahlen
N := Anzahl der Würfe
M := Anzahl der Seiten
```
Beispiele:

```go
w 2d10          | werfe 2x 10-Seitige Würfel, geb die Summe aus
w 1w20 2d10     | werfe 1x 20-Seitigen Würfel + 2x 10-Seitige Würfel
r 1W10          | werfe 1x 10-Seitigen Würfel
r 1d20 -2d8     | werfe 1x D20 subtragiere 2x D8
```
Ausgabe:
```go
> w 1d20
Würfel: [+1D20] [1:20]
+1D20 --> 9
Ergebnis: 9
```

### Gegner Liste
Befehl: ```le``` oder ```list```
```go
> le
Gegner:
	 Imp
	 Goblin
	 Bandit
	 Baer
	 Golem
	 Soldat
	 Pixie
	 Magier
```
mit mehr Informationen: ```raws```
```go
> raws
Pixie: [2:12]TP, [1:4]ATD, [6]ATB, [14]RKL
Magier: [4:28]TP, [1:8]ATD, [6]ATB, [8]RKL
Imp: [5:15]TP, [1:7]ATD, [4]ATB, [8]RKL
Goblin: [3:12]TP, [1:4]ATD, [3]ATB, [8]RKL
Bandit: [2:12]TP, [2:5]ATD, [3]ATB, [10]RKL
Baer: [8:18]TP, [1:6]ATD, [4]ATB, [8]RKL
Golem: [5:32]TP, [4:9]ATD, [6]ATB, [4]RKL
Soldat: [3:18]TP, [2:8]ATD, [4]ATB, [7]RKL
```
### Gegnerklasse inspizieren
Befehl: ```raw <name>```
```go
> raw Imp   
[5:15]TP, [1:7]ATD, [4]ATB, [8]RKL
```

### Gegner "zum Leben erwecken"
Befehl: ```ne```
```go
> ne Imp
+2D6+3 --> 11           // TP wird ausgewürfelt
Imp:
	TP: 11
	ATD: [+1D4 +1D4 -1D4] [1:7]
	ATB: 4
	RKL: 8
```
Dieser Gegner wird dann als "last" indexiert.
Ein Gegner kann sich verteidigen ```last def [angriffswurf]``` oder selbst angreifen ```last att [rüstungsklasse]```.
```go
> last att
Rüstungsbonus> 13
+1D4 --> 1
+1D4 --> 2
-1D4 --> -1
Gegner trifft
Schaden: 2
```

### Eigene Gegnerklasse erstellen
Befehl: ```ce```
```go
> ce
Erstelle Gegner: <name> <hd> <atd> <atb> <rkb>
Name> Octocat
Lebens Würfel> 8w4
Angriff Würfel> 1w8
Angriff Bonus> 3
Rüstungsbonus> 8
// Nachdem die Klasse erstellt wurde wird last initialisiert
+8D4 --> 13 // Lebenspunkte von Octocat in last
```
Erstellte Gegnerklassen sind danach über ```raws```, ```raw <name>``` oder ```list``` zu finden.
Beim Halten werden nicht statische Gegnerklassen und nicht exportierte Gegnerklassen gelöscht.

### Gegnerklassen exportieren (JSON)
Befehl: ```export [datei.json]``` (Standart: dnd.json)

### Gegnerklassen importieren (JSON)
Befehl: ```import [datei.json]``` (Standart: dnd.json)

## Lizenz
    Copyright (c) 2016 Markus "mtib" Becker

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included
    in all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
    INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
    IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
    DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
    ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
    OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
