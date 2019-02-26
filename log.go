package main

import (
	"github.com/btcsuite/btclog"
	"os"
)

const (
	maxRejectReasonLen = 250
)
var (
	// backendLog is the logging backend used to create all subsystem loggers.
	// The backend must not be used before the log rotator has been initialized,
	// or data races and/or nil pointer dereferences will occur.
	backendLog = btclog.NewBackend(logWriter{})

	// logRotator is one of the logging outputs.  It should be closed on
	// application shutdown.
	logRotator *rotator.Rotator

	adxrLog = backendLog.Logger("ADXR")
	amgrLog = backendLog.Logger("AMGR")
	cmgrLog = backendLog.Logger("CMGR")
	bcdbLog = backendLog.Logger("BCDB")
	bmgrLog = backendLog.Logger("BMGR")
	hcdLog  = backendLog.Logger("HC")
	chanLog = backendLog.Logger("CHAN")
	discLog = backendLog.Logger("DISC")
	indxLog = backendLog.Logger("INDX")
	minrLog = backendLog.Logger("MINR")
	peerLog = backendLog.Logger("PEER")
	rpcsLog = backendLog.Logger("RPCS")
	scrpLog = backendLog.Logger("SCRP")
	srvrLog = backendLog.Logger("SRVR")
	stkeLog = backendLog.Logger("STKE")
	txmpLog = backendLog.Logger("TXMP")
)
// logWriter 实现了io.Writer，同时向标准输出框和write-end pip(log rotator初始化的)输出。
// TODO 也许可以用io.MultiWriter(writer1, writer2)实现
type logWriter struct {}
func (logWriter) Write(p []byte)(n int,err error){
	os.Stdout.Write(p)

}