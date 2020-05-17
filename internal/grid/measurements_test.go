package grid

import "testing"

func assertDimensions(t *testing.T, grid Grid, instructions DrawingInstructions, expectedWidth, expectedHeight int) {
	width, height := CalculateDimensions(&grid, instructions)
	if expectedWidth != width {
		t.Errorf("Expecting width %d but got %d", expectedWidth, width)
	}
	if expectedHeight != height {
		t.Errorf("Expecting height %d but got %d", expectedHeight, height)
	}
}

func TestCalculatingDimensions(t *testing.T) {
	assertDimensions(t,
		Grid{
			width:  1,
			height: 1,
		},
		DrawingInstructions{
			NodeSize: 10,
			Margin:   1,
		},
		12,
		12,
	)
	assertDimensions(t,
		Grid{
			width:  10,
			height: 10,
		},
		DrawingInstructions{
			NodeSize: 10,
			Margin:   10,
		},
		210,
		210,
	)
	assertDimensions(t,
		Grid{
			width:  2,
			height: 1,
		},
		DrawingInstructions{
			NodeSize: 10,
			Margin:   5,
		},
		35,
		20,
	)
}
