package fx

import "go.uber.org/fx/internal/fxlog"

type ILogger interface {
	Printf(format string, v ...interface{})
}

type logger struct {
	logger *fxlog.Logger
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}


func (app *App) provideLogger() ILogger {
	return &logger{
		logger: app.logger,
	}
}
