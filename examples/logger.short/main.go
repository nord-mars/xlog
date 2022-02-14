package main

import (
	"github.com/nord-mars/xlog/pkg/xlog"
)

// Global interface
var iLog xlog.LogInterface

func main() {

	debugLevel := 3

	Log := xlog.NewShort("/tmp/main_short.log", "xLogShortExample", debugLevel)

	iLog = Log // Inject Loger to packet

	Log.Write(0, xlog.INFO, "------------------START-----------------------")
	Log.Write(1, xlog.INFO, "this is an information message!")
	Log.Write(2, xlog.INFO, "this is an", "concated", "string!")
	Log.Write(3, xlog.INFO, "this is     print (debugLevel < 3)!")
	Log.Write(4, xlog.INFO, "this is not print (debugLevel < 4)!")

	wraperFirst()
}

func wraperFirst() {
	wraperSecond()
}

func wraperSecond() {
	wraperFatal()
}

// CALL STACK example
func wraperFatal() {
	iLog.Write(0, xlog.FATAL, "we crashed")
}
