package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

var (
	object = flag.String("object", "org.kde.kwalletd6", "")
	path   = flag.String("path", "/modules/kwalletd6", "")
	method = flag.String("method", "org.kde.KWallet.isEnabled", "")
)

func main() {
	flag.Parse()

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	var s bool
	obj := conn.Object(*object, dbus.ObjectPath(*path))
	err = obj.Call(*method, 0).Store(&s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to call method %s %s %s: %s\n", *object, *path, *method, err)
		os.Exit(1)
	}

	fmt.Println("Result from calling:")
	fmt.Println(s)
}
