package main

import (
	"flag"
	"os"
	"time"

	errorly "github.com/TheRockettek/Errorly-Web/internal"
	"github.com/rs/zerolog"
)

func main() {
	lFlag := flag.String("level", "info", "Log level to use (debug/info/warn/error/fatal/panic/no/disabled/trace)")
	removeStaleEntries := flag.Bool("remove-stale-entries", false, "Purces unused integrations, issues and comments")

	flag.Parse()

	level, err := zerolog.ParseLevel(*lFlag)
	if err != nil {
		level = zerolog.InfoLevel
	}

	logger := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.Stamp,
	}

	log := zerolog.New(logger).With().Timestamp().Logger()

	if level != zerolog.NoLevel {
		log.Info().Str("logLevel", level.String()).Msg("Using logging")
	}

	errorly, err := errorly.NewErrorly(logger, level)
	if err != nil {
		log.Panic().Err(err).Msgf("Cannot create errorly: %s", err)
	}

	err = errorly.Open(*removeStaleEntries)
	if err != nil {
		log.Panic().Err(err).Msgf("Cannot open errorly: %s", err)
	}
}
