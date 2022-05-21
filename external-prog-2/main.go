package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func checkError(e error) {
	if e != nil {
		log.Println(e)
	}
}

func main() {
	log.Println("sub program...")

	cmd := exec.Command("ping", "-n", "8", "127.0.0.1")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	checkError(err)

	time.Sleep(10 * time.Second)
}
