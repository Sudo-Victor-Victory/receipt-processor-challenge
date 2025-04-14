FROM golang:latest

WORKDIR /app

# Creates a go module named workspace
RUN go mod init workspace

# Utilizes Gin - a Golang web framework for fast API Dev.
RUN go get github.com/gin-gonic/gin
RUN go get github.com/google/uuid

COPY main.go .

