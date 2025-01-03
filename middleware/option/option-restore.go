package option

import (
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

		if err := _existKey(keys, key, _checkKeyErr); err != nil {
			return err
		}

		keys[key] = struct{}{}

		if err := opt.Save(obj); err != nil {
			return err
		}

		if err := opt.Apply(obj); err != nil {
			return err
		}
	}

	action()

	for _, opt := range slices.Backward(opts) {
		if err := opt.Restore(obj); err != nil {
			return err
		}
	}

	return nil
}
