package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecCommand(module string, command string, args ...string) []byte {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("Hmmm, something went wrong (" + module + ")")
		os.Exit(1)
	}

	return stdout
}
