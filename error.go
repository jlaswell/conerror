/*
Copyright 2017 John Laswell

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package conerror provides contextual errors.
package conerror

import (
	"context"
	"errors"
)

// ConError is a contextual error. It can be used as a std error and to hold
// extra information via the Get() and Set() methods available.
type ConError interface {
	error

	// Get contextual information.
	Get(key interface{}) interface{}
	// Set contextual information.
	Set(key, val interface{})
}

// Error provides an implementation of the ConError interface.
type Error struct {
	ctx context.Context
	err error
}

// New create a new ConError from an error message.
func New(text string) ConError {
	return &Error{
		ctx: context.Background(),
		err: errors.New(text),
	}
}

// NewFromError creates a new ConError from an error.
func NewFromError(err error) ConError {
	return &Error{
		ctx: context.Background(),
		err: err,
	}
}

// Error returns the message for this error.
func (e *Error) Error() string {
	return e.err.Error()
}

// Get contextual information.
func (e *Error) Get(key interface{}) interface{} {
	return e.ctx.Value(key)
}

// Set contextual information.
func (e *Error) Set(key, val interface{}) {
	e.ctx = context.WithValue(e.ctx, key, val)
}
