FROM golang:1.25-alpine

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy

