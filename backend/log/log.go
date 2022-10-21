package log

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initialize zerolog logger.
func Init() {
	// output console when develop local machine
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func Info(ctx context.Context) *zerolog.Event {
	return log.Info()
}
