package core

import (
	"fmt"
	"os"
	"os/exec"
)

func Edit() {
	editCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("vi \"%s\"", AppConfig.Path))
	editCmd.Stdin = os.Stdin
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr

	if err := editCmd.Run(); err != nil {
		fmt.Printf("Failed editing file: %v\n", err)
		os.Exit(1)
	}
}
