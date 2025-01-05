//nolint:govet // Example file
package main_test

// tag::documentation[]
import (
	"github.com/itbasis/tools/builder/cmd"
)

// If arguments were not passed, they are taken from the `os.Args`.
func ExampleRunWithoutArguments() {
	cmd.InitApp().Run()
}

func ExampleRunWithArguments() {
	cmd.InitApp().Run("generate", "--debug")
}

// end::documentation[]
