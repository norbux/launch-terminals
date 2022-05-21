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

	cmd := exec.Command("dapr")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	checkError(err)

	time.Sleep(10 * time.Second)
}
