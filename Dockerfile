FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .


RUN go build -o main .




FROM alpine:latest

COPY --from=build /app/main .

EXPOSE 8080

cmd ["main"]
