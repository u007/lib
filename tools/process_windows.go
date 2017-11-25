package tools

import (
	"fmt"
)

func info(msg string) {
	fmt.Println(prog_name + "(i): " + msg)
}

func error_log(msg string) {
	fmt.Println(prog_name + "(e): " + msg)
}
