FROM golang:latest
LABEL authors="dmitry"

WORKDIR /app

COPY ./src/main/go.mod ./src/main/go.sum ./src/main/

RUN cd src/main/ && go mod download

COPY . .

WORKDIR /app/src/main

RUN go build -o bot .

CMD ["/app/src/main/bot"]