package npcgen

import "strconv"

// Stats represent the NPC's statistics (Strength, Dexterity, Constitution,
// Intelligence, Wisdom, Charisma).
type Stats struct {
	str int
	dex int
	con int
	iq  int
	wis int
	cha int
}

func (s *Stats) String() string {
	f := func(i int) string {
		return strconv.Itoa(i) + "\t"
	}

	ret := "\nSTR\tDEX\tCON\tINT\tWIS\tCHA\n" +
		f(s.str) + f(s.dex) + f(s.con) + f(s.iq) + f(s.wis) + f(s.cha)

	return ret
}
