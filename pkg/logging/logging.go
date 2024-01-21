package logging

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
		if err != nil {
			logrus.Error("failed to write log")
		}
	}

	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) GetLoggerWithFields(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)

	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0777)

	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0770)

	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{file, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	e = logrus.NewEntry(l)
}

func WithField(ctx context.Context, key string, value interface{}) *Logger {
	return &Logger{
		LoggerFromContext(ctx).WithField(key, value),
	}
}

func WithFields(ctx context.Context, fields logrus.Fields) *Logger {
	return &Logger{
		LoggerFromContext(ctx).WithFields(fields),
	}
}

func WithQuery(ctx context.Context, sql string, table string, args []interface{}) *Logger {
	return &Logger{
		LoggerFromContext(ctx).WithFields(map[string]interface{}{
			"sql":   sql,
			"table": table,
			"args":  args,
		}),
	}
}

func WithError(ctx context.Context, err error) *Logger {
	return &Logger{
		LoggerFromContext(ctx).WithError(err),
	}
}
