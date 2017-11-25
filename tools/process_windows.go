package tools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"
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
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	// fmt.Printf("in all caps: %q\n", out.String())
	return out.String(), nil
}

func PIDIsRunning(file string) (bool, error) {
	if _, err := os.Stat(file); err != nil {
		// info("pid file missing?" + file)
		return false, nil
	}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return false, fmt.Errorf("cannot read pid file?" + file)
	}
	if string(dat) == "" {
		// info("pid empty")
		return false, nil
	}

	pid, err := strconv.Atoi(string(dat))
	if err != nil {
		// err_f("unable to convert to pid '%s'", string(dat))
		return false, nil
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		// err_f("process error: %s", err.Error())
		return false, nil
	} else {
		// info_f("process exists %d, %#v", pid, proc)
	}

	err = proc.Signal(syscall.Signal(0))
	// info_f("Pid %d err? %v", pid, err)
	if err != nil {
		return false, nil
	}

	return true, nil
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

func info(msg string) {
	fmt.Println(prog_name + "(i): " + msg)
}

func err_f(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	error_log(msg)
}

func error_log(msg string) {
	fmt.Println(prog_name + "(e): " + msg)
}
