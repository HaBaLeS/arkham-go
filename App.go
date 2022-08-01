package main

type App struct {
	Runtime *RuntimeApp
	Web     *WebApp
}

func NewApp() *App {
	app := &App{}
	app.Runtime = NewRuntime(app)
	app.Web = NewWebApp(app)
	return app
}

func (app *App) InitApp() {
	app.Runtime.Init(app)
	app.Web.Init(app)
}
