package meshlog

import (
	"context"

	"golang.org/x/exp/slog"
)

var ErrControllerCode = "11071"

type Controller struct {
	enabled bool
	base    *LogrusHandler
}

func (c Controller) Info(ctx context.Context, d description, i ...interface{}) {
	var msg string
	c.base.Info(msg)
}

func (c Controller) Debug(ctx context.Context, d description, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) Warn(err error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) Error(err error) {
	// var msg string
	// return c.base.Error(ErrController(err, msg))
}

func (c Controller) Enabled(level slog.Level) bool {
	return c.enabled
}

func (c Controller) Handle(r slog.Record) error {
	return nil
}

func (c Controller) WithAttrs(attrs []slog.Attr) slog.Handler {
	c.base.WithGroup("controller")
	return c
}

func (c Controller) WithGroup(name string) slog.Handler {
	return c.base.WithGroup("controller")
}

func ErrController(err error, msg string) error {
	return nil
}

func (l *Logger) ControllerLogger() Interface {
	return &Controller{
		enabled: true,
		// base:    l,
	}
}
