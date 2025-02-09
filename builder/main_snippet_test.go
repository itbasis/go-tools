package main_test

// tag::documentation[]
import (
	"github.com/itbasis/go-tools/builder/cmd"
)

// If arguments were not passed, they are taken from the `os.Args`.
func SnippetRunWithoutArguments() {
	cmd.InitApp().Run()
}

func SnippetRunWithArguments() {
	cmd.InitApp().Run("generate", "--debug")
}

// end::documentation[]
