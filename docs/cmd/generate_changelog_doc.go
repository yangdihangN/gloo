package main

import (
	"context"
	"fmt"
	"os"

	"github.com/solo-io/go-utils/errors"

	"github.com/solo-io/go-utils/changelogutils"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Printf("unable to run: %v\n", err)
		os.Exit(1)
	}
}

const (
	glooDocGen  = "gloo"
	glooEDocGen = "glooe"
	validArgs
)

var (
	InvalidInputError = func(arg string) error {
		return errors.Errorf("invalid input, must provide exactly one argument, either '%v' or '%v', (provided %v)",
			glooDocGen,
			glooEDocGen,
			arg)
	}
)

func run(ctx context.Context) error {
	args := os.Args
	if len(args) != 2 {
		return InvalidInputError(fmt.Sprintf("%v", len(args)-1))
	}
	target := args[1]

	var repoRootPath, repo, changelogDirPath string
	switch target {
	case glooDocGen:
		repoRootPath = ".."
		repo = "gloo"
		changelogDirPath = "../changelog"
	case glooEDocGen:
		repoRootPath = "../../solo-projects"
		repo = "solo-projects"
		changelogDirPath = "../../solo-projects/changelog"
	default:
		return InvalidInputError(target)
	}

	// consider writing to stdout to enhance makefile/io readability `go run cmd/main.go > changelogSummary.md`
	owner := "solo-io"
	w := os.Stdout
	err := changelogutils.GenerateChangelogFromLocalDirectory(ctx, repoRootPath, owner, repo, changelogDirPath, w)
	return err
}
