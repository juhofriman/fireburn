package grid

import "strings"

type Color struct {
	Value string
}

func ResolveColor(mappings map[string]string, color Color) string {
	if mappings[color.Value] != "" {
		return mappings[color.Value]
	}
	if !strings.HasPrefix(color.Value, "#") {
		// Just panic for now
		panic("Can't resolve color")
	}
	return color.Value
}
