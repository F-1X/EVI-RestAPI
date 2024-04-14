FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM golang:latest

WORKDIR /app

COPY --from=builder /app/key-firebase.json .
COPY --from=builder /app/.env .
COPY --from=builder /app/main .

EXPOSE 8888


CMD ["./main"]