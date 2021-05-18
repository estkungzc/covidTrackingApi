package service

import (
	"encoding/json"
	"time"

	"github.com/estkungzc/covidTrackingApi/model"
	"github.com/estkungzc/covidTrackingApi/utils"
	"github.com/gojektech/heimdall/v6/httpclient"
)

type CovidTrackerService interface {
	GetCovidStat() (*model.CovidClientResponse, error)
	GetCovidSummary(covidStat *model.CovidClientResponse) model.CovidSummaryResponse
}

type CovidTrackerServiceImp struct {
}

func NewCovidTrackerServiceImp() CovidTrackerService {
	return &CovidTrackerServiceImp{}
}

// GetCovidStat from api : http://static.wongnai.com/devinterview/covid-cases.json
func (s *CovidTrackerServiceImp) GetCovidStat() (*model.CovidClientResponse, error) {
	// Create a new HTTP client with a default timeout
	timeout := 1000 * time.Millisecond
	request := utils.CustomHttpClient{
		Client:      httpclient.NewClient(httpclient.WithHTTPTimeout(timeout)),
		Url:         "http://static.wongnai.com/devinterview/covid-cases.json",
		BodyRequest: nil,
	}

	data := model.CovidClientResponse{}

	body, err := utils.MakeRequest(request)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// GetCovidSummary group data from covid status each case
func (s *CovidTrackerServiceImp) GetCovidSummary(covidStat *model.CovidClientResponse) model.CovidSummaryResponse {
	data := covidStat.Data

	result := model.CovidSummaryResponse{
		Province: make(map[string]int),
		AgeGroup: model.AgeGroup{
			Young:  0,
			Middle: 0,
			Old:    0,
			NA:     0,
		},
	}

	// summary response
	for _, v := range data {
		// summary by province
		if v.ProvinceEn != nil {
			result.Province[*v.ProvinceEn]++
		}

		// summary by age group
		switch {
		case v.Age == nil:
			result.AgeGroup.NA++
		case 0 <= *v.Age && *v.Age <= 30:
			result.AgeGroup.Young++
		case 30 < *v.Age && *v.Age <= 60:
			result.AgeGroup.Middle++
		case *v.Age > 60:
			result.AgeGroup.Old++
		}
	}

	return result
}
