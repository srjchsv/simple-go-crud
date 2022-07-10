package router

import (
	"github.com/gin-gonic/gin"
	"github.com/srjchsv/simple-go-crud/internal/handlers"
	"github.com/thinkerou/favicon"
)

func Register() *gin.Engine {
	r := gin.Default()
	r.Static("/styles", "./templates/styles")
	r.Use(favicon.New("./templates/resources/favicon.ico"))
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", handlers.Home)
	r.POST("/new", handlers.CreateUser)
	r.POST("/delete/:id", handlers.DeleteUser)

	return r
}
