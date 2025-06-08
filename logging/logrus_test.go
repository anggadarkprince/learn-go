package logging

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogging(t *testing.T) {
	// Initialize the logger
	logger := logrus.New()

	// Log a message
	fmt.Println("This is a simple log message")
	logger.Println("This is a simple log message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}

func TestLoggingLevel(t *testing.T) {
	// Initialize the logger
	logger := logrus.New()

	// Set the log level to Info
	logger.SetLevel(logrus.InfoLevel)

	// Log messages at different levels
	logger.Trace("This is a trace message") // This will not be logged
	logger.Debug("This is a debug message") // This will not be logged
	logger.Info("This is an info message")   // This will be logged
	logger.Warn("This is a warning message") // This will be logged
	logger.Error("This is an error message") // This will be logged
	//logger.Fatal("This is an fatal message") // This will be logged (Call os exit after logging)
	//logger.Panic("This is an panic message") // This will be logged (Call panic() after logging)
}

func TestOutput(t *testing.T) {
	// Initialize the logger
	logger := logrus.New()

	// Set the output to a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	logger.SetOutput(file)

	// Log a message
	logger.Info("This message will be logged to the file")
}

func TestJsonFormatter(t *testing.T) {
	// Initialize the logger
	logger := logrus.New()

	// Set the formatter to JSON
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Log a message
	logger.Info("This is a JSON formatted log message")
	logger.WithField("username", "angga.ari").
		WithField("name", "Angga Ari").
		Info("User logged in")
	logger.WithFields(logrus.Fields{
		"event": "test_event",
		"user":  "test_user",
	}).Info("This is a JSON formatted log message")
}

func TestEntry(t *testing.T) {
	// Initialize the logger
	logger := logrus.New()

	// Create a new entry
	entry := logrus.NewEntry(logger)

	// Log a message using the entry
	entry.Info("This is a log message using an entry")
	entry.WithField("username", "angga.ari").Info("User logged in")
	entry.WithFields(logrus.Fields{
		"event": "test_event",
		"user":  "test_user",
	}).Info("This is a log message using an entry with fields")
}

type SampleHook struct {

}

func (hook *SampleHook) Levels() []logrus.Level {
	// Trigger the hook for Info and Warn levels
	// You can customize this to include other levels as needed
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.WarnLevel,
	}
}

func (hook *SampleHook) Fire(entry *logrus.Entry) error {
	// Custom logic for the hook
	// For example, you can log to a different output or perform some action
	// Store the entry in a database, send an email, etc.
	fmt.Printf("Sample hook fired: %s - %s (%s)\n", entry.Level, entry.Message, entry.Time)
	return nil
}

func TestHook(t *testing.T) {
	// Initialize the logger
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	// Log messages
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message") // This will not trigger the hook
}

func TestSingleton(t *testing.T) {
	// Initialize the logger
	logrus.Info("This is a simple log message")
	logrus.Error("This is an error message")

	// Set the formatter to JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Log a message
	logrus.Info("This is a JSON formatted log message")
}