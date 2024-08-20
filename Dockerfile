FROM golang:1.22 as builder

LABEL authors="kmkm2"

WORKDIR /cmd/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/main cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main /app/main

WORKDIR /app

EXPOSE 8080

CMD ["/app/main"]
