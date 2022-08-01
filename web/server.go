package web

import (
	"arkham-go/runtime"
	"arkham-go/web/templates"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func NewServer(ps *runtime.PlaySession) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	templ := template.Must(template.New("").ParseFS(templates.Templates, "*.gohtml"))
	r.SetHTMLTemplate(templ)
	r.StaticFS("/img", http.Dir("leech-img"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.gohtml", 0)
	})
	return r
}
