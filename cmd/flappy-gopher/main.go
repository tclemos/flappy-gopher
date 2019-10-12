package main

import (
	"fmt"

	"github.com/tclemos/flappy-gopher/game"
)

func main() {
	fmt.Println("Flappy Gopher \\ʕ◔ϖ◔ʔ/")

	g := game.New()
	if err := g.Init(); err != nil {
		fmt.Println("Initializing game: ", err)
	}
}
