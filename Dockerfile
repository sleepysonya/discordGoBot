FROM golang:1.20 AS build_chonky

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY ./birthday ./birthday
COPY ./reminder ./reminder
COPY ./util ./util
COPY ./random ./random
COPY ./test ./test
COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-golang

CMD ["/docker-golang"]