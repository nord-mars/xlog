package xlog

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// inject interface
type LogInterface interface {
	Write(debugLevel int, messagetype messageType, format string, message ...interface{})
	WriteI(debugLevel int, format string, message ...interface{})
	WriteW(debugLevel int, format string, message ...interface{})
	WriteE(debugLevel int, format string, message ...interface{})
	WriteD(debugLevel int, format string, message ...interface{})
}

// Main class
type Logger struct {
	*log.Logger
	level     int
	flags     int
	callLevel int
}

//
var def_flag int

func init() {
	def_flag = log.Ldate | log.Ltime | log.Lmicroseconds
}

// Constract new logger object - open/create the log file
//  filename - full path to log file. Example: [/var/log/server/my_sever.log]
//  level    - debug level
//  flags    - modify
//    file flag:
//           FILE_PID  : add PID  to filename.PID.log
//           FILE_DATE : add DATE to filename.YYYY-MM-DD.log
//           FILE_TIME : add TIME to filename.hh:mm:ss.log
//    line log flag:
//           log.flag: Ldate | Ltime | Lmicroseconds | Llongfile | Lshortfile | LUTC | Lmsgprefix | LstdFlags
//    line custom flag:
//           LINE_CALL : add __FILE__:__LINE__ . WARNING: debug only - two time slow
//           LINE_PID  :
//           LINE_HOST :
//           LINE_APP  :
func New(logname string, level int, flags int) *Logger {
	// make file name
	var ext = path.Ext(logname)
	filename := strings.TrimSuffix(path.Clean(logname), ext)

	// FILENAME: add PID
	if (flags & FILE_PID) == FILE_PID {
		dir := path.Dir(filename)
		if dir == "." {
			dir = ""
		}
		filename = fmt.Sprintf("%s/%s.%d", dir, path.Base(filename), os.Getpid())
	}

	// FILENAME: add YYYY-MM-DD
	if (flags & FILE_DATE) == FILE_DATE {
		now := time.Now()
		filename = fmt.Sprintf("%s.%04d-%02d-%02d", filename, now.Year(), now.Month(), now.Day())
	}

	// FILENAME: add HH:MM:SS
	if (flags & FILE_TIME) == FILE_TIME {
		now := time.Now()
		filename = fmt.Sprintf("%s.%02d:%02d:%02d", filename, now.Hour(), now.Minute(), now.Second())
	}

	filename = fmt.Sprintf("%s%s", filename, ext)

	// LOGLINE: add [PID]
	var prefix string = ""
	if (flags & LINE_PID) == LINE_PID {
		prefix = fmt.Sprintf("[%d] ", os.Getpid())
	}

	// LOGLINE: add [HOSTNAME]
	if (flags & LINE_HOST) == LINE_HOST {
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		prefix += fmt.Sprintf("[%s] ", name)
	}

	// LOGLINE: add [program name]
	if (flags & LINE_APP) == LINE_APP {
		filename := filepath.Base(os.Args[0])
		prefix += fmt.Sprintf("[%s] ", filename)
	}

	// flag by default
	if flags == 0 {
		flags = log.Ldate | log.Ltime
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		log.Fatal(err)
	}

	return &Logger{
		log.New(file, prefix, flags),
		level,
		flags,
		1,
	}
}

// Add Program(Application) name to log prefix.
// Field ProgramName usefull for log analizator
func (self *Logger) SetProgramName(name string) {
	self.SetPrefix(fmt.Sprintf("%s[%s] ", self.Prefix(), name))
}

// get debug level
func (self *Logger) DebugLevel(level int) int {
	return self.level
}

// Change debug level
// Runtime up/down level
// Up if problem.
// Down to speedup.
func (self *Logger) SetDebugLevel(level int) {
	self.level = level
}

// Append the message to log file.
//   debugLevel  - compare with object level to print or not
//   messagetype - INFO / WARNING / ERROR / FATAL
//      INFO / WARNING / ERROR - append message to log
//      FATAL - append message, call stack
//   format  - string message format
//   message - string varables array
func (self *Logger) Write(debugLevel int, messagetype messageType, format string, message ...interface{}) {

	if self.level < debugLevel {
		return
	}

	// Append to message: __FILE__:__LINE__
	var dbg string = ""
	if (self.flags & LINE_CALL) == LINE_CALL {
		_, filename, line, _ := runtime.Caller(self.callLevel)
		dbg = fmt.Sprintf("%s:%d: ", path.Base(filename), line)
	}

	msg := ""
	switch messagetype {
	case INFO:
		msg = fmt.Sprintf("%s  INFO - "+format, dbg, message)
	case WARN:
		msg = fmt.Sprintf("%s  WARN - "+format, dbg, message)
	case ERROR:
		msg = fmt.Sprintf("%s ERROR - "+format, dbg, message)
	case FATAL:
		_, filename, line, ok := runtime.Caller(self.callLevel)
		if ok {
			msg = fmt.Sprintf("%s:%d: FATAL - "+format, filename, line, message)
			stackSlice := make([]byte, 512)
			count := runtime.Stack(stackSlice, false)
			if count > 0 {
				msg += fmt.Sprintf("  CALL STACK:\n%s", stackSlice[0:count])
			}
		}
	}

	self.Printf(msg)
}

// Wrapper: write INFO
func (self *Logger) WriteI(debugLevel int, format string, message ...interface{}) {
	self.callLevel = 2
	self.Write(debugLevel, INFO, format, message...)
}

// Wrapper: write WARNING
func (self *Logger) WriteW(debugLevel int, format string, message ...interface{}) {
	self.callLevel = 2
	self.Write(debugLevel, WARN, format, message...)
}

// Wrapper: write ERROR
func (self *Logger) WriteE(debugLevel int, format string, message ...interface{}) {
	self.callLevel = 2
	self.Write(debugLevel, ERROR, format, message...)
}

// Wrapper: write DUMP
func (self *Logger) WriteD(debugLevel int, format string, message ...interface{}) {
	self.callLevel = 2
	self.Write(debugLevel, FATAL, format, message...)
}
