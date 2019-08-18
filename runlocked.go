package main

import (
	"context"
	"os"
	"os/exec"
	"time"

	"github.com/gofrs/flock"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fileLock := flock.New(".gloo-helm.lock")

	locked, err := fileLock.TryLockContext(ctx, time.Second/10)

	if err != nil {
		panic(err)
	}
	if !locked {
		panic("can't acquire helm lock")
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()

	if locked {
		// do work
		fileLock.Unlock()
	}
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		panic(err)
	}
}
