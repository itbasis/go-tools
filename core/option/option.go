package option

import (
	"log/slog"
)

const (
	_msgAlreadyBefore     = "the option has already been applied before"
	_msgAlreadyAndIgnored = "the option has already been applied before and will be ignored"
)

type LazyOptionFunc[T any] func() Option[T]

type Option[T any] interface {
	Key() Key
	Apply(*T) error
}

func ApplyOptions[O ~[]Option[T], T any](obj *T, opts O, defaults map[Key]LazyOptionFunc[T]) error {
	var keys = make(map[Key]struct{}, len(opts))

	for _, opt := range opts {
		if err := _applyOption[T](keys, _checkKeyErr, opt, obj); err != nil {
			return err
		}
	}

	for _, lazyOpt := range defaults {
		if err := _applyOption[T](keys, _checkKeySilent, lazyOpt(), obj); err != nil {
			return err
		}
	}

	return nil
}

func _applyOption[T any](keys map[Key]struct{}, checkKey _checkKey, opt Option[T], obj *T) error {
	var (
		key               = opt.Key()
		slogAttrOptionKey = _slogAttrOptionKey(key)
	)

	if err := _existKey(keys, key); err != nil {
		switch checkKey {
		case _checkKeyErr:
			slog.Error(_msgAlreadyBefore, slogAttrOptionKey)

			return err

		case _checkKeyWarn:
			slog.Warn(_msgAlreadyAndIgnored, slogAttrOptionKey)

			return nil

		case _checkKeySilent:
			slog.Debug(_msgAlreadyAndIgnored, slogAttrOptionKey)

			return nil
		}
	}

	slog.Debug("apply option", slogAttrOptionKey, slog.Int("check_key", int(checkKey)))

	keys[key] = struct{}{}

	return opt.Apply(obj) //nolint:wrapcheck // TODO
}
