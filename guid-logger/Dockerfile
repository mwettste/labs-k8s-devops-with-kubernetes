FROM golang:1.17.1 as builder

WORKDIR /app
COPY go.mod go.sum *.go ./

ENV GOOS linux
ENV CGO_ENABLED 0

RUN go mod download && go build -o app /app/guid-logger.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app /app/

CMD ./app