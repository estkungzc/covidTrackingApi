package service

import (
	"encoding/json"
	"log"
)

type CovidTrackerService struct {
}

func NewCovidTrackerService() *CovidTrackerService {
	return &CovidTrackerService{}
}

func (s *CovidTrackerService) GetCovidSummary() map[string]interface{} {
	covidSummaryData := []byte(`
	{
		"Province": {
			"Samut Sakhon": 3613,
			"Bangkok": 2774
		},
		"AgeGroup": {
			"0-30": 300,
			"31-60": 150,
			"61+": 250,
			"N/A": 4
		}
	}`)

	var data map[string]interface{}
	err := json.Unmarshal([]byte(covidSummaryData), &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
