FROM golang:1.16.4-alpine3.13

WORKDIR /app

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 go build -o main .

EXPOSE 4000

CMD ["./main"]
