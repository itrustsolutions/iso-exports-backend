package logger

import (
	"os"
	"time"

	"github.com/itrustsolutions/iso-exports-backend/utils/config"
	"github.com/rs/zerolog"
)

func Initialize() (zerolog.Logger, error) {
	appConfig := config.GetConfigOrExist()

	var logger zerolog.Logger

	if appConfig.App.Env == "local" {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		logger = zerolog.New(output).With().Timestamp().Logger()
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	return logger, nil
}
