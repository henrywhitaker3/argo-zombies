package logger

import (
	"context"

	"github.com/henrywhitaker3/ctxgen"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	key = "logger"
)

var (
	zl *zap.SugaredLogger
)

func Wrap(ctx context.Context, level zap.AtomicLevel) context.Context {
	return ctxgen.WithValue(ctx, key, newLogger(level))
}

func Logger(ctx context.Context) *zap.SugaredLogger {
	l, ok := ctxgen.ValueOk[*zap.SugaredLogger](ctx, key)
	if !ok {
		return newLogger(zap.NewAtomicLevelAt(zapcore.ErrorLevel))
	}
	return l
}

func newLogger(level zap.AtomicLevel) *zap.SugaredLogger {
	if zl == nil {
		logConfig := zap.NewProductionConfig()
		logConfig.OutputPaths = []string{"stdout"}
		logConfig.Level = level
		logger, _ := logConfig.Build()
		zl = logger.Sugar()
	}
	return zl
}
