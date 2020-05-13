package grid

// Node represents a single node in Grid
type Node struct {
	x, y int
}

// Grid is the main container for diagrams. Grids can be nested and they also retain pointers
// to parent. You should not construct this struct directly, but instead use NewGrid and grid.Group()
type Grid struct {
	width  int
	height int
	color  string

	parent   *Grid
	children []*Grid

	placement Node

	icons []*Icon
}

// NewGrid creates new Grid and returns pointer.
// This creates new root grid and leaves parent reference to nil.
func NewGrid(width, height int, color string) *Grid {
	return &Grid{
		width:     width,
		height:    height,
		color:     color,
		placement: Node{0, 0},
	}
}

// Group adds new group no grid and returns pointer to created grid
func (g *Grid) Group(placement Node, width, height int, color string) *Grid {
	subGrid := NewGrid(width, height, color)
	subGrid.placement = placement
	subGrid.parent = g
	return subGrid
}

// PlaceIcon places Icon into Grid
func (g *Grid) PlaceIcon(icon Icon) {
	g.icons = append(g.icons, &icon)
}