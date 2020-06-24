package utils

import (
	"os"
	"os/signal"
	"strings"

	"github.com/go-kit/kit/log"
)

// TrimWhitespace will remove the tabs and newlines from a string
func TrimWhitespace(str string) string {
	r := strings.Replace(str, "\t", "", -1)
	r = strings.Replace(str, "\n", "", -1)
	return r
}

// GetEnvVar ...
func GetEnvVar(key string, def string) (string, bool) {
	variable := os.Getenv(key)
	if variable == "" {
		return def, false
	}
	return variable, true
}

// GetLogger will return a gokit logger with default params
func GetLogger(serviceName string) log.Logger {
	lgr := log.NewLogfmtLogger(os.Stderr)
	lgr = log.WithPrefix(lgr, "service", serviceName)
	lgr = log.WithPrefix(lgr, "timestamp", log.DefaultTimestampUTC)
	return lgr
}

// GetDoneChannel will create a new channel to listen for done signals on
func GetDoneChannel() *chan os.Signal {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	return &done
}
