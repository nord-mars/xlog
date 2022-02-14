package xlog

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// Main interface
type LogInterface interface {
	Write(debugLevel int, messagetype messageType, message ...string)
}

// Main class
type Xlog struct {
	*log.Logger
	level int
	flags int
}

//
var def_flag int

func init() {
	def_flag = log.Ldate | log.Ltime | log.Lmicroseconds
}

// Constract new logger object - open/create the log file
//  filename - full path to log file example: [/var/log/server/my_sever.log]
//  level    - debug lavel
//  lflag    - log.flag: Ldate | Ltime | Lmicroseconds | Llongfile | Lshortfile | LUTC | Lmsgprefix | LstdFlags
//             logger.FILE_LINE: add __FILE__:__LINE__ . WARNING: debug only - two time slow
//             logger.FILE_PID      : add PID to filename.PID.log
//             logger.FILE_DATE     : add DATE to filename.YYYY-MM-DD.log
//             logger.FILE_TIME     : add TIME to filename.hh:mm:ss.log
//             logger.LINE_PID      :
func New(basename string, level int, flags int) *Xlog {
	// make file name
	var ext = path.Ext(basename)
	filename := strings.TrimSuffix(path.Clean(basename), ext)

	if (flags & FILE_PID) == FILE_PID {
		dir := path.Dir(filename)
		if dir == "." {
			dir = ""
		}
		filename = fmt.Sprintf("%s%s.%d", dir, path.Base(filename), os.Getpid())
	}

	if (flags & FILE_DATE) == FILE_DATE {
		now := time.Now()
		filename = fmt.Sprintf("%s.%04d-%02d-%02d", filename, now.Year(), now.Month(), now.Day())
	}

	if (flags & FILE_TIME) == FILE_TIME {
		now := time.Now()
		filename = fmt.Sprintf("%s.%02d:%02d:%02d", filename, now.Hour(), now.Minute(), now.Second())
	}

	filename = fmt.Sprintf("%s%s", filename, ext)
	//fmt.Println(filename)

	// add [PID]
	var prefix string = ""
	if (flags & LINE_PID) == LINE_PID {
		prefix = fmt.Sprintf("[%d] ", os.Getpid())
	}

	// flag by default
	if flags == 0 {
		flags = log.Ldate | log.Ltime
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		log.Fatal(err)
	}

	return &Xlog{
		log.New(file, prefix, flags),
		level,
		flags,
	}
}

// Append the message to log file.
//   debugLevel  - compare with object lavel to print or not
//   messagetype - INFO / WARNING / ERROR / FATAL
//      INFO / WARNING / ERROR - append message to log
//      FATAL - append message, call STACK and EXIT the programm
//   message - strig varables array
func (l *Xlog) Write(debugLevel int, messagetype messageType, message ...string) {

	if l.level < debugLevel {
		return
	}

	// Append to message: __FILE__:__LINE__
	var dbg string = ""
	if (l.flags & LINE_CALL) == LINE_CALL {
		_, filename, line, _ := runtime.Caller(1)
		dbg = fmt.Sprintf("%s:%d: ", path.Base(filename), line)
	}

	switch messagetype {
	case INFO:
		l.Printf("%s  INFO - %s", dbg, message)
	case WARNING:
		l.Printf("%s  WARN - %s", dbg, message)
	case ERROR:
		l.Printf("%s ERROR - %s", dbg, message)
	case FATAL:
		_, filename, line, _ := runtime.Caller(1)
		l.Printf("%s:%d: FATAL - %s", filename, line, message)
		stackSlice := make([]byte, 512)
		count := runtime.Stack(stackSlice, false)
		if count > 0 {
			l.Printf("  CALL STACK:\n%s", stackSlice[0:count])
		}
	}
}
