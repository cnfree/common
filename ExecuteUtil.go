package common

import (
	"sync"
	"os/exec"
	"strings"
	"os"
)

type executeUtil struct {
	mutex sync.Mutex
}

var Execute = executeUtil{}

func (this executeUtil) RunIn(dir, name string, args ...string) error {
	_, err := this.RunInWithOutput(dir, name, args...)
	return err
}

func (this executeUtil) RunInWithOutput(dir, name string, args ...string) (string, error) {
	c := exec.Command(name, args...)
	c.Dir = dir
	o, err := c.CombinedOutput()
	return string(o), err
}

func (this executeUtil) Execute(cmd string, dir string, bindSTD bool, envs []string) (cmder *exec.Cmd, err error) {
	parts := strings.Fields(cmd)
	command := parts[0]
	args := parts[1:len(parts)]

	commander := exec.Command(command, args...)

	if dir != "" {
		commander.Dir = dir
	}

	if bindSTD {
		commander.Stderr = os.Stderr
		commander.Stdout = os.Stdout
		commander.Stdin = os.Stdin
	}

	commander.Env = append(os.Environ(), envs...)

	if err = commander.Start(); err != nil {
		return
	}

	cmder = commander

	return
}
