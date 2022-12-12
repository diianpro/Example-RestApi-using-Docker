# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 go build -o simple-app

CMD [ "./simple-app" ]