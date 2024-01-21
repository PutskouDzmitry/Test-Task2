package logging

import "context"

type contextLogger struct{}

func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, contextLogger{}, logger)
}

func LoggerFromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(contextLogger{}).(Logger); ok {
		return l
	}
	return *GetLogger()
}
