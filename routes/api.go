package routes

import (
	"pokedex/controller"
	"pokedex/middlewares"

	"github.com/gin-gonic/gin"
)

func InitializeApiRoutes(server *gin.Engine) {

	api := server.Group("/api")
	{
		api.POST("/migrate", middlewares.BasicAuth(), controller.MigrateAndPopulateTable)

		v1 := api.Group("/v1")
		{
			v1.GET("/list", controller.GetPokeList)

			v1.GET("/details", controller.GetPokeDetails)

			v1.PATCH("/update", controller.UpdatePokemon)
		}
	}
}
