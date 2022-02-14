// Package logger ... documentation
package xlog

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// Main interface
type ConsoloeInterface interface {
	WriteConsole(messagetype messageType, message ...string)
}

// Main class
type console struct {
	pid int
}

//var def_flag int
func init() {
	//    def_flag = log.Ldate | log.Ltime | log.Lmicroseconds
}

// Constract new Object
func NewConsole() *console {
	return &console{
		os.Getpid(),
	}
}

// Append the message to log file.
//   messagetype - INFO / WARNING / ERROR / DEBUG / FATAL
//      INFO / WARNING / ERROR - append message to log
//      FATAL - append message, call STACK and EXIT the programm
//   message     - strig varables array
func (l *console) WriteConsole(messagetype messageType, message ...string) {
	switch messagetype {
	case INFO:
		fmt.Printf(InfoColor, fmt.Sprintf("\nInformation: \n%s\n", message))
	case WARNING:
		fmt.Printf(WarningColor, fmt.Sprintf("\nWarning: \n%s\n", message))
	case ERROR:
		fmt.Printf(ErrorColor, fmt.Sprintf("\nError: \n%s\n", message))
	case FATAL:
		_, filename, line, _ := runtime.Caller(1)
		fmt.Printf(FatalColor,
			fmt.Sprintf("\nFatal: \n%s\n%s\n",
				filename+":"+strconv.Itoa(line),
				message))
		stackSlice := make([]byte, 512)
		count := runtime.Stack(stackSlice, false)
		if count > 0 {
			fmt.Printf(FatalColor,
				fmt.Sprintf("%d stack:\n%s",
					l.pid,
					stackSlice[0:count]))
		}
	}
}
