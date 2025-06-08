package logging

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

func TestSlog(t *testing.T) {
	slog.Info("hello, world")
	slog.Info("message", "k1", "v1", "k2", "v2")
	slog.Info("hello, world", "user", os.Getenv("USER"))

	logger := slog.Default()
	logger.Info("hello, world", "user", os.Getenv("USER"))
}

func TestSlogText(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("hello, world", "user", os.Getenv("USER"))

	logger.Info("Usage Statistics",
		slog.Int("current-memory", 50),
		slog.Int("min-memory", 20),
		slog.Int("max-memory", 80),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)
}

func TestSlogJson(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello, world", "user", os.Getenv("USER"))

	logger.Info("Usage Statistics",
		slog.Int("current-memory", 50),
		slog.Int("min-memory", 20),
		slog.Int("max-memory", 80),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)
}

func TestSlogAttributes(t *testing.T) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world",
    slog.String("user", os.Getenv("USER")))
}

func TestSlogFile(t *testing.T) {
	file, err := os.OpenFile("slog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	logger := slog.New(slog.NewJSONHandler(file, nil))
	logger.Info("This message will be logged to the file")
}