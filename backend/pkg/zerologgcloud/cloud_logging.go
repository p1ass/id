package zerologgcloud

import (
	"time"

	"github.com/rs/zerolog"
)

// SetCloudLoggingFieldFormat sets zerolog field names that are compatible with the Cloud Logging format.
func SetCloudLoggingFieldFormat() {
	logLevelSeverity := map[zerolog.Level]string{
		zerolog.DebugLevel: "DEBUG",
		zerolog.InfoLevel:  "INFO",
		zerolog.WarnLevel:  "WARNING",
		zerolog.ErrorLevel: "ERROR",
		zerolog.PanicLevel: "CRITICAL",
		zerolog.FatalLevel: "CRITICAL",
	}

	zerolog.LevelFieldName = "severity"
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return logLevelSeverity[l]
	}

	// https://cloud.google.com/logging/docs/agent/logging/configuration?hl=ja#timestamp-processing
	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}
