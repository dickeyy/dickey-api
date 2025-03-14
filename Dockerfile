FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o dickey-api .

EXPOSE 8000

CMD ["./dickey-api"]