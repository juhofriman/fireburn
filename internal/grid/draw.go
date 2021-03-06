package grid

import (
	"fmt"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

type drawingContext struct {
	relativeX, relativeY int
}

func getNodePosition(node Node, grid *Grid, drawingInstructions DrawingInstructions, drawContext drawingContext) (int, int) {
	x := drawContext.relativeX + drawingInstructions.NodeSize*node.X + drawingInstructions.Margin*node.X + drawingInstructions.NodeSize/2
	y := drawContext.relativeY + drawingInstructions.NodeSize*node.Y + drawingInstructions.Margin*node.Y + drawingInstructions.NodeSize/2
	return x, y
}

func drawDesignHelpers(context *gg.Context, grid *Grid, drawingInstructions DrawingInstructions, drawContext drawingContext) {
	context.SetHexColor("#aaaaaa")
	context.SetLineWidth(1)
	if drawingInstructions.DesignMode {
		for i := 0; i < grid.width; i++ {
			context.DrawLine(
				float64(drawingInstructions.NodeSize*i+drawingInstructions.Margin*i+drawContext.relativeX),
				float64(drawContext.relativeY),
				float64(drawingInstructions.NodeSize*i+drawingInstructions.Margin*i+drawContext.relativeX),
				float64(drawingInstructions.NodeSize*grid.height+drawingInstructions.Margin*grid.height-drawingInstructions.Margin+drawContext.relativeY))
			context.DrawLine(
				float64(drawingInstructions.NodeSize*i+drawingInstructions.Margin*i+drawContext.relativeX+drawingInstructions.NodeSize),
				float64(drawContext.relativeY),
				float64(drawingInstructions.NodeSize*i+drawingInstructions.Margin*i+drawContext.relativeX+drawingInstructions.NodeSize),
				float64(drawingInstructions.NodeSize*grid.height+drawingInstructions.Margin*grid.height-drawingInstructions.Margin+drawContext.relativeY))
		}

		for i := 0; i < grid.height; i++ {
			context.DrawLine(
				float64(drawContext.relativeX),
				float64(drawingInstructions.NodeSize*i+drawingInstructions.Margin*i+drawContext.relativeY),
				float64(drawingInstructions.NodeSize*grid.width+drawingInstructions.Margin*grid.width-drawingInstructions.Margin+drawContext.relativeX),
				float64(drawingInstructions.NodeSize*i+drawingInstructions.Margin*i+drawContext.relativeY))
			context.DrawLine(
				float64(drawContext.relativeX),
				float64(drawingInstructions.NodeSize*i+drawingInstructions.Margin*i+drawContext.relativeY+drawingInstructions.NodeSize),
				float64(drawingInstructions.NodeSize*grid.width+drawingInstructions.Margin*grid.width-drawingInstructions.Margin+drawContext.relativeX),
				float64(drawingInstructions.NodeSize*i+drawingInstructions.Margin*i+drawContext.relativeY+drawingInstructions.NodeSize))
		}

		context.Stroke()

		for i := 0; i < grid.width; i++ {
			for n := 0; n < grid.height; n++ {
				x, y := getNodePosition(Node{X: i, Y: n}, grid, drawingInstructions, drawContext)
				context.DrawStringAnchored(fmt.Sprintf("(%d, %d)", i, n),
					float64(x),
					float64(y),
					0.5,
					0.5)
			}
		}
	}
}

func drawIcons(context *gg.Context, grid *Grid, drawingInstructions DrawingInstructions, drawContext drawingContext) {
	for _, icon := range grid.icons {
		if im, err := gg.LoadImage(icon.Src); err == nil {
			fmt.Println("Drawing icon")
			resizedImage := imaging.Resize(im, drawingInstructions.NodeSize, drawingInstructions.NodeSize, imaging.Lanczos)
			x, y := getNodePosition(icon.Placement, grid, drawingInstructions, drawContext)
			context.DrawImageAnchored(resizedImage, x, y, 0.5, 0.5)
			if icon.Label != "" {
				context.SetHexColor(icon.LabelColor)
				context.DrawStringAnchored(icon.Label, float64(x), float64(y+drawingInstructions.NodeSize/2), 0.5, 1.5)
			}
		} else {
			fmt.Println(err)
		}
	}
}

func draw(context *gg.Context, grid *Grid, drawingInstructions DrawingInstructions, drawContext drawingContext) {
	fmt.Printf("Drawing: %v\n", grid)
	context.SetHexColor(ResolveColor(drawingInstructions.ColorMappings, grid.color))
	if grid.parent == nil {
		context.Clear()
	} else {
		context.DrawRoundedRectangle(
			float64(drawContext.relativeX-drawingInstructions.Margin),
			float64(drawContext.relativeY-drawingInstructions.Margin),
			float64(grid.width*drawingInstructions.NodeSize+grid.width*drawingInstructions.Margin+drawingInstructions.Margin),
			float64(grid.height*drawingInstructions.NodeSize+grid.height*drawingInstructions.Margin+drawingInstructions.Margin),
			float64(grid.roundness),
		)
		context.Fill()
	}

	drawDesignHelpers(context, grid, drawingInstructions, drawContext)

	drawIcons(context, grid, drawingInstructions, drawContext)

	for _, child := range grid.children {
		draw(context, child, drawingInstructions, drawingContext{
			child.placement.X*drawingInstructions.NodeSize + child.placement.X*drawingInstructions.Margin + drawContext.relativeX,
			child.placement.Y*drawingInstructions.NodeSize + child.placement.Y*drawingInstructions.Margin + drawContext.relativeY,
		})
	}
}

func drawFromRoot(context *gg.Context, grid *Grid, drawingInstructions DrawingInstructions) {
	draw(context, grid, drawingInstructions, drawingContext{drawingInstructions.Margin, drawingInstructions.Margin})
}
