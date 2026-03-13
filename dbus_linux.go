//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus"
)

var dbusConn *dbus.Conn

func initDbus() error {
	var err error
	dbusConn, err = dbus.SessionBus()
	return err
}

func executeDBusMethod(object, path, method, args string) {
	if dbusConn == nil {
		fmt.Fprintln(os.Stderr, "dbus is not available!")
		return
	}
	call := dbusConn.Object(object, dbus.ObjectPath(path)).Call(method, 0, args)
	if call.Err != nil {
		fmt.Fprintf(os.Stderr, "dbus call failed: %s\n", call.Err)
	}
}
