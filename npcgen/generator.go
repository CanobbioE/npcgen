package npcgen

import "math/rand"
import "github.com/CanobbioE/npcgen/npcgen/lists"

func generateNPC() *NPC {
	npc := NPC{}

	npc.genderIsFemale = rand.Int()%2 == 0
	npc.genName()
	npc.age = lists.Ages[rand.Int()%len(lists.Ages)]
	npc.race = lists.Races[rand.Int()%len(lists.Races)]
	npc.nation = lists.Nations[rand.Int()%len(lists.Nations)]

	return &npc
}
