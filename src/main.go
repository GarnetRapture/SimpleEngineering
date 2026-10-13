package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"simpleharness/src/internal/wizard"
)

//go:embed all:template
var templateFiles embed.FS

func main() {
	template, err := fs.Sub(templateFiles, "template")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(wizard.Run(os.Stdin, os.Stdout, os.Stderr, template))
}
