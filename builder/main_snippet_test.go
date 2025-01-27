//nolint:govet // Snippets file
package main_test

// tag::documentation[]
import (
	"github.com/itbasis/go-tools/builder/cmd"
)

// end::documentation[]
// nolint:unused
// tag::documentation[]
// If arguments were not passed, they are taken from the `os.Args`.
func SnippetRunWithoutArguments() {
	cmd.InitApp().Run()
}

// end::documentation[]
// nolint:unused
// tag::documentation[]
func SnippetRunWithArguments() {
	cmd.InitApp().Run("generate", "--debug")
}

// end::documentation[]
