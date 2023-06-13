#!/bin/bash
# syntax=docker/dockerfile:1
FROM --platform=linux/amd64 golang:1.20-alpine AS build
# set the working directory inside the container
WORKDIR /app

# copy the source code into the container
COPY src/ ./src/
COPY go.mod .
COPY go.sum .
# Uncomment this line to run the image locally
# COPY .env .

# install the dependencies
RUN go mod download

# build the application
RUN go build -o executable ./src/main.go

# set the environment variables
ENV PORT=3011

# expose the port that the application listens on
EXPOSE $PORT

# run the application
CMD ["./executable"]


