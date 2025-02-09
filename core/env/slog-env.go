package env

import (
	"log/slog"

	itbasisCoreLog "github.com/itbasis/go-tools/core/log"
)

func SlogAttrEnv[E ListOrMap](env E) slog.Attr {
	if list, ok := any(env).(List); ok {
		return itbasisCoreLog.SlogAttrMap("env", SlicesToMap(list))
	}

	return itbasisCoreLog.SlogAttrMap("env", any(env).(Map))
}
