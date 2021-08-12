package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	ShellSh   = "sh"
	ShellBash = "bash"
)

var shellMap = map[string]string{
	ShellSh:   "/bin/sh",
	ShellBash: "/bin/bash",
}

type LogFormatterParams struct {
	TimeStamp time.Time
	User      string
	Output    string
	Error     error
	Command   string
	Shell     string
	Latency   time.Duration
	PID       string
}

func (p *LogFormatterParams) format() error {
	formatCommand := fmt.Sprintf("echo '%s' >> /proc/%s/fd/1 2>&1", p.Output, p.PID)
	cmd := exec.Command(shellMap[p.Shell], "-c", formatCommand)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (p *LogFormatterParams) getPid(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	p.PID = strings.TrimSpace(string(fd))
	return nil
}

func (p *LogFormatterParams) exec() {
	cmd := exec.Command("su", "-", p.User, shellMap[p.Shell], "-c", p.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		p.Error = err
	}
	p.Output = string(out)
}

func Collect(shell string, user string, command string, pidPath string) error {
	param := LogFormatterParams{
		User:    user,
		Shell:   shell,
		Command: command,
	}
	if err := param.getPid(pidPath); err != nil {
		return err
	}

	param.TimeStamp = time.Now()
	param.exec()
	param.Latency = time.Now().Sub(param.TimeStamp)

	fmt.Printf("%+v\n", param)
	if err := param.format(); err != nil {
		return err
	}
	return nil
}
