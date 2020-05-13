package grid

// Icon is an image in graph which is placed into certain position in grid
type Icon struct {
	id        string
	src       string
	placement Node
	label     string
}
