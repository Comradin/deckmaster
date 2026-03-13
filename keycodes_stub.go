//go:build !linux && !darwin
// +build !linux,!darwin

package main

import (
	"strconv"
	"strings"
)

// keycodes is empty on unsupported platforms.
var keycodes = map[string]int{}

// pasteKeys returns the keycode string for the paste action.
func pasteKeys() string { return "29-47" }

func formatKeycodes(key string) string {
	if v, ok := keycodes[key]; ok {
		return strconv.Itoa(v)
	}
	return key
}

func parseKeycode(key string) (int, bool) {
	key = formatKeycodes(strings.TrimSpace(key))
	v, err := strconv.Atoi(key)
	if err != nil {
		return 0, false
	}
	return v, true
}
