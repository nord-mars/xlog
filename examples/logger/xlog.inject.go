package main

import "github.com/nord-mars/xlog/pkg/xlog"

// Global interface
// var iLog xlog.LogInterface

func SetIlog(xlog xlog.LogInterface) {
	iLog = xlog
}
