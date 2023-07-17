FROM golang:1.20-alpine

RUN apk add build-base
RUN apk add git

WORKDIR /app

COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on

RUN go mod download

COPY . .

WORKDIR /app/cmd/server

RUN go build -o /app/shop main.go

EXPOSE 8000

WORKDIR /app

CMD ["/app/shop"]