package main

import (
	"log"
	"os"
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

	//cmd := exec.Command("cmd", "sub/main.exe")
	cmd := &exec.Cmd{
		Path: "cmd",
		// Using cmd
		Args: []string{"/c", "start", "cmd", "/k", ".\\sub\\main.exe"},

		// Using wt
		// Args: []string{"/c start wt sub/main.exe"},
	}
	//cmd.Stdin = strings.NewReader("sub/main.exe")

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// err := cmd.Run()
	err := cmd.Start()
	CheckError(err)

	time.Sleep(10 * time.Second)
}
