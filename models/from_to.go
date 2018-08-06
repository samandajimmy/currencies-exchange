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

func CreateFromTo(from, to string) string {
	if db.Create(&FromTo{From: from, To: to}).Error != nil {
		// Create failed, do something e.g. return, panic etc.
		return "create failed! please try again later"
	}
	return "data has been successfully created!"
}

func DeleteFromTo(id string) string {
	var fromTo FromTo

	err := db.Where("id = ?", id).Find(&fromTo).RecordNotFound()

	if err {
		return "data not found!"
	}

	db.Delete(&fromTo)
	return "data has bean successfully deleted!"
}
