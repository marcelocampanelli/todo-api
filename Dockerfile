FROM golang:1.21-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ./ /app/

RUN go mod tidy && go mod download && go mod verify

CMD ["air", "-c", ".air.toml"]