package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCovidStatResponse(t *testing.T) {
	covidTrackerService := NewCovidTrackerServiceImp()

	_, err := covidTrackerService.GetCovidStat()

	require.NoError(t, err)
}

func TestCovidSummaryResponse(t *testing.T) {
	// prepare test cases
	//testCases := []struct {
	//	name       string
	//	input      []model.CovidClientResponse
	//	buildStubs func(store *mock_service.MockCovidTrackerService)
	//	expected   model.CovidSummaryResponse
	//}

	t.Error("Not Implemented...")

}
