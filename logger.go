package fx

//import (
//	"fmt"
//	"go.uber.org/fx/internal/fxlog"
//	"io"
//)
//
//type ILogger interface {
//	io.Closer
//	Printf(format string, v ...interface{})
//	Error(s string, err error)
//	NameChange(newName string)
//}
//
//type logger struct {
//	logger *fxlog.Logger
//}
//
//func (l *logger) NameChange(newName string) {
//	local := l.logger
//	if local != nil {
//		local.Printf(fmt.Sprintf("Change name from %v to %v", local.Name, newName))
//		local.Name = newName
//	}
//}
//
//func (l *logger) Close() error {
//	l.logger = nil
//	return nil
//}
//
//func (l *logger) Error(s string, err error) {
//	local := l.logger
//	if local != nil {
//		local.Printf("ERROR: %v(%v)", s, err.Error())
//	}
//}
//
//func (l *logger) Printf(format string, v ...interface{}) {
//	local := l.logger
//	if local != nil {
//		local.Printf(format, v...)
//	}
//}
//
//func (app *App) provideLogger() ILogger {
//	return &logger{
//		logger: app.logger,
//	}
//}
