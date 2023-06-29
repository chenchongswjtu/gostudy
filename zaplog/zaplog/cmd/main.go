package main

import (
	"github.com/chenchong/zaplog/zaplog"
)

func main() {
	logger := zaplog.NewLogger()
	logger.Debugf("this is %s", "debug")
	logger.Infof("this is %s", "info")
	logger.Warnf("this is %s", "warn")
	logger.Errorf("this is %s", "error")

	logger1 := zaplog.MustLogger("test")
	logger1.Debugf("this is %s", "debug")
	logger1.Infof("this is %s", "info")
	logger1.Warnf("this is %s", "warn")
	logger1.Errorf("this is %s", "error")

	zaplog.L.Debugf("this is %s", "debug")
	zaplog.L.Infof("this is %s", "info")
	zaplog.L.Warnf("this is %s", "warn")
	zaplog.L.Errorf("this is %s", "error")

}
