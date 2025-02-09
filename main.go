package main

import (
	"TracerExample/handler"
	"TracerExample/scheduler"
	"fmt"
	"github.com/FedosOnGIT/TracerLib/uploadBatch"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile("logs/tracer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v", err)
		os.Exit(1)
	}

	logger, _ := uploadBatch.New(file, zapcore.WarnLevel, zapcore.ErrorLevel, uploadBatch.Configuration{
		Module:      "Tracer",
		Service:     "Example",
		OsVersion:   "5.15.62-13.el7.x86_64",
		Vendor:      "Linux",
		Host:        "localhost",
		DataCenter:  "datacenter",
		CloudMinion: "runner",
		VersionName: "1.0",
		Environment: "test",
		DeviceID:    "123",
	})

	go scheduler.StartBackgroundTask(logger)

	tracerHandler := handler.New(logger)
	http.HandleFunc("/warn", tracerHandler.HandleWithWarn)
	http.HandleFunc("/warnf", tracerHandler.HandleWithWarnf)
	http.HandleFunc("/error", tracerHandler.HandleWithError)
	http.HandleFunc("/errorf", tracerHandler.HandleWithErrorf)

	logger.Warn("Starting server")

	if err := http.ListenAndServe(":8086", nil); err != nil {
		logger.Fatalf("failed to start server, %s", err)
	}
}
