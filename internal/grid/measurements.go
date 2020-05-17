package grid

// CalculateDimensions for image based on Grid and DrawingInstructions
func CalculateDimensions(grid *Grid, instructions DrawingInstructions) (int, int) {
	width := grid.width*instructions.NodeSize + grid.width*instructions.Margin + instructions.Margin
	height := grid.height*instructions.NodeSize + grid.height*instructions.Margin + instructions.Margin
	return width, height
}
