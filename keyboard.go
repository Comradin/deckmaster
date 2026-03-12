package main

// Keyboard is a platform-agnostic interface for emulating keyboard input.
type Keyboard interface {
	KeyPress(keycode int) error
	KeyDown(keycode int) error
	KeyUp(keycode int) error
	Close() error
}
