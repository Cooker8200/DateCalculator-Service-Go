FROM golang:1.21

# TODO: pass aws creds as params

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app .

EXPOSE 3001

CMD ["./app"]
