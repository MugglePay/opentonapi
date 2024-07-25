package core

import "errors"

var ErrAccountNotFound = errors.New("account not found")
var ErrEntityNotFound = errors.New("entity not found")
var ErrTooManyEntities = errors.New("too many entities")
var ErrNotKeyBlock = errors.New("block must be a key block")
