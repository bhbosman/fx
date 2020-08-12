package fx

type logName struct {
	Name string
}

func LogName(name string) Option {
	return &logName{Name: name}
}

func (self *logName) String() string {
	return "Name logger"
}

func (self *logName) apply(app *App) {
	app.logger.Name = self.Name
}
