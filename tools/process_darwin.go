package tools

import (
	"fmt"
	"io/ioutil"
	"log/syslog"
	"os"
	"strconv"
	"syscall"
)

var logger, err = syslog.New(syslog.LOG_INFO, prog_name)

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

func info(msg string) {
	fmt.Println(prog_name + "(i): " + msg)
	logger.Info(msg)
}

func error_log(msg string) {
	fmt.Println(prog_name + "(e): " + msg)
	logger.Err(msg)
}
