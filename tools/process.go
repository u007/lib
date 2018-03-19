package tools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

var prog_name string = "process"

func SystemExecInDir(path string, command string, args ...string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	err = os.Chdir(path)
	if err != nil {
		return "", err
	}
	out, err := SystemExec(command, args...)
	if err != nil {
		os.Chdir(cwd)
		return out, err
	}

	err = os.Chdir(cwd)
	// fmt.Printf("in all caps: %q\n", out.String())
	return out, err
}

// @return pid
func SystemRunBg(command string, args ...string) (int, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		return 0, err
	}
	// fmt.Printf("in all caps: %q\n", out.String())
	return cmd.Process.Pid, nil
}

func SystemExec(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()
	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	// fmt.Printf("in all caps: %q\n", out.String())
	return out.String(), nil
}

func IsCurrentPIDByFile(file string) (bool, error) {
	pid := os.Getpid()

	if _, err := os.Stat(file); err != nil {
		// info("pid file missing?" + file)
		return false, nil
	}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return false, fmt.Errorf("cannot read pid file?" + file)
		// return false
	}
	if string(dat) == "" {
		// info("pid missing")
		return false, nil
	}

	file_pid, err := strconv.Atoi(string(dat))
	if err != nil {
		// err_f("err process: %s", err.Error())
		return false, nil
	}
	return pid == file_pid, nil
}

func info_f(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	info(msg)
}

func err_f(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	error_log(msg)
}
