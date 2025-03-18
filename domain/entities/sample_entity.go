package entities

import "gorm.io/gorm"

type SampleEntity struct {
	gorm.Model
	Name string `json:"name"`
}