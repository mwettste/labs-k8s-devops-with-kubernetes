FROM golang:1.17.1 as builder

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app/
COPY go.sum go.mod *.go ./
RUN go mod download && go build -o app /app/todo-app.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/ /app/
ENV PORT=8080

CMD ./app