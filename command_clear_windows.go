//go:build windows
// +build windows

package main

import (
	"os"
	"os/exec"
)

func callbackClear(cfg *config, args ...string) error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}
