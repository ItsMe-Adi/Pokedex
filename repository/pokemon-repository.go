package repository

import (
	"pokedex/models"

	"gorm.io/gorm"
)

var Db *gorm.DB

func CreatePokemonTableWithRecords(records []models.Pokemon) {
	Db.AutoMigrate(&models.Pokemon{})
	Db.Create(records)
}

func GetPokeList(limit int, offset int) []models.Poke {
	list := []models.Poke{}
	Db.Limit(limit).Offset(offset).Select("name", "pokedex_number", "type1", "type2").Find(&list)
	return list
}

func GetPokeDetails(id int) models.Pokemon {
	details := models.Pokemon{}
	Db.First(&details, id)
	return details
}
