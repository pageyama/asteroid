package main

import (
	"fmt"
	"os"

	"github.com/pageyama/asteroid/asteroid"
)

func main() {

	winWidth := int32(800)
	winHeight := int32(600)
	fps := uint32(60)

	if err := asteroid.Run(winWidth, winHeight, fps); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
