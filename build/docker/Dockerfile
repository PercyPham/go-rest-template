FROM golang:1.14-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .


FROM alpine
WORKDIR /app
COPY --from=builder /app/.env /app/.env
COPY --from=builder /app/main /app/main
CMD ["/app/main"]
