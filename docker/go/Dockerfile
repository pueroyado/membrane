FROM golang:latest

WORKDIR /app

ADD /app /app

RUN go mod download

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]
