FROM golang:1.17.1 as builder

WORKDIR /app/
COPY . .
RUN go build -o app /app/hash-generator.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/ /app/
CMD ./app