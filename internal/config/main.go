package config

import (
	"fmt"
	"time"

	"github.com/pronuu/lincoln/internal/constants"
	"github.com/pronuu/lincoln/internal/utils"
)

// Config ...
type Config struct {
	Port         string
	Host         string
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

// NewConfig ...
func NewConfig() *Config {
	env, _ := utils.GetEnvVar(constants.Environment, "production")

	var portKey string
	if env == "production" {
		portKey = constants.HerokuPort
	} else {
		portKey = constants.LincolnPort
	}

	port, _ := utils.GetEnvVar(portKey, "7998")
	host, _ := utils.GetEnvVar(constants.LincolnAddress, "localhost")
	address := fmt.Sprintf("%s:%s", host, port)

	writeTimeoutStr, _ := utils.GetEnvVar(constants.LincolnWriteTimeout, "15")
	writeTimeout, _ := time.ParseDuration(writeTimeoutStr)

	readTimeoutStr, _ := utils.GetEnvVar(constants.LincolnReadTimeout, "15")
	readTimeout, _ := time.ParseDuration(readTimeoutStr)

	idleTimeoutStr, _ := utils.GetEnvVar(constants.LincolnIdleTimeout, "60")
	idleTimeout, _ := time.ParseDuration(idleTimeoutStr)

	return &Config{
		Port:         port,
		Host:         host,
		Address:      address,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
}
