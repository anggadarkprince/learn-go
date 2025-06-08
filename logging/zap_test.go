package logging

import (
	"fmt"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestZapSugarLog(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "http://example.com")
}

func TestZapLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	logger.Info("Failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Error("Failed to fetch URL", zap.String("url", "http://example.com"))
}

func TestZapDevelopment(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any

	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	// logger.Fatal("This is a fatal message") // This will call os exit after logging
	// logger.Panic("This is a panic message") // This will call panic() after logging
}

func TestZapLogLevel(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Set the log level to Info
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	// logger.Fatal("This is a fatal message") // This will call os exit after logging
	// logger.Panic("This is a panic message") // This will call panic() after logging
}

func TestZapSetLevel(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Set the log level to Info
	logger = logger.WithOptions(zap.IncreaseLevel(zap.WarnLevel))

	logger.Debug("This is a debug message") // This will not be logged
	logger.Info("This is an info message")   // This will not be logged
	logger.Warn("This is a warning message") // This will be logged
	logger.Error("This is an error message") // This will be logged
}

func TestZapJson(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	logger.Info("hello, world",
		zap.String("user", "example_user"),
	)

	logger.Info("Usage Statistics",
		zap.Int("current-memory", 50),
		zap.Int("min-memory", 20),
		zap.Int("max-memory", 80),
		zap.Int("cpu", 10),
		zap.String("app-version", "v0.0.1-beta"),
	)
}

func TestZapFile(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Set the output to a file
	file, err := os.OpenFile("zap.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(file),
		zap.InfoLevel,
	))

	logger.Info("This message will be logged to the file")

	logger.Info("hello, world",
		zap.String("user", "example_user"),
	)
}

func TestZapCustomHook(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Add the hook to the logger
	logger = logger.WithOptions(zap.Hooks(func(entry zapcore.Entry) error {
		fmt.Println("test hooks test hooks")
		return nil
	}))

	logger.Error("This is an error message")
}