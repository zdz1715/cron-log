package pkg

import (
	"fmt"
	"os/exec"
	"time"
)

const (
	PID       = "crond.pid"
	ShellSh   = "sh"
	ShellBash = "bash"
)

var shellMap = map[string]string{
	ShellSh:   "/bin/sh",
	ShellBash: "/bin/bash",
}

type LogFormatter func(params LogFormatterParams) string

type LogFormatterParams struct {
	TimeStamp time.Time
	User      string
	Output    string
	Error     error
	Command   string
	Shell     string
	Latency   time.Duration
}

func (p *LogFormatterParams) exec() {
	cmd := exec.Command("su", "-", p.User, shellMap[p.Shell], "-c", p.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		p.Error = err
	}
	p.Output = string(out)
}

func Collect(shell string, user string, command string) {
	param := LogFormatterParams{
		User:    user,
		Shell:   shell,
		Command: command,
	}
	param.TimeStamp = time.Now()
	param.exec()
	param.Latency = time.Now().Sub(param.TimeStamp)

	fmt.Printf("%+v\n", param)
}
