package xlog

// Log level
type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

// Linux console color
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	FatalColor   = "\033[1;35m%s\033[40;5m"
)

// logger flags
const (
	LINE_PID  = 512  // append PID  to log - 1 << 9
	LINE_CALL = 1024 // add __FILE__:__LINE__
	FILE_DATE = 2048 // append DATE to filename.PID.log
	FILE_TIME = 4096 // append TILE to filename.PID.log
	FILE_PID  = 8192 // add PID to filename.PID.log - 1 << 9
)
