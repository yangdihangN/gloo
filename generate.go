package main

import (
	"github.com/solo-io/solo-kit/pkg/code-generator/cmd"
	"github.com/solo-io/solo-kit/pkg/code-generator/docgen/options"
	"github.com/solo-io/solo-kit/pkg/utils/log"
)

//go:generate go run generate.go

func main() {

	//err := version.CheckVersions()
	//if err != nil {
	//	log.Fatalf("generate failed!: %v", err)
	//}

	log.Printf("starting generate")
	docs := cmd.DocsOptions{
		Output: options.Restructured,
	}
	if err := cmd.Run("projects", true, &docs, nil, nil); err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
}
