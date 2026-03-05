//go:build ignore
package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
)

func main() {
    fmt.Println("Starting postgres")
    exec.Command("docker-compose", "up", "-d").Run()

    fmt.Println("Waiting for postgres")
    time.Sleep(3 * time.Second)

    fmt.Println("Starting gin")
    cmd := exec.Command("go", "run", "main.go")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}