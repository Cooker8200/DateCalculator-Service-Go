FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o app .

EXPOSE 3001

CMD ["./app"]
