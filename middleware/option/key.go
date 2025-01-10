package option

import (
	"log/slog"

	"github.com/pkg/errors"
)

type Key string

const _msgAlreadyKey = "option has already been applied"

type _checkKey int

const (
	_checkKeySilent _checkKey = iota
	_checkKeyWarn   _checkKey = iota
	_checkKeyErr    _checkKey = iota
)

func _existKey(keys map[Key]struct{}, key Key) error {
	slogAttrOptionKey := _slogAttrOptionKey(key)

	if _, exist := keys[key]; !exist {
		slog.Debug("key not found", slogAttrOptionKey)

		return nil
	}

	slog.Debug(_msgAlreadyKey, slogAttrOptionKey)

	return errors.New(_msgAlreadyKey)
}
