package main

import (
	"fmt"
	"log"

	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func defaultLinkHandler(name, ref string) string {
	return fmt.Sprintf(":ref:`%s <%s>`", name, ref)
}

func main() {
	app := cmd.GlooCli(version.Version)
	disableAutoGenTag(app)
	emptyStr := func(s string) string { return "" }

	err := doc.GenReSTTreeCustom(app, "./docs/cli", emptyStr, defaultLinkHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func disableAutoGenTag(c *cobra.Command) {
	c.DisableAutoGenTag = true
	for _, c := range c.Commands() {
		disableAutoGenTag(c)
	}
}
