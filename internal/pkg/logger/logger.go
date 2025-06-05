package logger

import (
	"log/slog"
	"os"
	"runtime/debug"
)

type Config struct {
	Level              slog.Level
	IncludeProgramInfo bool
}

func MustLogger(config Config) *slog.Logger {
	opts := slog.HandlerOptions{Level: config.Level}

	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &opts),
	)

	if config.IncludeProgramInfo {
		buildInfo, _ := debug.ReadBuildInfo()

		log = log.With(slog.Group("program_info", "os", os.Getpid(), "go_version", buildInfo.GoVersion))
	}

	slog.SetDefault(log)

	return log
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
