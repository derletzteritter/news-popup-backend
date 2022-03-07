FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["/news-popup-backend"]