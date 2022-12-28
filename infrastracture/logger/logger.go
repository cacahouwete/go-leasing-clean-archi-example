package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// New create new logger depending on level given.
func New(level string) *zerolog.Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "panic":
		l = zerolog.PanicLevel
	case "fatal":
		l = zerolog.FatalLevel
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	case "trace":
		l = zerolog.TraceLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3

	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &logger
}
