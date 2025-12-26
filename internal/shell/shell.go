package shell

import (
	"errors"
	"os"
	"os/exec"
)

type Data struct {
	FailedCommand string
	Error         string
}

func (d Data) String() string {
	return "{" + d.FailedCommand + ", " + d.Error + "}"
}

type Provider interface {
	Get() (error, Data)
}

type DefaultShellProvider struct{}

func NewShellProvider() Provider {
	return DefaultShellProvider{}
}

func (DefaultShellProvider) Get() (error, Data) {
	lastCmd := os.Getenv("LAST_FAILED_CMD")
	if lastCmd == "" {
		return errors.New("your history is clean, no need to use whai"), Data{}
	}

	cmd := exec.Command("/bin/sh", "-c", lastCmd)

	output, err := cmd.CombinedOutput()

	if err == nil {
		return err, Data{}
	}

	return nil, Data{
		FailedCommand: lastCmd,
		Error:         string(output),
	}
}
