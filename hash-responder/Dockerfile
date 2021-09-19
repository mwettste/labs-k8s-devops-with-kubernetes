FROM golang:1.17.1 as builder

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app/
COPY *.go go.mod go.sum ./
RUN go build -o app /app/hash-responder.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/app /app/

ENV PORT=5050

CMD ./app