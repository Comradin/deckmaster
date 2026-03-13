//go:build darwin
// +build darwin

package main

/*
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>

void deckmaster_key_event(uint16_t keycode, int down) {
    CGEventRef event = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)keycode, down != 0);
    if (event == NULL) return;
    CGEventPost(kCGSessionEventTap, event);
    CFRelease(event);
}
*/
import "C"

type cgKeyboard struct{}

func initKeyboard() (Keyboard, error) {
	return &cgKeyboard{}, nil
}

func (k *cgKeyboard) KeyPress(keycode int) error {
	C.deckmaster_key_event(C.uint16_t(keycode), 1)
	C.deckmaster_key_event(C.uint16_t(keycode), 0)
	return nil
}

func (k *cgKeyboard) KeyDown(keycode int) error {
	C.deckmaster_key_event(C.uint16_t(keycode), 1)
	return nil
}

func (k *cgKeyboard) KeyUp(keycode int) error {
	C.deckmaster_key_event(C.uint16_t(keycode), 0)
	return nil
}

func (k *cgKeyboard) Close() error { return nil }
