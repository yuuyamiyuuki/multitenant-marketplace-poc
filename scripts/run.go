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

	fmt.Println("Updating swagger docs")
	swagCmd := exec.Command("swag", "init")
	swagCmd.Stdout = os.Stdout
	swagCmd.Stderr = os.Stderr

	if err := swagCmd.Run(); err != nil {
		fmt.Printf("Warning: Swagger generation failed: %v\n", err)
	}

	fmt.Println("Starting gin")
	ginCmd := exec.Command("go", "run", "main.go")
	ginCmd.Stdout = os.Stdout
	ginCmd.Stderr = os.Stderr

	if err := ginCmd.Run(); err != nil {
		fmt.Printf("Server exited with error: %v\n", err)
	}
}
