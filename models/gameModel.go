package models

import "gorm.io/gorm"

type Figure struct {
	gorm.Model
	Name  string
	Color string
}

type Move struct {
	gorm.Model
	X        string
	Y        string
	FigureID uint `gorm:"foreignkey:ID"`
	GameID   uint `gorm:"foreignkey:ID"`
}

type Game struct {
	gorm.Model
	Moves   []Move `gorm:"foreignkey:GameID"`
	Players []User `gorm:"many2many:game_players;"`
}
