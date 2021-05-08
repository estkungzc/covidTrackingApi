FROM golang:1.16.4 AS builder

WORKDIR /app
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -a -o /app/build/covidTrackingApi .

FROM alpine

WORKDIR /app
COPY --from=builder /app/build/covidTrackingApi .
COPY --from=builder /app/config ./config

CMD [ "./covidTrackingApi"]