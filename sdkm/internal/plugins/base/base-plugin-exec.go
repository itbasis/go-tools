package base

import (
	"io"
	"os"

	itbasisMiddlewareExec "github.com/itbasis/tools/middleware/exec"
	itbasisMiddlewareOs "github.com/itbasis/tools/middleware/os"
	"github.com/pkg/errors"
)

func (receiver *basePlugin) Exec(
	cli string,
	overrideEnv map[string]string,
	stdIn io.Reader, stdOut, stdErr io.Writer,
	args []string,
) error {
	var envMap = itbasisMiddlewareOs.MergeEnvAsMap(os.Environ(), overrideEnv)

	envMap["SDKM_BACKUP_PATH"] = envMap["PATH"]
	envMap["PATH"] = itbasisMiddlewareOs.CleanPath(envMap["PATH"], itbasisMiddlewareOs.ExecutableDir())

	if err := os.Setenv("PATH", envMap["PATH"]); err != nil {
		return errors.Wrap(err, "error setting PATH environment variable")
	}

	cmd, err := itbasisMiddlewareExec.NewExecutable(
		cli,
		itbasisMiddlewareExec.WithArgs(args...),
		itbasisMiddlewareExec.WithCustomIn(stdIn),
		itbasisMiddlewareExec.WithCustomOut(stdOut, stdErr),
		itbasisMiddlewareExec.WithEnvAsMap(envMap),
	)
	if err != nil {
		return errors.Wrap(err, "error executing plugin")
	}

	if err := cmd.Execute(); err != nil {
		return errors.Wrap(err, itbasisMiddlewareExec.ErrFailedExecuteCommand.Error())
	}

	return nil
}
