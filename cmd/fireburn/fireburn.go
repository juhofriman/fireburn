package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/juhofriman/fireburn/internal/grid"
	"github.com/juhofriman/fireburn/internal/yamlspec"
	"github.com/spf13/cobra"
)

func parseDimensions(dim string) (int, int) {
	fmt.Printf("parsing %s\n", dim)
	dimensions := strings.SplitN(dim, "x", 2)
	var width, height = 0, 0
	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		panic("Aaargh")
	}
	height, err2 := strconv.Atoi(dimensions[1])
	if err2 != nil {
		panic("Aaargh")
	}
	return width, height
}

func createChildren(parent *grid.Grid, container *yamlspec.Container) *grid.Grid {
	for _, child := range container.Children {
		width, height := parseDimensions(child.Nodes)
		group := parent.Group(grid.NodeOf(child.Placement), width, height, child.Color, child.Roundness)
		createChildren(group, &child)
	}
	return parent
}

func main() {

	var cmdScaffolf = &cobra.Command{
		Use:   "scaffold GRID_SIZE [FILE]",
		Short: "Scaffold new fireburn yaml template",
		Long: `Scaffolds new yaml template with given size. 
GRID_SIZE must be suplied in form 10x5, which equals 10 wide and 5 nodes tall grid`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			spec := yamlspec.FireburnYAMLSpecification{}

			spec.Root.Nodes = args[0]
			spec.Root.Color = "#fefefe"
			spec.Root.Roundness = 0

			spec.Output.NodeSize = 50
			spec.Output.Margin = 10

			data := yamlspec.FireburnSpecificationToString(spec)

			fmt.Println(string(data))
		},
	}

	var designMode bool

	var cmdRender = &cobra.Command{
		Use:   "render [FILE]",
		Short: "Renders given template. Outputted image file name will match given template.",
		Long:  `Render given fireburn template`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			spec := yamlspec.ReadFireburnFile(args[0])
			fmt.Printf("%v\n", spec)
			width, height := parseDimensions(spec.Root.Nodes)
			rootGrid := grid.NewGrid(width, height, spec.Root.Color, spec.Root.Roundness)

			createChildren(rootGrid, &spec.Root)

			context := grid.DrawGrid(rootGrid, grid.DrawingInstructions{
				NodeSize:      spec.Output.NodeSize,
				Margin:        spec.Output.Margin,
				DesignMode:    designMode,
				ColorMappings: spec.Color,
			})

			context.SavePNG("out.png")
		},
	}

	cmdRender.Flags().BoolVarP(&designMode, "design", "d", false, "Output in design mode")

	var rootCmd = &cobra.Command{Use: "fireburn"}
	rootCmd.AddCommand(cmdScaffolf, cmdRender)
	rootCmd.Execute()

	// Reminder how the API works
	// rootGrid := grid.NewGrid(6, 10, "#000000", 0)
	// subGrid := rootGrid.Group(grid.Node{X: 1, Y: 1}, 4, 3, "#fee2b3", 20)
	// nested := subGrid.Group(grid.Node{X: 1, Y: 1}, 2, 1, "#ffa299", 50)
	// large := rootGrid.Group(grid.Node{X: 1, Y: 5}, 4, 4, "#ad6989", 40)

	// nested.PlaceIcon(grid.Icon{
	// 	ID:         "icon-1",
	// 	Src:        "/Users/juhofr/icons/aws-icons/PNG Dark/Storage/Amazon-FSx@4x.png",
	// 	Placement:  grid.Node{X: 0, Y: 0},
	// 	Label:      "FX 1",
	// 	LabelColor: "#ffffff",
	// })
	// nested.PlaceIcon(grid.Icon{
	// 	ID:         "icon-1",
	// 	Src:        "/Users/juhofr/icons/aws-icons/PNG Dark/Storage/Amazon-FSx@4x.png",
	// 	Placement:  grid.Node{X: 1, Y: 0},
	// 	Label:      "FX 2",
	// 	LabelColor: "#ffffff",
	// })

	// large.PlaceIcon(grid.Icon{
	// 	ID:        "icon-1",
	// 	Src:       "/Users/juhofr/icons/aws-icons/PNG Dark/Storage/Amazon-FSx@4x.png",
	// 	Placement: grid.Node{X: 1, Y: 0},
	// })

	// context := grid.DrawGrid(rootGrid, grid.DrawingInstructions{
	// 	NodeSize:   100,
	// 	Margin:     30,
	// 	DesignMode: false,
	// })

	// context.SavePNG("out.png")
}
