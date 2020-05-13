package main

import (
	"fmt"

	"github.com/juhofriman/fireburn/internal/grid"
)

func main() {
	fmt.Println("FIREBURN")

	grid := grid.NewGrid(1, 1, "#ffffff")

	fmt.Printf("%v\n", grid)
}
