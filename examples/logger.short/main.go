package main

import (
	"github.com/nord-mars/xlog/pkg/xlog"
)

// Global interface
var iLog xlog.LogInterface

func main() {

	debugLevel := 3

	myLog := xlog.NewShort("/tmp/main_short.log", "xLogShortExample", debugLevel)

	iLog = myLog // Inject Loger to packet

	myLog.Write(0, xlog.INFO, "------------------START-----------------------")
	myLog.Write(1, xlog.INFO, "this is an information message!")
	myLog.Write(2, xlog.INFO, "this is an", "concated", "string!")
	myLog.Write(3, xlog.INFO, "this is     print (debugLevel < 3)!")
	myLog.Write(4, xlog.INFO, "this is not print (debugLevel < 4)!")

	wrapperFirst()
}

func wrapperFirst() {
	wrapperSecond()
}

func wrapperSecond() {
	wrapperFatal()
}

// CALL STACK example
func wrapperFatal() {
	iLog.Write(0, xlog.FATAL, "we crashed")
}
