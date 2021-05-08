# covidTrackingApi

## Description
- API summary COVID-19 stats using https://covid19.th-stat.com/api/open/cases
- Structure base on https://github.com/thanawatpetchuen/covidTrackingApi

## Tools
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper) - Configuration
- [Heimdall](https://github.com/gojek/heimdall) - HTTP Client

## API
| resource | description |
|:----------------------------|:----------------------------------|
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
