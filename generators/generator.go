package npcgen

import "math/rand"
import "github.com/CanobbioE/npcgen/generators/lists"

// GenerateNPC returns the pointer to a newly generated random npc.
func GenerateNPC() *NPC {
	npc := NPC{}

	npc.genderIsFemale = rand.Int()%2 == 0
	npc.GenName()
	npc.age = lists.Ages[rand.Int()%len(lists.Ages)]
	npc.race = lists.Races[rand.Int()%len(lists.Races)]
	npc.nation = lists.Nations[rand.Int()%len(lists.Nations)]
	npc.voice = lists.Voices[rand.Int()%len(lists.Voices)]
	npc.character = lists.Characters[rand.Int()%len(lists.Characters)]
	npc.GenAspect()

	return &npc
}
