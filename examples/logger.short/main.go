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

	myLog.WriteI(0, "------------------START-----------------------")
	myLog.WriteI(1, "this is an information message")
	myLog.WriteI(2, "this is an", "concated", "string!")
	myLog.WriteI(3, "this is     print (debugLevel < 3)!")
	myLog.WriteI(4, "this is not print (debugLevel < 4)!")

	myLog.WriteE(0, "this is the ERROR")
	myLog.WriteW(0, "this is the WaRNING")
	myLog.WriteD(0, "this is the DUMP")

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
