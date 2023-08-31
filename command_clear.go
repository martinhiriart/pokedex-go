package main

import (
	"os"
	"os/exec"
)

func callbackClear(cfg *config, args ...string) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}
