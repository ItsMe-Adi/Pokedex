package service

import (
	"pokedex/helpers"
	"pokedex/models"
	"pokedex/repository"
)

func GetPokeList(request models.Listrequest) []models.Poke {
	return repository.GetPokeList(request.Limit, helpers.GetOffset(request.Page, request.Limit))
}

func GetPokeDetails(request models.Detailsrequest) models.Pokemon {
	return repository.GetPokeDetails(request.Id)
}
