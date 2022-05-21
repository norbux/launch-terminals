package main

import (
	"log"
	"os/exec"
	"time"
)

func CheckError(e error) {
	if e != nil {
		log.Println(e)
	}
}

func main() {
	log.Println("testing commands...")

	cmd := &exec.Cmd{
		Path: "cmd",
		// Using cmd
		Args: []string{"/c", "start", "cmd", "/k", ".\\external-prog-1\\main.exe"},

		// Using wt
		// Args: []string{"/c start wt external-prog-1/main.exe"},
	}

	err := cmd.Start()
	CheckError(err)

	cmd2 := &exec.Cmd{
		Path: "cmd",
		// Using cmd
		Args: []string{"/c", "start", "cmd", "/k", ".\\external-prog-2\\main.exe"},

		// Using wt
		// Args: []string{"/c start wt external-prog-2/main.exe"},
	}

	err = cmd2.Start()
	CheckError(err)

	time.Sleep(10 * time.Second)
}
