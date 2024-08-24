FROM golang:1.23.0-alpine3.19

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /rrqd

EXPOSE 8080

CMD [ "/rrqd" ]