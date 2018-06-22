package npcgen

type Aspect struct {
	hair   *Hair
	eyes   string
	height string
	skin   string
}

type Hair struct {
	color  string
	length string
}

func Aspect() *Aspect {

}
