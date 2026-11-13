FROM golang:1.26-alpine3.23 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o main cmd/server/main.go

# --- RUNTIME STAGE --- #

FROM alpine:3.23

WORKDIR /app

COPY --from=builder app/main .

EXPOSE 8080

CMD ["./main"]