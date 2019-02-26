package main

import (
	"fmt"
	"github.com/naichadouban/mylog/mylog"
	"github.com/naichadouban/mylog/mylog/rotator"
	"mylog/test"
	"os"
	"path/filepath"
)

const (
	maxRejectReasonLen = 250
)
var (
	// backendLog is the logging backend used to create all subsystem loggers.
	// The backend must not be used before the log rotator has been initialized,
	// or data races and/or nil pointer dereferences will occur.
	backendLog = mylog.NewBackend(logWriter{})

	// logRotator is one of the logging outputs.  It should be closed on
	// application shutdown.
	logRotator *rotator.Rotator

	mainLog = backendLog.Logger("MAIN")
	testLog = backendLog.Logger("TEST")

)
// logWriter 实现了io.Writer，同时向标准输出框和write-end pip(log rotator初始化的)输出。
// TODO 也许可以用io.MultiWriter(writer1, writer2)实现
type logWriter struct {}
func (logWriter) Write(p []byte)(n int,err error){
	os.Stdout.Write(p)
	logRotator.Write(p)
	return len(p),nil
}
func init(){
	test.UseLogger(testLog)
}
// subsystemLoggers maps each subsystem identifier to its associated logger.
var subsystemLoggers = map[string]mylog.Logger{
	"TEST":testLog,
}
// initLogRotator initializes the logging rotater to write logs to logFile and
// create roll files in the same directory.  It must be called before the
// package-global log rotater variables are used.
func initLogRotator(logFile string) {
	logDir, _ := filepath.Split(logFile)
	err := os.MkdirAll(logDir, 0700)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create log directory: %v\n", err)
		os.Exit(1)
	}
	r, err := rotator.New(logFile, 10*1024, false, 3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create file rotator: %v\n", err)
		os.Exit(1)
	}

	logRotator = r
}
// setLogLevel sets the logging level for provided subsystem.  Invalid
// subsystems are ignored.  Uninitialized subsystems are dynamically created as
// needed.
// 设置某一个子系统的日志等级
func setLogLevel(subsystemID string, logLevel string) {
	// Ignore invalid subsystems.
	logger, ok := subsystemLoggers[subsystemID]
	if !ok {
		return
	}

	// Defaults to info if the log level is invalid.
	level, _ := mylog.LevelFromString(logLevel)
	logger.SetLevel(level)
}
// 设置全部子系统的日志等级
func setLogLevels(logLevel string) {
	// Configure all sub-systems with the new logging level.  Dynamically
	// create loggers as needed.
	for subsystemID := range subsystemLoggers {
		setLogLevel(subsystemID, logLevel)
	}
}
