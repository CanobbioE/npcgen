package npcgen

import (
	"math/rand"

	"github.com/CanobbioE/npcgen/npcgen/lists"
)

type NPC struct {
	name           string
	age            string
	genderIsFemale bool
	race           string
	nation         string
	voice          string
	character      string
	aspect         *Aspect
}

func (n *NPC) genName() {
	var names []string
	if n.genderIsFemale {
		names = lists.FemaleNames
	} else {
		names = lists.MaleNames
	}
	n.name = names[rand.Int()%len(names)]
}
