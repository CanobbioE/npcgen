package npcgen

import (
	"math/rand"
	"strconv"

	"github.com/CanobbioE/npcgen/generators/lists"
)

// Aspect represents the NPC aspects (hairstyle, eyes color,
// height and skin color
type Aspect struct {
	hairs  *Hairs
	eyes   string
	height string
	skin   string
}

// Hairs represents the NPC hairstyle.
type Hairs struct {
	color string
	style string
}

// GenHairs generates a random hairstyle for the NPC.
func (a *Aspect) GenHairs() {
	hairs := Hairs{}
	hairs.style = lists.Hairs[rand.Int()%len(lists.Hairs)]
	hairs.color = lists.Colors[rand.Int()%len(lists.Colors)]
	a.hairs = &hairs
}

// GenHeight generates a random height for the NPC, based on
// its race and age.
func (a *Aspect) GenHeight(race string, age string) {
	minHeight := 160
	maxHeight := 200
	switch race {
	case "Halfling":
		minHeight = 90
		maxHeight = 100
	case "Nano":
		minHeight = 120
		maxHeight = 150
	case "Dragonide":
		minHeight = 180
		maxHeight = 210
	case "Gnomo":
		minHeight = 90
		maxHeight = 120
	case "Mezzorco":
		minHeight = 150
		maxHeight = 210
	}

	if age == "bambino" {
		minHeight = minHeight * 4 / 5
		maxHeight = maxHeight * 4 / 5
	}

	a.height = strconv.Itoa(rand.Int()%(maxHeight-minHeight) + minHeight)
}

// String returns a string representation of the Aspect object.
func (a *Aspect) String() string {
	ret := "\n\t" + "Capelli: " + a.hairs.style +
		" " + a.hairs.color +
		"\n\t" + "Occhi: " + a.eyes +
		"\n\t" + "Altezza: " + a.height +
		"\n\t" + "Pelle: " + a.skin
	return ret
}
