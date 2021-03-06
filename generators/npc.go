package npcgen

import (
	"math/rand"

	"github.com/CanobbioE/npcgen/generators/lists"
)

// NPC represents an npc description (name, age, gender, race,
// nationality, voice, character and ascpect)
type NPC struct {
	name           string
	age            string
	genderIsFemale bool
	race           string
	nation         string
	voice          string
	character      string
	aspect         *Aspect
	stats          *Stats
}

// GenName randomly generates the NPC's name, based on its gender.
func (n *NPC) GenName() {
	var names []string
	if n.genderIsFemale {
		names = lists.FemaleNames
	} else {
		names = lists.MaleNames
	}
	n.name = names[rand.Int()%len(names)]
}

// GenAspect generates a random aspect for the NPC.
func (n *NPC) GenAspect() {
	aspect := Aspect{}
	aspect.GenHairs()
	aspect.eyes = lists.Colors[rand.Int()%len(lists.Colors)]
	aspect.GenHeight(n.race, n.age)
	aspect.skin = lists.Colors[rand.Int()%len(lists.Colors)]
	n.aspect = &aspect
}

// String returns a string representation of the NPC object.
func (n *NPC) String() string {
	gender := "M"
	if n.genderIsFemale {
		gender = "F"
	}

	ret := "Nome: " + n.name +
		"\n" + "Eta: " + n.age +
		"\n" + "Sesso: " + gender +
		"\n" + "Razza: " + n.race +
		"\n" + "Nazione: " + n.nation +
		"\n" + "Voce: " + n.voice +
		"\n" + "Carattere: " + n.character +
		"\n" + "Aspetto: " + n.aspect.String() +
		"\n" + n.stats.String()

	return ret
}

// GenStats randomly generate statistics for the NPC.
func (n *NPC) GenStats() {
	stats := Stats{}

	stats.str = rand.Int() % 20
	stats.dex = rand.Int() % 20
	stats.con = rand.Int() % 20
	stats.iq = rand.Int() % 20
	stats.wis = rand.Int() % 20
	stats.cha = rand.Int() % 20

	n.stats = &stats
}
