FROM golang:1.23.5-alpine as builder

WORKDIR /app

COPY . .

RUN go get github.com/swaggo/swag/gen@v1.16.4
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4
RUN swag init -g cmd/server/main.go
RUN go mod tidy
RUN go mod download
RUN go build -o main.bin cmd/server/main.go

FROM alpine as release

WORKDIR /app

RUN apk add --no-cache bash
COPY --from=builder /app/main.bin /app/main.bin
COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

EXPOSE 8080

ENTRYPOINT ["/app/wait-for-it.sh", "postgres", "5432", "--", "./main.bin"]