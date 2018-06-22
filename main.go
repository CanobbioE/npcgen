package main

import (
	"fmt"

	"github.com/CanobbioE/npcgen/generators"
)

func main() {
	fmt.Println(npcgen.GenerateNPC().String())
}
