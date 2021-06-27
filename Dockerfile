FROM golang:1.13

COPY ./main.go /app/main.go

WORKDIR /app

RUN go build -o main .

EXPOSE 9000

CMD ["./main"]