FROM golang:latest

WORKDIR /app
COPY /app /app

RUN go mod download
RUN go build main.go

EXPOSE 3000

CMD ["./main"]
