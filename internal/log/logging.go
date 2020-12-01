// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"github.com/sirupsen/logrus"
)

type LogFormat string

const (
	Syslog = "syslog"

	LogFormatText LogFormat = "text"
	LogFormatJSON LogFormat = "json"

	// DefaultLogLevelStr is the string representation of DefaultLogLevel. It
	// is used to allow for injection of the logging level via go's ldflags in
	// unit tests, as only injection with strings via ldflags is allowed.
	DefaultLogLevelStr string = "debug"

	// DefaultLogFormat is the string representation of the default logrus.Formatter
	// we want to use (possible values: text or json)
	DefaultLogFormat LogFormat = LogFormatText
)

var (
	// DefaultLogger is the base logrus logger. It is different from the logrus
	// default to avoid external dependencies from writing out unexpectedly
	DefaultLogger = InitializeDefaultLogger()

	// LevelStringToLogrusLevel maps string representations of logrus.Level into
	// their corresponding logrus.Level.
	LevelStringToLogrusLevel = map[string]logrus.Level{
		"panic":   logrus.PanicLevel,
		"error":   logrus.ErrorLevel,
		"warning": logrus.WarnLevel,
		"info":    logrus.InfoLevel,
		"debug":   logrus.DebugLevel,
	}
)

// LogOptions maps configuration key-value pairs related to logging.
type LogOptions map[string]string

// InitializeDefaultLogger returns a logrus Logger with a custom text formatter.
func InitializeDefaultLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = GetFormatter(DefaultLogFormat)
	logger.SetLevel(LevelStringToLogrusLevel[DefaultLogLevelStr])
	return logger
}

// GetFormatter returns a configured logrus.Formatter with some specific values
// we want to have
func GetFormatter(format LogFormat) logrus.Formatter {
	switch format {
	case LogFormatText:
		return &logrus.TextFormatter{
			DisableTimestamp: true,
			DisableColors:    true,
		}
	case LogFormatJSON:
		return &logrus.JSONFormatter{
			DisableTimestamp: true,
		}
	}

	return nil
}
