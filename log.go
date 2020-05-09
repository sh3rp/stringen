package stringen

import (
	"os"

	"github.com/rs/zerolog"
)

var LOGGER = zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
