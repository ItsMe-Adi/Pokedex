package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"pokedex/models"
	"pokedex/repository"
	"strconv"
)

func readCsv(path string) [][]string {
	file, err := os.Open(path) // "assets/pokemon.csv"
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	fileReader := csv.NewReader(file)
	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return records
}

func prepareCsvRecords(records [][]string) []models.Pokemon {
	var data []models.Pokemon
	for i := 1; i < len(records); i++ {
		dex_num, _ := strconv.Atoi(records[i][32])
		height, _ := strconv.ParseFloat(records[i][27], 32)
		weight, _ := strconv.ParseFloat(records[i][38], 32)
		gen_num, _ := strconv.Atoi(records[i][39])
		record := models.Pokemon{
			PokdexNumber:   dex_num,
			Name:           records[i][30],
			Classification: records[i][24],
			Height:         float32(height),
			Weight:         float32(weight),
			Generation:     gen_num,
			Type1:          records[i][36],
			Type2:          records[i][37],
		}
		data = append(data, record)
	}
	return data
}

func MigrateAndPopulateTable() []models.Pokemon {
	records := prepareCsvRecords(readCsv("assets/pokemon.csv"))
	repository.CreatePokemonTableWithRecords(records)
	return records
}
