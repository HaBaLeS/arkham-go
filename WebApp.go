package main

import (
	"arkham-go/web"
	"github.com/gin-gonic/gin"
)

type WebApp struct {
	Gin *gin.Engine
}

func NewWebApp(app *App) *WebApp {
	return &WebApp{
		Gin: web.NewServer(app.Runtime.PlaySession),
	}
}

func (web *WebApp) Init(app *App) {
	if err := web.Gin.SetTrustedProxies([]string{}); err != nil {
		panic(err)
	}
}
