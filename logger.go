package context

import (
	"fmt"

	"runtime"

	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
	"golang.org/x/net/context"
)

type loggerKey struct{}

// Logger provides a leveled-logging interface.
type Logger interface {
	log.Interface
}

// WithLogger creates a new context with provided logger.
func WithLogger(ctx context.Context, logger Logger) context.Context {
	return WithValue(ctx, loggerKey{}, logger)
}

// GetLoggerWithField returns a logger instance with the specified field key
// and value without affecting the context. Extra specified keys will be
// resolved from the context.
func GetLoggerWithField(ctx context.Context, key, value interface{}, keys ...interface{}) Logger {
	return getApexLogger(ctx, keys...).WithField(fmt.Sprint(key), value)
}

// GetLoggerWithFields returns a logger instance with the specified fields
// without affecting the context. Extra specified keys will be resolved from
// the context.
func GetLoggerWithFields(ctx context.Context, fields map[interface{}]interface{}, keys ...interface{}) Logger {
	// must convert from interface{} -> interface{} to string -> interface{} for logrus.
	lfields := make(log.Fields, len(fields))
	for key, value := range fields {
		lfields[fmt.Sprint(key)] = value
	}

	return getApexLogger(ctx, keys...).WithFields(lfields)
}

// GetLogger returns the logger from the current context, if present. If one
// or more keys are provided, they will be resolved on the context and
// included in the logger. While context.Value takes an interface, any key
// argument passed to GetLogger will be passed to fmt.Sprint when expanded as
// a logging key field. If context keys are integer constants, for example,
// its recommended that a String method is implemented.
func GetLogger(ctx context.Context, keys ...interface{}) Logger {
	return getApexLogger(ctx, keys...)
}

// getApexLogger returns the apex logger for the context. If one more keys
// are provided, they will be resolved on the context and included in the
// logger. Only use this function if specific apex/log functionality is
// required.
func getApexLogger(ctx context.Context, keys ...interface{}) *log.Entry {
	var logger *log.Entry

	// Get a logger, if it is present.
	loggerInterface := ctx.Value(loggerKey{})
	if loggerInterface != nil {
		if lgr, ok := loggerInterface.(*log.Entry); ok {
			logger = lgr
		}
	}

	if logger == nil {
		fields := log.Fields{}

		// Fill in the instance id, if we have it.
		instanceID := ctx.Value("instance.id")
		if instanceID != nil {
			fields["instance.id"] = instanceID
		}

		fields["go.version"] = runtime.Version()
		// If no logger is found, just return the standard logger.
		// TODO: (apex/log) has nil handler for standard logger
		// https://github.com/apex/log/issues/6
		if alog, ok := log.Log.(*log.Logger); ok && alog.Handler == nil {
			log.SetHandler(discard.Default)
		}
		logger = log.Log.WithFields(fields)

	}

	fields := log.Fields{}
	for _, key := range keys {
		v := ctx.Value(key)
		if v != nil {
			fields[fmt.Sprint(key)] = v
		}
	}

	return logger.WithFields(fields)
}
