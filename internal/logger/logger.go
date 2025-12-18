package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger returns a Zap Logger.
// It checks the APP_ENV environment variable.
func InitLogger() (*zap.Logger, error) {
	env := os.Getenv("APP_ENV") // "dev" or "prod"

	var config zap.Config

	if env == "dev" {
		// Development: Human-readable, colored, stack traces on warnings
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		// Production: JSON format, faster, no colors
		config = zap.NewProductionConfig()
		// Make the timestamp readable (ISO8601) instead of Unix epoch
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
