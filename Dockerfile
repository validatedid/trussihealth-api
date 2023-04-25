#!/bin/bash
# syntax=docker/dockerfile:1
FROM --platform=linux/amd64 golang:1.20-alpine AS build
# set the working directory inside the container
WORKDIR /app

# copy the source code into the container
COPY src/ ./src/
COPY go.mod .
COPY go.sum .
# How to run docker locally
# Uncomment
# COPY .env .
# and run
# docker run -e APP_ENV=local -p 3011:3011 -t trussicontainer
# OR
# Set APP_ENV=deployment in your .env file and run
# docker run --env-file=.env -p 3011:3011 -t trussicontainer


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


