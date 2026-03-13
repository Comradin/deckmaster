//go:build darwin
// +build darwin

package main

import (
	"strconv"
	"strings"
)

// macOS Carbon virtual keycodes
var keycodes = map[string]int{
	"A":          0x00,
	"S":          0x01,
	"D":          0x02,
	"F":          0x03,
	"H":          0x04,
	"G":          0x05,
	"Z":          0x06,
	"X":          0x07,
	"C":          0x08,
	"V":          0x09,
	"B":          0x0B,
	"Q":          0x0C,
	"W":          0x0D,
	"E":          0x0E,
	"R":          0x0F,
	"Y":          0x10,
	"T":          0x11,
	"Num1":       0x12,
	"Num2":       0x13,
	"Num3":       0x14,
	"Num4":       0x15,
	"Num6":       0x16,
	"Num5":       0x17,
	"Equal":      0x18,
	"Num9":       0x19,
	"Num7":       0x1A,
	"Minus":      0x1B,
	"Num8":       0x1C,
	"Num0":       0x1D,
	"Rightbrace": 0x1E,
	"O":          0x1F,
	"U":          0x20,
	"Leftbrace":  0x21,
	"I":          0x22,
	"P":          0x23,
	"Enter":      0x24,
	"L":          0x25,
	"J":          0x26,
	"Apostrophe": 0x27,
	"K":          0x28,
	"Semicolon":  0x29,
	"Backslash":  0x2A,
	"Comma":      0x2B,
	"Slash":      0x2C,
	"N":          0x2D,
	"M":          0x2E,
	"Dot":        0x2F,
	"Tab":        0x30,
	"Space":      0x31,
	"Grave":      0x32,
	"Backspace":  0x33,
	"Esc":        0x35,
	"Rightmeta":  0x36,
	"Leftmeta":   0x37,
	"Leftshift":  0x38,
	"Capslock":   0x39,
	"Leftalt":    0x3A,
	"Leftctrl":   0x3B,
	"Rightshift": 0x3C,
	"Rightalt":   0x3D,
	"Rightctrl":  0x3E,
	"F17":        0x40,
	"Kpdot":      0x41,
	"Kpmult":     0x43,
	"Kpplus":     0x45,
	"Numlock":    0x47,
	"Volumeup":   0x48,
	"Volumedown": 0x49,
	"Mute":       0x4A,
	"Kpslash":    0x4B,
	"Kpenter":    0x4C,
	"Kpminus":    0x4E,
	"F18":        0x4F,
	"F19":        0x50,
	"Kpequal":    0x51,
	"Kp0":        0x52,
	"Kp1":        0x53,
	"Kp2":        0x54,
	"Kp3":        0x55,
	"Kp4":        0x56,
	"Kp5":        0x57,
	"Kp6":        0x58,
	"Kp7":        0x59,
	"F20":        0x5A,
	"Kp8":        0x5B,
	"Kp9":        0x5C,
	"F5":         0x60,
	"F6":         0x61,
	"F7":         0x62,
	"F3":         0x63,
	"F8":         0x64,
	"F9":         0x65,
	"F11":        0x67,
	"F13":        0x69,
	"F16":        0x6A,
	"F14":        0x6B,
	"F10":        0x6D,
	"F12":        0x6F,
	"F15":        0x71,
	"Help":       0x72,
	"Home":       0x73,
	"Pageup":     0x74,
	"Delete":     0x75,
	"F4":         0x76,
	"End":        0x77,
	"F2":         0x78,
	"Pagedown":   0x79,
	"F1":         0x7A,
	"Left":       0x7B,
	"Right":      0x7C,
	"Down":       0x7D,
	"Up":         0x7E,
}

// pasteKeys returns the keycode string for the paste action (Cmd+V on macOS).
func pasteKeys() string { return "55-9" }

func formatKeycodes(key string) string {
	for k, v := range keycodes {
		if strings.EqualFold(key, k) {
			return strconv.Itoa(v)
		}
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
