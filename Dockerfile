FROM golang:latest
LABEL authors="dmitry"

RUN cd / && git clone https://github.com/dream-4ild/anon_bot.git && git checkout dev

WORKDIR /anon_bot/src/main/

RUN go mod download

RUN go build -o bot .

CMD ["/anon_bot/src/main/bot"]