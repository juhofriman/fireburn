package main

import (
	"fmt"

	"github.com/juhofriman/fireburn/internal/grid"
)

func main() {
	// This will be the main entry point with CLI stuff
	// ATM it's just fixed graph for demostrating and testing features
	fmt.Println("FIREBURN")

	rootGrid := grid.NewGrid(6, 10, "#000000")
	subGrid := rootGrid.Group(grid.Node{X: 1, Y: 1}, 4, 3, "#fee2b3")
	subGrid.Group(grid.Node{X: 1, Y: 1}, 2, 1, "#ffa299")
	rootGrid.Group(grid.Node{X: 1, Y: 5}, 4, 4, "#ad6989")

	context := grid.DrawGrid(rootGrid, grid.DrawingInstructions{
		NodeSize:   50,
		Margin:     10,
		DesignMode: true,
	})

	context.SavePNG("out.png")
}
