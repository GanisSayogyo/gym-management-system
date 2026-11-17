package config

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger membuat logger yang menulis ke terminal
// dan file logs/app.log secara bersamaan.
func NewLogger() *slog.Logger {
	// Membuat folder logs jika belum tersedia.
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic("gagal membuat folder logs: " + err.Error())
	}

	// Konfigurasi rotasi file log.
	rotator := &lumberjack.Logger{
		Filename:   filepath.Join("logs", "app.log"),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     14,
		Compress:   true,
	}

	// Mengirim log ke terminal dan file.
	writer := io.MultiWriter(os.Stdout, rotator)

	// Membuat handler dengan format JSON.
	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: parseLevel(GetEnv("LOG_LEVEL", "info")),
	})

	logger := slog.New(handler)

	// Menjadikan logger sebagai logger default.
	slog.SetDefault(logger)

	return logger
}

// parseLevel menentukan tingkat log.
func parseLevel(value string) slog.Level {
	switch strings.ToLower(value) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
