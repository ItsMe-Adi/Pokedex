package routes

import (
	"pokedex/controller"

	"github.com/gin-gonic/gin"
)

func InitializeViewRoutes(server *gin.Engine) {
	server.GET("/pokedex", controller.GetPokedexView)
}
