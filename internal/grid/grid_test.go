package grid

import "testing"

func TestNewGrid(t *testing.T) {
	color := "#ffffff"
	height := 10
	width := 20
	grid := NewGrid(width, height, color)
	if got := grid.color; got != color {
		t.Errorf("NewGrid() did not return given color: %q, want %q", got, color)
	}
	if got := grid.width; got != width {
		t.Errorf("NewGrid() did not return given height: %q, want %q", got, width)
	}
	if got := grid.height; got != height {
		t.Errorf("NewGrid() did not return given height: %q, want %q", got, height)
	}
	if got := grid.parent; got != nil {
		t.Errorf("NewGrid() did not return grid with parent pointer nil")
	}
	if got := grid.children; got != nil {
		t.Errorf("NewGrid() did not return grid with children nil")
	}
	expectedNode := Node{0, 0}
	if got := grid.placement; got != expectedNode {
		t.Error("NewGrid() did not return grid with placement to {0, 0}")
	}
}

func TestGroup(t *testing.T) {
	grid := NewGrid(10, 10, "#ffffff")
	node := Node{0, 0}
	color := "#ffffff"
	height := 10
	width := 20
	subGrid := grid.Group(node, width, height, color)
	if got := subGrid.color; got != color {
		t.Errorf("grid.Group() did not return given color: %q, want %q", got, color)
	}
	if got := subGrid.width; got != width {
		t.Errorf("grid.Group() did not return given height: %q, want %q", got, width)
	}
	if got := subGrid.height; got != height {
		t.Errorf("grid.Group() did not return given height: %q, want %q", got, height)
	}
	if got := subGrid.children; got != nil {
		t.Errorf("grid.Group() did not return grid with children nil")
	}
	if got := subGrid.placement; got != node {
		t.Errorf("grid.Group() did not return grid with placement to %v", node)
	}

	if subGrid.parent != grid {
		t.Error("Created subgrid did not place pointer to parent")
	}
}

func TestGridIconPlacement(t *testing.T) {
	grid := NewGrid(10, 10, "#ffffff")
	grid.PlaceIcon(Icon{
		id:        "icon",
		src:       "image.png",
		placement: Node{1, 1},
		label:     "Icon Label",
	})

	if len(grid.icons) != 1 {
		t.Errorf("After adding icon the size of icons is not 1, was=%d", len(grid.icons))
	}

	grid.PlaceIcon(Icon{
		id:        "icon2",
		src:       "image.png",
		placement: Node{1, 1},
		label:     "Icon Label",
	})

	if len(grid.icons) != 2 {
		t.Errorf("After adding icon the size of icons is not 2, was=%d", len(grid.icons))
	}
}
