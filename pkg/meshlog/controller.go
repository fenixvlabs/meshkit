package meshlog

import (
	"context"
	"github.com/fenixvlabs/meshkit/pkg/errors"
	"golang.org/x/exp/slog"
)

var ErrControllerCode = "11071"

type Controller struct {
	enabled bool
	// using logrus logger now
	base *Logger
}

func (c Controller) Info(ctx context.Context, d description, i ...interface{}) {
	c.base.Info()
}

func (c Controller) Debug(ctx context.Context, d description, i ...interface{}) {
	c.base.Debug()
}

func (c Controller) Warn(err error) {
	c.base.Warn(err)
}

func ErrController(err error, msg string) error {
	return errors.NewErrorDescription(ErrControllerCode, errors.Alert, []string{msg}, []string{err.Error()}, []string{}, []string{})
}

func (c Controller) Error(err error, msg string) {
	c.base.Error(ErrController(err, msg))
}

func (c Controller) Enabled(level slog.Level) bool {
	return c.enabled
}

func (c Controller) Handle(r slog.Record) error {
	return nil
}

/*
func (c Controller) WithAttrs(attrs []slog.Attr) slog.Handler {
	c.base.WithGroup("controller")
	return c
}

func (c Controller) WithGroup(name string) slog.Handler {
	return c.base.WithGroup("controller")
}
*/

func (l *Logger) ControllerLogger() slog.Logger {
	return slog.Logger{}
}
