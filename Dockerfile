FROM golang:latest

WORKDIR /app

# Creates a go module named workspace
RUN go mod init workspace

# Utilizes Gin - a Golang web framework for fast API Dev.
RUN go get github.com/gin-gonic/gin
RUN go get github.com/google/uuid
RUN go get github.com/stretchr/testify
RUN go get github.com/stretchr/testify/assert@v1.10.0

COPY main.go .
COPY point_calculation_logic.go . 
COPY main_test.go . 

