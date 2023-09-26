FROM golang:1.21-alpine

WORKDIR /app

RUN go mod download && go mod verify

RUN go install github.com/cosmtrek/air@latest

COPY ./ /app/

CMD ["air", "-c", ".air.toml"]