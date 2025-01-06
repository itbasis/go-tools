package option

import (
	"fmt"
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

func _existKey(keys map[Key]struct{}, key Key, checkKey _checkKey) error {
	slogAttrOptionKey := _slogAttrOptionKey(key)

	if _, exist := keys[key]; !exist {
		slog.Debug("key not found", slogAttrOptionKey)

		return nil
	}

	switch checkKey {
	case _checkKeySilent:
		slog.Debug(_msgAlreadyKey, slogAttrOptionKey)

		return nil

	case _checkKeyWarn:
		slog.Warn(_msgAlreadyKey, slogAttrOptionKey)

		return nil

	case _checkKeyErr:
		slog.Error(_msgAlreadyKey, slogAttrOptionKey)

	default:
		slog.Error(fmt.Sprintf("unknown checkKey enum: %d", checkKey), slogAttrOptionKey)
	}

	return errors.New(_msgAlreadyKey)
}
