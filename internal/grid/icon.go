package grid

// Icon is an image in graph which is placed into certain position in grid
type Icon struct {
	ID         string
	Src        string
	Placement  Node
	Label      string
	LabelColor string
}
