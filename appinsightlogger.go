package main

import (
	"fmt"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
)

type AppInsigtsLogger struct {
	client appinsights.TelemetryClient
}

func InitLogger(instrumentationKey string) AppInsigtsLogger {
	fmt.Println("Init logger")
	return AppInsigtsLogger{appinsights.NewTelemetryClient(instrumentationKey)}
}

//
func ShutdownLogger(logger AppInsigtsLogger) {
	fmt.Println("Closing channel...")
	logger.client.Channel().Close(10 * time.Second)

	fmt.Println("30 seconds timeout")
	time.Sleep(30 * time.Second)
}

func ToString(format interface{}) string {
	return fmt.Sprintf("%v", format)
}

// implemenet https://github.com/nuclio/logger/blob/master/logger.go interface
// for now does not handle varagrs

func (logger AppInsigtsLogger) Error(format interface{}, vars ...interface{}) {
	message := ToString(format)
	fmt.Println(message)
	telemetry := appinsights.NewTraceTelemetry(message, appinsights.Error)
	logger.client.Track(telemetry)
}

func (logger AppInsigtsLogger) Warn(format interface{}, vars ...interface{}) {
	message := ToString(format)
	telemetry := appinsights.NewTraceTelemetry(message, appinsights.Warning)
	logger.client.Track(telemetry)
}

func (logger AppInsigtsLogger) Info(format interface{}, vars ...interface{}) {
	message := ToString(format)
	telemetry := appinsights.NewTraceTelemetry(message, appinsights.Information)
	logger.client.Track(telemetry)
}

func (logger AppInsigtsLogger) Debug(format interface{}, vars ...interface{}) {
	// debug will use the *Verbose* severity level
	message := ToString(format)
	telemetry := appinsights.NewTraceTelemetry(message, appinsights.Verbose)
	logger.client.Track(telemetry)
}

func (logger AppInsigtsLogger) ErrorWith(format interface{}, vars ...interface{}) {
	message := ToString(format)
	telemetry := appinsights.NewTraceTelemetry(message, appinsights.Error)
	logger.client.Track(telemetry)
}

func (logger AppInsigtsLogger) WarnWith(format interface{}, vars ...interface{}) {
	message := ToString(format)
	telemetry := appinsights.NewTraceTelemetry(message, appinsights.Warning)
	logger.client.Track(telemetry)
}

func (logger AppInsigtsLogger) InfoWith(format interface{}, vars ...interface{}) {
	message := ToString(format)
	telemetry := appinsights.NewTraceTelemetry(message, appinsights.Information)
	logger.client.Track(telemetry)
}

func (logger AppInsigtsLogger) DebugWith(format interface{}, vars ...interface{}) {
	message := ToString(format)
	telemetry := appinsights.NewTraceTelemetry(message, appinsights.Verbose)
	logger.client.Track(telemetry)
}

// Flush flushes buffered logs
func (logger AppInsigtsLogger) Flush() {
	logger.client.Channel().Flush()
}

// GetChild returns a child logger, if underlying logger supports hierarchal logging
func (logger AppInsigtsLogger) GetChild(name string) {
	return
}

var logger AppInsigtsLogger

func main() {

	logger = InitLogger("<app insight instrumentation key>")

	logger.Error(fmt.Sprintf("%s Error message", time.Now().String()))
	logger.Warn(fmt.Sprintf("%s Warn message", time.Now().String()))
	logger.Info(fmt.Sprintf("%s Info message", time.Now().String()))
	logger.Debug(fmt.Sprintf("%s Debug message", time.Now().String()))

	ShutdownLogger(logger)
}
