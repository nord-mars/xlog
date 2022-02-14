package main

import (
	"fmt"
	"time"

	"github.com/nord-mars/xlog/pkg/xlog"
)

func main() {

	debugLevel := 3

	Log := xlog.NewShort("/tmp/main_speed.log", "xlogSpeedExample", debugLevel)

	Log.Write(0, xlog.INFO, "------------------START-----------------------")

	// check speed
	start := time.Now()

	for i := 0; i < 10000; i++ {
		Log.Write(3, xlog.INFO, "  for loop")
	}

	//
	now := time.Now()
	diff := now.Sub(start)
	fmt.Println("diff: ", diff)

}
