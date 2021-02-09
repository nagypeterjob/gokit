package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

// Logger mock interface
type Logger interface {
	Info(string, ...zap.Field)
	Debug(string, ...zap.Field)
	Error(string, ...zap.Field)
}

// New creates a new zapcore logger instance with production config.
// If `logFilePath` is provided & file does not exist, logger will try to create it.
func New(level string, logFilepath string, stackTrace bool) (*zap.Logger, error) {
	var (
		atom = zap.NewAtomicLevel()
		err  = atom.UnmarshalText([]byte(level))
		cfg  = zap.NewProductionConfig()
	)
	if err != nil {
		return nil, err
	}

	cfg.DisableStacktrace = stackTrace
	cfg.Level = atom

	if logFilepath != "" {
		directory := filepath.Dir(logFilepath)

		if _, err := os.Stat(directory); os.IsNotExist(err) {
			if err := os.MkdirAll(directory, os.ModePerm); err != nil {
				return nil, err
			}
		}

		cfg.OutputPaths = []string{logFilepath}
	}

	return cfg.Build()
}
