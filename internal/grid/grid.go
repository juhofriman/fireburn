package grid

import "github.com/fogleman/gg"

// DrawingInstructions which defines nodeSize and general Grid margin in pixels
type DrawingInstructions struct {
	NodeSize, Margin int
	DesignMode       bool
}

// Node represents a single node in Grid
type Node struct {
	X, Y int
}

// Grid is the main container for diagrams. Grids can be nested and they also retain pointers
// to parent. You should not construct this struct directly, but instead use NewGrid and grid.Group()
type Grid struct {
	width  int
	height int

	color     string
	roundness int

	parent   *Grid
	children []*Grid

	placement Node

	icons []*Icon
}

// NewGrid creates new Grid and returns pointer.
// This creates new root grid and leaves parent reference to nil.
func NewGrid(width, height int, color string, roundness int) *Grid {
	return &Grid{
		width:     width,
		height:    height,
		color:     color,
		roundness: roundness,
		placement: Node{0, 0},
	}
}

// Group adds new group no grid and returns pointer to created grid
func (g *Grid) Group(placement Node, width, height int, color string, roundness int) *Grid {
	subGrid := NewGrid(width, height, color, roundness)
	subGrid.placement = placement
	subGrid.parent = g
	g.children = append(g.children, subGrid)
	return subGrid
}

// PlaceIcon places Icon into Grid
func (g *Grid) PlaceIcon(icon Icon) {
	g.icons = append(g.icons, &icon)
}

// DrawGrid grid draws grid
func DrawGrid(grid *Grid, drawingInstructions DrawingInstructions) *gg.Context {

	width, height := CalculateDimensions(grid, drawingInstructions)

	context := gg.NewContext(width, height)

	drawFromRoot(context, grid, drawingInstructions)

	return context
}
