package logging

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestZerolog(t *testing.T){
	// UNIX Time is faster and smaller than most timestamps
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Log a message
	log.Print("hello world")
	log.Info().Msg("hello, world")
	log.Info().Str("user", "example_user").Msg("hello, world")
	log.Info().Str("user", "example_user").Msgf("hello, world %s", "with format")
	log.Info().Fields(map[string]interface{}{
		"user": "example_user",
		"role": "admin",
	}).Msg("hello, world with fields")
}

func TestZerologLevel(t *testing.T) {
	// Set the log level to Info
	log.Info().Msg("This is an info message")
	log.Warn().Msg("This is a warning message")
	log.Error().Msg("This is an error message")
	// log.Fatal().Msg("This is a fatal message") // This will call os exit after logging
	// log.Panic().Msg("This is a panic message") // This will call panic() after logging
}

func TestZeroLogSetLevel(t *testing.T) {
	// Set the log level to Info
	log.Logger = log.Level(zerolog.WarnLevel)

	log.Debug().Msg("This is a debug message") // This will not be logged
	log.Info().Msg("This is an info message")   // This will not be logged
	log.Warn().Msg("This is a warning message") // This will be logged
	log.Error().Msg("This is an error message") // This will be logged
}

func TestZerologJson(t *testing.T) {
	// Log a message in JSON format
	log.Info().Str("user", "example_user").Msg("hello, world")
	log.Info().Fields(map[string]interface{}{
		"user": "example_user",
		"role": "admin",
	}).Msg("hello, world with fields")
}

func TestZerologFile(t *testing.T) {
	// Set the output to a file
	file, err := os.OpenFile("zerolog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	log.Logger = log.Output(file)

	// Log a message
	log.Info().Msg("This message will be logged to the file")
}

func TestZerologHook(t *testing.T) {
	// Create a custom hook
	hook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		if level == zerolog.ErrorLevel {
			e.Str("error", "custom error message")
		}
	})

	// Add the hook to the logger
	log.Logger = log.Hook(hook)

	// Log an error message
	log.Error().Msg("This is an error message with a custom hook")
}