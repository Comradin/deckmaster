//go:build !linux && !darwin
// +build !linux,!darwin

package main

import "errors"

type stubKeyboard struct{}

func initKeyboard() (Keyboard, error) {
	return nil, errors.New("keyboard emulation not supported on this platform")
}

func (k *stubKeyboard) KeyPress(keycode int) error { return errors.New("unsupported") }
func (k *stubKeyboard) KeyDown(keycode int) error  { return errors.New("unsupported") }
func (k *stubKeyboard) KeyUp(keycode int) error    { return errors.New("unsupported") }
func (k *stubKeyboard) Close() error               { return nil }
