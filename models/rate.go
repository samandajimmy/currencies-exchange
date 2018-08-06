package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

const dateFormat = "2006-01-02"

type Rate struct {
	gorm.Model
	Rate     float64 `json:"rate,omitempty"`
	Date     string  `json:"date,omitempty"`
	FromToID uint    `json:"from_to_id,omitempty"`
	FromTo   FromTo  `json:"from_to,omitempty"`
}

type ExchangeRate struct {
	FromToId uint    `json:"from_to_id,omitempty"`
	From     string  `json:"from,omitempty"`
	To       string  `json:"to,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
	Date     string  `json:"date,omitempty"`
	WeekAvg  string  `json:"week_avg,omitempty"`
	Min      float64 `json:"min,omitempty"`
	Max      float64 `json:"max,omitempty"`
	Variance float64 `json:"variance,omitempty"`
}

type RecentRate struct {
	RateDetails ExchangeRate
	ExRates     []ExchangeRate
}

func setFromTo(rates []Rate) []Rate {
	for idx, rate := range rates {
		var fromTo FromTo
		db.Where("id = ?", rate.FromToID).Find(&fromTo)
		rates[idx].FromTo = fromTo
	}
	return rates
}

func GetAllRates() (string, error) {
	var rates []Rate
	db.Find(&rates)

	rates = setFromTo(rates)

	res, err := json.Marshal(rates)
	return string(res), err
}

func GetRatesByDate(date string) (string, error) {
	var exRates []ExchangeRate
	var tempExRates []ExchangeRate
	var rates []Rate
	var tempRates []Rate

	formatedDate, _ := time.Parse(dateFormat, date)
	weekBefore := formatedDate.AddDate(0, 0, -6).Format(dateFormat)

	db.Where("date = ?", date).Find(&rates)

	if len(rates) == 0 {
		return "no rates available!", err
	}

	db.Table("from_tos").
		Select("from_tos.id as from_to_id, avg(rates.rate) as week_avg").
		Joins("left join rates on rates.from_to_id = from_tos.id").
		Where("rates.date BETWEEN ? AND ?", weekBefore, date).
		Group("from_tos.id").Scan(&tempExRates)

	if len(tempExRates) == 0 {
		return "no currencies exchange available!", err
	}

	for _, rate := range rates {
		var fromTo FromTo
		avgRate := "insufficient data"

		db.Where("id = ?", rate.FromToID).Find(&fromTo)
		db.Where("from_to_id = ?", rate.FromToID).
			Where("date BETWEEN ? AND ?", weekBefore, date).
			Find(&tempRates)

		for _, tempRate := range tempExRates {
			if (tempRate.FromToId == rate.FromToID) && (len(tempRates) >= 7) {
				avgRate = tempRate.WeekAvg
				break
			}
		}

		exRate := ExchangeRate{From: fromTo.From, To: fromTo.To, Rate: rate.Rate, WeekAvg: avgRate}
		exRates = append(exRates, exRate)
	}

	res, err := json.Marshal(exRates)
	return string(res), err
}

func GetRecentRates(from, to string) (string, error) {
	var recentRate RecentRate
	var fromTo FromTo
	var lastRate Rate
	var tempRate ExchangeRate
	var exRates []ExchangeRate
	var rates []Rate
	var tempRates []Rate
	var notFound bool

	notFound = db.Where("from_tos.from = ? AND from_tos.to = ?", from, to).Find(&fromTo).RecordNotFound()

	if notFound {
		return "currency exchange not found!", err
	}

	db.Where("from_to_id = ?", fromTo.ID).Find(&tempRates)

	if len(tempRates) == 0 {
		return "rates not found!", err
	}

	lastRate = tempRates[len(tempRates)-1]

	formatedDate, _ := time.Parse(dateFormat, lastRate.Date)
	weekBefore := formatedDate.AddDate(0, 0, -6).Format(dateFormat)

	notFound = db.Table("from_tos").
		Select("from_tos.id as from_to_id, avg(rates.rate) as week_avg, min(rates.rate) as min, max(rates.rate) as max").
		Joins("left join rates on rates.from_to_id = from_tos.id").
		Where("rates.from_to_id = ? AND rates.date BETWEEN ? AND ?", fromTo.ID, weekBefore, lastRate.Date).
		Group("from_tos.id").Scan(&tempRate).RecordNotFound()

	if notFound {
		return "Rate details not found!", err
	}

	db.Where("from_to_id = ? AND date BETWEEN ? AND ?", fromTo.ID, weekBefore, lastRate.Date).Find(&rates)

	if len(rates) == 0 {
		return "rates not found", err
	}

	for _, rate := range rates {
		exRate := ExchangeRate{Rate: rate.Rate, Date: rate.Date}
		exRates = append(exRates, exRate)
	}

	variance := tempRate.Max - tempRate.Min
	rateDetails := ExchangeRate{From: fromTo.From, To: fromTo.To, WeekAvg: tempRate.WeekAvg, Variance: variance}

	recentRate = RecentRate{RateDetails: rateDetails, ExRates: exRates}

	res, err := json.Marshal(recentRate)
	return string(res), err
}

func CreateRate(rate Rate) string {
	if db.Create(&rate).Error != nil {
		return "create failed! please try again later"
	}

	return "data has been successfully created!"
}
