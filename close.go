package fx

import (
	"context"
)


func (self *App) provideAppCloser() IAppCloser{
	return &appCloser{
		app: self,
	}
}




type IAppCloser interface {
	Close(ctx context.Context) error
}

type appCloser struct {
	app *App
}



func (self *appCloser) Close(ctx context.Context) error {
	return self.app.Stop(ctx)
}
