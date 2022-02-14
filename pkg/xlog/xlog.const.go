package xlog

// Log level
type messageType int

const (
	INFO messageType = 0 + iota
	WARN
	ERROR
	FATAL
)

// Linux console color
const (
	ColorInfo    = "\033[1;34m%s\033[0m"
	ColorWarning = "\033[1;33m%s\033[0m"
	ColorError   = "\033[1;31m%s\033[0m"
	ColorFatal   = "\033[1;35m%s\033[40;5m"
)

// xlog flags
// logger flags: append fields
const (
	// FILE - filename
	FILE_DATE = 1 << 10 // FILE: append DATE to filename.DATE.log
	FILE_TIME = 1 << 11 // FILE: append TIME to filename.TIME.log
	FILE_PID  = 1 << 12 // FILE: append PID  to filename.PID.log
	// LINE - log fields
	LINE_PID  = 1 << 13 // LINE: add PID
	LINE_CALL = 1 << 14 // LINE: add __FILE__:__LINE__
	LINE_HOST = 1 << 15 // LINE: add HOSTNAME
	LINE_APP  = 1 << 16 // LINE: add Application name
)
