# covidTrackingApi

## Description

-   API summary COVID-19 stats using https://covid19.th-stat.com/api/open/cases
-   Structure base on https://github.com/thanawatpetchuen/ginstructure

## Tools

-   [Gin Web Framework](https://github.com/gin-gonic/gin)
-   [Viper](https://github.com/spf13/viper) - Configuration
-   [Heimdall](https://github.com/gojek/heimdall) - HTTP Client
-   Go modules

## API

| resource             | description                      |
| :------------------- | :------------------------------- |
| `GET /covid/summary` | returns summarize COVID-19 stats |

Response of summarize COVID-19 stats

```json
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
}
```

## Environment

Change environments in config/config.yaml

```yaml
env: debug || release || test
service-name: covid-tracking-api
http:
    port: 3000
```

## Project Structure

    .
    ├── app
    │   └── app.go
    ├── config
    │   └── config.yaml
    ├── handler
    │   └── handler.go
    ├── internal
    │   ├── config
    │   │   └── loader.go
    │   └── server
    │       └── server.go
    ├── model
    │   └── config.go
    ├── router
    │   └── httpRouter.go
    ├── service
    │   ├── mock
    │   │   └── covid_tracker.mock.go
    │   ├── covid_tracker.go
    │   └── covid_tracker_test.go
    ├── utils
    │   └── utils.go
    ├── docker-compose.yml
    ├── Dockerfile
    ├── Makefile
    ├── main.go
    └── README.md
