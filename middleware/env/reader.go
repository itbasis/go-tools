package env

import (
	"log/slog"

	itbasisMiddlewareLog "github.com/itbasis/tools/middleware/log"

	"github.com/caarlos0/env/v11"
)

func ReadEnvConfig(cfg any) error {
	return ReadEnvConfigWithOptions(cfg, nil)
}

func ReadEnvConfigWithOptions(cfg any, opts *env.Options) error {
	logger := slog.With(slog.Any("config", cfg), slog.Any("envOptions", opts))

	logger.Debug("reading config...")

	var err error
	if opts == nil {
		err = env.Parse(cfg)
	} else {
		err = env.ParseWithOptions(cfg, *opts)
	}

	if err != nil {
		logger.Error("failed to parse config", itbasisMiddlewareLog.SlogAttrError(err))

		return ErrFailedReadConfiguration
	}

	logger.Debug("read config.")

	return nil
}