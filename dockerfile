# Use official Go image with exact version
FROM golang:1.24.3

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -tags netgo -ldflags '-s -w' -o app

EXPOSE 8080
CMD ["./app"]
