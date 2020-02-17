package log

import "github.com/yamakiller/magicLibs/log"

var (
	defaultLog log.LogAgent
)

//WithLog 设置全局日志模块
func WithLog(l log.LogAgent) {
	defaultLog = l
}

//Info ...
func Info(fmrt string, args ...interface{}) {
	defaultLog.Info("", fmrt, args...)
}

//Error ...
func Error(fmrt string, args ...interface{}) {
	defaultLog.Error("", fmrt, args...)
}

//Debug ...
func Debug(fmrt string, args ...interface{}) {
	defaultLog.Debug("", fmrt, args...)
}

//Warning ...
func Warning(fmrt string, args ...interface{}) {
	defaultLog.Warning("", fmrt, args...)
}

//Trace ...
func Trace(fmrt string, args ...interface{}) {
	defaultLog.Trace("", fmrt, args...)
}

//Fatal ...
func Fatal(fmrt string, args ...interface{}) {
	defaultLog.Fatal("", fmrt, args...)
}

//Panic ...
func Panic(fmrt string, args ...interface{}) {
	defaultLog.Panic("", fmrt, args...)
}
