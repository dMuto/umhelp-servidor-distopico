package logger

import (
	"os"

	"github.com/dMuto/umhelp-servidor-distopico/config"
	"github.com/rs/zerolog"
)

func New(cfg *config.Config) *zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Str("service", cfg.InternalConfig.ServiceName).Timestamp().Logger()

	return &logger
}
