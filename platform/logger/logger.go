package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Logger interface {
	Info(ctx context.Context, messge string, fields ...zap.Field)
	Error(ctx context.Context, message string, fields ...zap.Field)
	Fatal(ctx context.Context, message string, fields ...zap.Field)
	Warn(ctx context.Context, message string, fields ...zap.Field)
	Panic(ctx context.Context, message string, fields ...zap.Field)
	Debug(ctx context.Context, message string, fields ...zap.Field)
	GetZapLogger() *zap.Logger
	Named(name string) *logger
	With(fields ...zap.Field) *logger
	with(fields ...zap.Field) *zap.Logger
	extract(ctx context.Context) []zap.Field
}

type logger struct {
	log *zap.Logger
}

func NewLogger(log *zap.Logger) Logger {
	return logger{
		log: log,
	}
}

// Error implements Logger.
func (l logger) Error(ctx context.Context, message string, fields ...zap.Field) {
	l.with(l.extract(ctx)...).Error(message, fields...)
}

// Fatal implements Logger.
func (l logger) Fatal(ctx context.Context, message string, fields ...zap.Field) {
	l.with(l.extract(ctx)...).Fatal(message, fields...)
}

// Info implements Logger.
func (l logger) Info(ctx context.Context, messge string, fields ...zap.Field) {
	l.with(l.extract(ctx)...).Info(messge, fields...)
}

// GetZapLogger implements Logger.
func (l logger) GetZapLogger() *zap.Logger {
	return l.log
}

// Debug implements Logger.
func (l logger) Debug(ctx context.Context, message string, fields ...zap.Field) {
	l.with(l.extract(ctx)...).Debug(message, fields...)
}

// Named implements Logger.
func (l logger) Named(name string) *logger {
	log := l.log.Named(name)
	return &logger{
		log: log,
	}
}

// Panic implements Logger.
func (l logger) Panic(ctx context.Context, message string, fields ...zap.Field) {
	l.with(l.extract(ctx)...).Panic(message, fields...)
}

// Warn implements Logger.
func (l logger) Warn(ctx context.Context, message string, fields ...zap.Field) {
	l.with(l.extract(ctx)...).Warn(message, fields...)
}

// extract implements Logger.
func (l logger) extract(ctx context.Context) []zap.Field {
	var fields []zap.Field

	fields = append(fields, zap.String("time", time.Now().Format(time.RFC3339)))

	if val, ok := ctx.Value("user-id").(string); ok {
		fields = append(fields, zap.String("user_id", val))
	}

	if val, ok := ctx.Value("x-request-id").(string); ok {
		fields = append(fields, zap.String("x-request-id", val))
	}

	if val, ok := ctx.Value("x-request-start-time").(time.Time); ok {
		fields = append(fields, zap.Float64("x-request-start-time", float64(time.Since(val).Milliseconds())))
	}

	return fields
}

// With implements Logger.
func (l logger) With(fields ...zap.Field) *logger {
	log := l.log.With(fields...)
	return &logger{
		log: log,
	}
}

// With implements Logger.
func (l logger) with(fields ...zap.Field) *zap.Logger {
	log := l.log.With(fields...)
	return log
}
