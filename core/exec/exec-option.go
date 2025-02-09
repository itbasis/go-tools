package exec

import (
	"os/exec"

	itbasisCoreOption "github.com/itbasis/go-tools/core/option"
)

type Option = itbasisCoreOption.Option[exec.Cmd]
type RestoreOption = itbasisCoreOption.RestoreOption[exec.Cmd]
