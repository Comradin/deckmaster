//go:build !linux && !darwin
// +build !linux,!darwin

package main

import (
	"errors"
	"image"
	"time"
)

// ActiveWindowChangedEvent gets emitted when the active window changes.
type ActiveWindowChangedEvent struct {
	Window Window
}

// WindowClosedEvent gets emitted when a window gets closed.
type WindowClosedEvent struct {
	Window Window
}

// Window describes an application window (stub for unsupported platforms).
type Window struct {
	ID    uint32
	Class string
	Name  string
	Icon  image.Image
}

// Xorg is a stub on unsupported platforms; window tracking is not supported.
type Xorg struct{}

// Connect always returns an error on unsupported platforms.
func Connect(display string) (*Xorg, error) {
	return nil, errors.New("X11 window tracking is not available on this platform")
}

// Close is a no-op on unsupported platforms.
func (x *Xorg) Close() {}

// TrackWindows is a no-op on unsupported platforms.
func (x *Xorg) TrackWindows(ch chan interface{}, interval time.Duration) {}

// CloseWindow returns an error on unsupported platforms.
func (x *Xorg) CloseWindow(w Window) error {
	return errors.New("window management not supported on this platform")
}

// RequestActivation returns an error on unsupported platforms.
func (x *Xorg) RequestActivation(w Window) error {
	return errors.New("window management not supported on this platform")
}
