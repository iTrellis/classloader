// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package class_loader

import (
	"errors"
	"reflect"
)

var ErrorType = reflect.TypeOf((*error)(nil)).Elem()

// reflective dynamic access errors
var (
	ErrFailedCreateInstance      = errors.New("failed create instance")
	ErrNotFoundConstructMethod   = errors.New("not found construct method")
	ErrBadActorInitFuncOutNumber = errors.New("the actor init function, result should error or nothing")
)
