package models

import (
	"gorm.io/gorm"
)

type Poke struct {
	Name           string  `json:"name,omitempty" gorm:"column:name"`
	PokdexNumber   int     `json:"pokedex_number,omitempty" gorm:"column:pokedex_number"`
	Classification string  `json:"classification,omitempty" gorm:"column:classification"`
	Height         float32 `json:"height,omitempty" gorm:"column:height"` // in m
	Weight         float32 `json:"weight,omitempty" gorm:"column:weight"` // in kg
	Generation     int     `json:"generation,omitempty" gorm:"column:generation"`
	Type1          string  `json:"type1,omitempty" gorm:"column:type1"`
	Type2          string  `json:"type2,omitempty" gorm:"column:type2;default:NULL"`
}

type Pokemon struct {
	gorm.Model
	Poke
}
