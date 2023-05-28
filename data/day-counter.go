package data

import (
	"encoding/json"
	"os"
	"time"
)

type SchoolData struct {
	SchoolDays   [][][]int `json:"school-days"`
	ThisYearEnds []int     `json:"this-year-ends"`
	SchoolEnds   []int     `json:"school-ends"`
}

func readSchoolData() (SchoolData, error) {
	data, err := os.ReadFile("static/json/ferien.json")
	if err != nil {
		return SchoolData{}, err
	}
	res := SchoolData{}
	err = json.Unmarshal(data, &res)
	return res, err
}

func calcWorkDaysBetween(data [][][]int, begin []int, end []int) int {
	res := 0
	for yIndex, yData := range data {
		for mIndex, mData := range yData {
			for dIndex, dData := range mData {
				if yIndex < begin[0] {
					continue
				}
				if yIndex == begin[0] && mIndex < begin[1] {
					continue
				}
				if yIndex == begin[0] && mIndex == begin[1] && dIndex <= begin[2] {
					continue
				}

				if yIndex > end[0] {
					continue
				}
				if yIndex == end[0] && mIndex > end[1] {
					continue
				}
				if yIndex == end[0] && mIndex == end[1] && dIndex > end[2] {
					continue
				}

				if dData == 1 {
					res += 1
				}
			}
		}
	}
	return res
}

func DayCounter() (map[string]any, error) {
	data, err := readSchoolData()
	if err != nil {
		return map[string]any{}, err
	}

	now := time.Now()
	year := now.Year() - 2023
	month := int(now.Month()) - 1
	day := now.Day() - 1

	schoolAll := calcWorkDaysBetween(data.SchoolDays, []int{year, month, day}, data.SchoolEnds)
	schoolThis := calcWorkDaysBetween(data.SchoolDays, []int{year, month, day}, data.ThisYearEnds)

	counterYear := 365 - now.YearDay()
	if now.Year()%4 == 0 && now.Year()%100 != 0 {
		counterYear = 366 - now.YearDay()
	}

	return map[string]any{
		"day-counter-school-this": schoolThis,
		"day-counter-school-all":  schoolAll,
		"day-counter-year":        counterYear,
	}, nil
}
