//go:build darwin
// +build darwin

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

// Window describes an application window (stub for macOS).
type Window struct {
	ID    uint32
	Class string
	Name  string
	Icon  image.Image
}

// Xorg is a stub on macOS; window tracking is not supported.
type Xorg struct{}

// Connect always returns an error on macOS, disabling window tracking.
func Connect(display string) (*Xorg, error) {
	return nil, errors.New("X11 window tracking is not available on macOS")
}

// Close is a no-op on macOS.
func (x Xorg) Close() {}

// TrackWindows is a no-op on macOS.
func (x *Xorg) TrackWindows(ch chan interface{}, interval time.Duration) {}

// CloseWindow returns an error on macOS.
func (x Xorg) CloseWindow(w Window) error {
	return errors.New("window management not supported on macOS")
}

// RequestActivation returns an error on macOS.
func (x Xorg) RequestActivation(w Window) error {
	return errors.New("window management not supported on macOS")
}
