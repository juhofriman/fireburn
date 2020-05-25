package yamlspec

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Container struct {
	Nodes     string
	Color     string
	Roundness int
	Children  []Container
	Placement string
}

type FireburnYAMLSpecification struct {
	Root   Container
	Output struct {
		NodeSize int `yaml:"node_size"`
		Margin   int
	}
	Color map[string]string
}

func ReadFireburnFile(path string) FireburnYAMLSpecification {
	spec := FireburnYAMLSpecification{}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	uerr := yaml.Unmarshal(content, &spec)
	if uerr != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("")

	return spec
}

func FireburnSpecificationToString(spec FireburnYAMLSpecification) []byte {
	content, err := yaml.Marshal(spec)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
