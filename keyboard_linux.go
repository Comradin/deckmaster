//go:build linux
// +build linux

package main

import "github.com/bendahl/uinput"

type uinputKeyboard struct {
	kb uinput.Keyboard
}

func initKeyboard() (Keyboard, error) {
	kb, err := uinput.CreateKeyboard("/dev/uinput", []byte("Deckmaster"))
	if err != nil {
		return nil, err
	}
	return &uinputKeyboard{kb: kb}, nil
}

func (k *uinputKeyboard) KeyPress(keycode int) error { return k.kb.KeyPress(keycode) }
func (k *uinputKeyboard) KeyDown(keycode int) error  { return k.kb.KeyDown(keycode) }
func (k *uinputKeyboard) KeyUp(keycode int) error    { return k.kb.KeyUp(keycode) }
func (k *uinputKeyboard) Close() error               { return k.kb.Close() }
