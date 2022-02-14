package main

import (
	"log"

	"github.com/nord-mars/xlog/pkg/xlog"
)

// Global interface
var iLog xlog.LogInterface

func main() {

	debugLevel := 3

	// EXAMPLE 1
	logFlags := log.Ldate | // dt: 2020/12/15
		log.Ltime | // dt: 06:19:41
		log.Lmicroseconds | // dt: .800297
		log.Lmsgprefix | // [PREFIX-PLACE] - before / after
		xlog.LINE_CALL | // prefix: add [__FILE__:__LINE__]
		xlog.LINE_PID | // prefix: add [PID]
		xlog.LINE_HOST | // prefix: add [HOSTMAME]
		xlog.LINE_APP | // prefix: add [RUN FILE NAME]
		xlog.FILE_PID | // filename: add PID
		xlog.FILE_DATE | // filename: add DATE
		xlog.FILE_TIME // filename: add TIME
	// EXAMPLE 2 + 3
	//	logFlags := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix | xlog.PID | xlog.DATE | xlog.TIME
	//	logFlags := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix | log.Lshortfile

	myLog := xlog.New("/tmp/log/main.log", debugLevel, logFlags)
	myLog.SetProgramName("xLogExample") // set microservice name

	iLog = myLog // Inject Logger to packet

	myLog.Write(0, xlog.INFO, "------------------START-----------------------")
	myLog.Write(1, xlog.INFO, "this is an information message!")
	myLog.Write(2, xlog.INFO, "this is an", "concated", "string!")
	myLog.Write(3, xlog.INFO, "this is     print (debugLevel < 3)!")
	myLog.Write(4, xlog.INFO, "this is not print (debugLevel < 4)!")

	// xlod
	myLog.Write(0, xlog.WARN, "this is an warning!")
	myLog.Write(0, xlog.ERROR, "this is an error!")
	//	Log.Write(0, xlog.FATAL,   "we crashed")

	// ------------------------------------
	// Print: Ouput RAW message to FILE.log
	myLog.Print("PRINT: one string message\n")
	myLog.Printf("PRINTF: %d\n", 4)
	myLog.Println("PRINTLN: ", "New line message")

	// Fatal: Output to FILE.log and console
	//	Log.Fatal(  "Programm stop, return [exit status 1]\n")
	//	Log.Fatalln("Programm stop, return [exit status 1]")
	//	Log.Fatalf( "Programm stop, return [exit status 1]\n")

	// Panic: Output console only
	//	Log.Panic(  "Programm stop, return [exit status 2] and print call stack to console.\n")
	//	Log.Panicf( "Programm stop, return [exit status 2] and print call stack to console. %d\n", 100)
	//	Log.Panicln("Programm stop, return [exit status 2] and print call stack to console.")

	// ------------------------------------
	// Call stack example
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
