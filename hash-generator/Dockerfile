FROM golang:1.17.1 as builder

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app/
COPY . .
RUN go build -o app /app/hash-generator.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/ /app/
CMD ./app