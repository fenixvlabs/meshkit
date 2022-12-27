package meshlog

import (
	"context"
	"github.com/fenixvlabs/meshkit/pkg/errors"
	log "github.com/sirupsen/logrus"

	"golang.org/x/exp/slog"
	"io"
	"os"
	"time"
)

const (
	AdapterMeshkit    = "meshkit"
	AdapterMesheryctl = "mesheryctl"
	AdapterConsul     = "meshery-consul"
)

const (
	JsonLogFormat = iota
	SyslogLogFormat
	TerminalLogFormat
)

type description string
type Format int

type Writer interface {
	Printf(string, ...interface{})
}

type Interface interface {
	Info(context.Context, description, ...interface{})
	Debug(context.Context, description, ...interface{})
	Warn(err error)
	Error(err error)
}

type commonHandler struct {
	Output io.Writer
	w      Writer
}

type HandlerOptions struct {
	Format Format
	Level  bool
	commonHandler
}

type LogrusHandler struct {
	Logger *log.Logger
	Writer
	commonHandler
}

type Logger struct {
	handler *log.Entry
}

func (l LogrusHandler) Enabled(level slog.Level) bool {
	return true
}

func (l LogrusHandler) Handle(r slog.Record) error {
	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) {
		fields[a.Key] = a.Value.Any()
	})

	entry := l.Logger.WithFields(fields)

	switch r.Level {
	case slog.LevelDebug:
		entry.Debug(r.Message)
	case slog.LevelInfo.Level():
		entry.Info(r.Message)
	case slog.LevelWarn:
		entry.Warn(r.Message)
	case slog.LevelError:
		entry.Error(r.Message)
	}

	return nil
}

func (l LogrusHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return l
}

func (l LogrusHandler) WithGroup(name string) slog.Handler {
	return l
}

func (l LogrusHandler) Error(err error) string {
	return ""
}

func (l LogrusHandler) Info(msg string) {
	return

}

type TerminalFormatter struct{}

func (t TerminalFormatter) Format(entry *log.Entry) ([]byte, error) {
	return append([]byte(entry.Message), '\n'), nil
}

func New(appName string, opts HandlerOptions) (Interface, error) {
	meshlog := log.New()

	switch opts.Format {
	case JsonLogFormat:
		meshlog.SetFormatter(&log.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	case SyslogLogFormat:
		meshlog.SetFormatter(&log.TextFormatter{
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
		})
	case TerminalLogFormat:
		meshlog.SetFormatter(new(TerminalFormatter))
	}

	log.SetOutput(os.Stdout)
	if opts.Output != nil {
		log.SetOutput(opts.Output)
	}

	log.SetLevel(log.InfoLevel)
	if opts.Level {
		log.SetLevel(log.DebugLevel)
	}

	_ = log.WithFields(log.Fields{"app": appName})
	return nil, nil
}

func (l *Logger) Info(description ...interface{}) {
	l.Info(log.InfoLevel, description)
}

func (l *Logger) Debug(description ...interface{}) {
	l.Debug(log.DebugLevel, description)
}

func (l *Logger) Error(err error) {
	l.handler.WithFields(log.Fields{
		"code":                  errors.GetCode(err),
		"severity":              errors.GetSeverity(err),
		"short-description":     errors.GetSDescription(err),
		"probable-cause":        errors.GetCause(err),
		"suggested-remediation": errors.GetRemedy(err),
	}).Log(log.ErrorLevel, err.Error())
}

func (l *Logger) Warn(err error) {
	l.handler.WithFields(log.Fields{
		"code":                  errors.GetCode(err),
		"severity":              errors.GetSeverity(err),
		"short-description":     errors.GetSDescription(err),
		"probable-cause":        errors.GetCause(err),
		"suggested-remediation": errors.GetRemedy(err),
	}).Log(log.WarnLevel, err.Error())
}
