package grid

// DrawingInstructions which defines nodeSize and general Grid margin in pixels
type DrawingInstructions struct {
	nodeSize int
	margin   int
}

// CalculateDimensions for image based on Grid and DrawingInstructions
func CalculateDimensions(grid Grid, instructions DrawingInstructions) (int, int) {
	width := grid.width*instructions.nodeSize + grid.width*instructions.margin + instructions.margin
	height := grid.height*instructions.nodeSize + grid.height*instructions.margin + instructions.margin
	return width, height
}
