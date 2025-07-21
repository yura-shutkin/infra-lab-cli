package config

import (
	"fmt"
	"os/exec"
)

func BinaryNotFoundError(binary string) error {
	return fmt.Errorf("%s not found", binary)
}

func IsBinaryInPath(binary string) bool {
	_, err := exec.LookPath(binary)
	if err != nil {
		return false
	}
	return true
}
