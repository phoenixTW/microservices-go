FROM golang:1.21.1-bullseye AS builder
WORKDIR /srv/go-app
COPY . .
RUN go build -o microservice


FROM debian:buster
WORKDIR /srv/go-app
COPY --from=builder /srv/go-app/config.json .
COPY --from=builder /srv/go-app/microservice .

CMD ["./microservice"]