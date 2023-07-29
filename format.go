package go_utils

import (
	"fmt"
	"os/exec"
)

func FormatGoCode(filePath string) error {
	cmd := exec.Command("gofmt", "-w", filePath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to goutil Go code: %v", err)
	}
	return nil
}
