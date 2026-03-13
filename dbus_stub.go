//go:build !linux && !darwin
// +build !linux,!darwin

package main

import (
	"fmt"
	"os"
)

func initDbus() error {
	return nil
}

func executeDBusMethod(object, path, method, args string) {
	fmt.Fprintln(os.Stderr, "dbus actions are not supported on this platform")
}
