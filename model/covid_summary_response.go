package model

type AgeGroup struct {
	Young  int `json:"0-30"`
	Middle int `json:"31-60"`
	Old    int `json:"61+"`
	NA     int `json:"N/A"`
}

type CovidSummaryResponse struct {
	Province map[string]int `json:"Province"`
	AgeGroup AgeGroup       `json:"AgeGroup"`
}
