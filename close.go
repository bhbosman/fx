package fx

//func (self *App) provideAppCloser() IAppCloser {
//	return &appCloser{
//		app: self,
//	}
//}

//type IAppCloser interface {
//	io.Closer
//}

//type appCloser struct {
//	app *App
//}

//func (self *appCloser) Close() error {
//	app := self.app
//	self.app = nil
//
//	stopCtx, cancel := context.WithTimeout(context.Background(), app.StopTimeout())
//	defer cancel()
//
//	if err := app.Stop(stopCtx); err != nil {
//		//app.logger.Fatalf("ERROR\t\tFailed to stop cleanly: %v", err)
//		return err
//	}
//	return nil
//}
