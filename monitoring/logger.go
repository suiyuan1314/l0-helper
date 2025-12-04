package monitoring

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// LogLevel represents the severity level of a log entry
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

// String returns the string representation of a LogLevel
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Logger represents a logger for the bot
type Logger struct {
	mu      sync.RWMutex
	writer  io.Writer
	level   LogLevel
	useJSON bool
}

// NewLoggerParams represents parameters for creating a new logger
type NewLoggerParams struct {
	Level   LogLevel
	UseJSON bool
	Output  io.Writer
}

// NewLogger creates a new logger instance
func NewLogger(params NewLoggerParams) (*Logger, error) {
	// Use stdout if no output is specified
	output := params.Output
	if output == nil {
		output = os.Stdout
	}

	return &Logger{
		writer:  output,
		level:   params.Level,
		useJSON: params.UseJSON,
	}, nil
}

// LogEntry represents a log entry
type LogEntry struct {
	Timestamp    time.Time         `json:"timestamp"`
	Level        string            `json:"level"`
	Message      string            `json:"message"`
	ChainName    string            `json:"chain_name,omitempty"`
	Transaction  *types.Transaction `json:"transaction,omitempty"`
	Address      common.Address    `json:"address,omitempty"`
	Amount       string            `json:"amount,omitempty"`
	Error        string            `json:"error,omitempty"`
	ExtraFields  map[string]interface{} `json:"extra_fields,omitempty"`
}

// Debug logs a debug message
func (l *Logger) Debug(message string, opts ...LogOption) {
	l.log(LogLevelDebug, message, opts...)
}

// Info logs an info message
func (l *Logger) Info(message string, opts ...LogOption) {
	l.log(LogLevelInfo, message, opts...)
}

// Warn logs a warning message
func (l *Logger) Warn(message string, opts ...LogOption) {
	l.log(LogLevelWarn, message, opts...)
}

// Error logs an error message
func (l *Logger) Error(message string, opts ...LogOption) {
	l.log(LogLevelError, message, opts...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(message string, opts ...LogOption) {
	l.log(LogLevelFatal, message, opts...)
	os.Exit(1)
}

// log logs a message with the specified level
func (l *Logger) log(level LogLevel, message string, opts ...LogOption) {
	// Check if the level is enabled
	if level < l.level {
		return
	}

	// Create log entry
	entry := LogEntry{
		Timestamp:   time.Now(),
		Level:       level.String(),
		Message:     message,
		ExtraFields: make(map[string]interface{}),
	}

	// Apply options
	for _, opt := range opts {
		opt(&entry)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.useJSON {
		// Log in JSON format
		jsonData, err := json.Marshal(entry)
		if err != nil {
			// Fallback to text format if JSON marshalling fails
			fmt.Fprintf(l.writer, "%s %s [ERROR] Failed to marshal JSON log entry: %v\n",
				entry.Timestamp.Format(time.RFC3339),
				entry.Level,
				err)
			return
		}
		fmt.Fprintf(l.writer, "%s\n", jsonData)
	} else {
		// Log in text format
		baseFormat := "%s %s %s"
		args := []interface{}{entry.Timestamp.Format(time.RFC3339), entry.Level, entry.Message}

		// Add optional fields
		if entry.ChainName != "" {
			baseFormat += " chain=%s"
			args = append(args, entry.ChainName)
		}
		if entry.Transaction != nil {
			baseFormat += " tx=%s"
			args = append(args, entry.Transaction.Hash().Hex())
		}
		// Compare to zero address using explicit zero value
		zeroAddr := common.Address{}
		if entry.Address != zeroAddr {
			baseFormat += " addr=%s"
			args = append(args, entry.Address.Hex())
		}
		if entry.Amount != "" {
			baseFormat += " amount=%s"
			args = append(args, entry.Amount)
		}
		if entry.Error != "" {
			baseFormat += " error=%s"
			args = append(args, entry.Error)
		}
		baseFormat += "\n"

		fmt.Fprintf(l.writer, baseFormat, args...)
	}
}

// LogOption represents an option for customizing a log entry
type LogOption func(*LogEntry)

// WithChainName adds the chain name to the log entry
func WithChainName(chainName string) LogOption {
	return func(e *LogEntry) {
		e.ChainName = chainName
	}
}

// WithTransaction adds the transaction to the log entry
func WithTransaction(tx *types.Transaction) LogOption {
	return func(e *LogEntry) {
		e.Transaction = tx
	}
}

// WithAddress adds the address to the log entry
func WithAddress(addr common.Address) LogOption {
	return func(e *LogEntry) {
		e.Address = addr
	}
}

// WithAmount adds the amount to the log entry
func WithAmount(amount string) LogOption {
	return func(e *LogEntry) {
		e.Amount = amount
	}
}

// WithError adds the error to the log entry
func WithError(err error) LogOption {
	return func(e *LogEntry) {
		if err != nil {
			e.Error = err.Error()
		}
	}
}

// WithExtraField adds an extra field to the log entry
func WithExtraField(key string, value interface{}) LogOption {
	return func(e *LogEntry) {
		e.ExtraFields[key] = value
	}
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.level = level
}

// GetLevel returns the current log level
func (l *Logger) GetLevel() LogLevel {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.level
}
