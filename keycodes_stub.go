//go:build !linux && !darwin
// +build !linux,!darwin

package main

import "strconv"

// keycodes is empty on unsupported platforms.
var keycodes = map[string]int{}

func formatKeycodes(key string) string {
	if v, ok := keycodes[key]; ok {
		return strconv.Itoa(v)
	}
	return key
}

func parseKeycode(key string) (int, bool) {
	key = formatKeycodes(key)
	v, err := strconv.Atoi(key)
	if err != nil {
		return 0, false
	}
	return v, true
}
