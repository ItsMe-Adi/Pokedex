package controller

import (
	"net/http"
	"pokedex/controller/constants"
	"pokedex/helpers"
	"pokedex/models"
	"strconv"

	"pokedex/service"

	"github.com/gin-gonic/gin"
)

func MigrateAndPopulateTable(ctx *gin.Context) {
	data := service.MigrateAndPopulateTable()
	ctx.JSON(http.StatusOK, helpers.StandardResponse(true, gin.H{"saved_data": data}, "Table successfully created with data"))
}

func GetPokedexView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func GetPokeList(ctx *gin.Context) {
	var request models.Listrequest
	if ctx.Query("limit") != "" {
		request.Limit, _ = strconv.Atoi(ctx.Query("limit"))
	} else {
		request.Limit = constants.Limit
	}
	if ctx.Query("page") != "" {
		request.Page, _ = strconv.Atoi(ctx.Query("page"))
	} else {
		request.Page = constants.Page
	}
	pokeList := service.GetPokeList(request)
	ctx.JSON(http.StatusOK, helpers.StandardResponse(true, gin.H{"list": pokeList}, "Pokemon List"))
}

func GetPokeDetails(ctx *gin.Context) {
	var request models.Detailsrequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helpers.StandardResponse(false, gin.H{}, err.Error()))
	} else {
		pokeDetails := service.GetPokeDetails(request)
		ctx.JSON(http.StatusOK, helpers.StandardResponse(true, gin.H{"details": pokeDetails}, "Pokemon Details"))
	}
}

func UpdatePokemon(ctx *gin.Context) {

}
