package option

import (
	"log/slog"
	"slices"
)

type RestoreOption[T any] interface {
	Option[T]

	Save(*T) error
	Restore(*T) error
}

func ApplyRestoreOptions[T any](obj *T, opts []RestoreOption[T], action func()) error {
	var keys = make(map[Key]struct{}, len(opts))

	for _, opt := range opts {
		var key = opt.Key()

		if err := _existKey(keys, key); err != nil {
			slog.Error(_msgAlreadyBefore, _slogAttrOptionKey(key))

			return err
		}

		keys[key] = struct{}{}

		if err := opt.Save(obj); err != nil {
			return err //nolint:wrapcheck // TODO
		}

		if err := opt.Apply(obj); err != nil {
			return err //nolint:wrapcheck // TODO
		}
	}

	action()

	for _, opt := range slices.Backward(opts) {
		if err := opt.Restore(obj); err != nil {
			return err //nolint:wrapcheck // TODO
		}
	}

	return nil
}
