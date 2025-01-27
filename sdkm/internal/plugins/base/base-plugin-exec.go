package base

import (
	"io"
	"os"

	itbasisCoreExec "github.com/itbasis/go-tools/core/exec"
	itbasisCoreOs "github.com/itbasis/go-tools/core/os"
	"github.com/pkg/errors"
)

func (receiver *basePlugin) Exec(
	cli string,
	overrideEnv map[string]string,
	stdIn io.Reader, stdOut, stdErr io.Writer,
	args []string,
) error {
	var envMap = itbasisCoreOs.MergeEnvAsMap(os.Environ(), overrideEnv)

	envMap["SDKM_BACKUP_PATH"] = envMap["PATH"]
	envMap["PATH"] = itbasisCoreOs.CleanPath(envMap["PATH"], itbasisCoreOs.ExecutableDir())

	if err := os.Setenv("PATH", envMap["PATH"]); err != nil {
		return errors.Wrap(err, "error setting PATH environment variable")
	}

	cmd, err := itbasisCoreExec.NewExecutable(
		cli,
		itbasisCoreExec.WithArgs(args...),
		itbasisCoreExec.WithCustomIn(stdIn),
		itbasisCoreExec.WithCustomOut(stdOut, stdErr),
		itbasisCoreExec.WithEnvAsMap(envMap),
	)
	if err != nil {
		return errors.Wrap(err, "error executing plugin")
	}

	if err := cmd.Execute(); err != nil {
		return errors.Wrap(err, itbasisCoreExec.ErrFailedExecuteCommand.Error())
	}

	return nil
}
