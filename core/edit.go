package core

import (
	"fmt"
	"os"
	"os/exec"
)

func Edit() error {
	editCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("vi \"%s\"", AppConfig.Path))
	editCmd.Stdin = os.Stdin
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr

	if err := editCmd.Run(); err != nil {
		return fmt.Errorf("failed editing file: %w", err)
	}

	return nil
}
