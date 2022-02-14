package xlog

import "log"

// Constract new logger object - short form
func NewShort(logname string, appName string, level int) *Logger {

	logFlags := log.Ldate | // dt: 2020/12/15
		log.Ltime | // dt: 06:19:41
		log.Lmicroseconds | // dt: .800297
		log.Lmsgprefix | // [PREFIX-PLACE] - before / after
		//                LINE_CALL       | // prefix: add [__FILE__:__LINE__]
		LINE_PID | // prefix: add [PID]
		LINE_HOST | // prefix: add [HOSTNAME]
		//                LINE_APP        | // prefix: add [RUN FILE NAME]
		FILE_PID | // filename: add PID
		FILE_DATE | // filename: add DATE
		FILE_TIME // filename: add TIME

	Log := New(logname, level, logFlags)
	Log.SetProgramName(appName)

	return Log
}
