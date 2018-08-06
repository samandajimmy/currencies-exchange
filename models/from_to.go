package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type FromTo struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
	Rate []Rate `json:"rates,omitempty"`

	gorm.Model
}

func GetAllFromTo() (string, error) {
	var currencies []FromTo
	db.Find(&currencies)

	res, err := json.Marshal(currencies)
	return string(res), err
}

func CreateFromTo(from, to string) {
	db.Create(&FromTo{From: from, To: to})
}

func DeleteFromTo(id string) {
	var fromTo FromTo

	db.Where("id = ?", id).Find(&fromTo)
	db.Delete(&fromTo)
}
