FROM golang:1.18-alpine3.16 as builder

RUN apk add --update --upgrade build-base

WORKDIR /app

COPY ../../.. .

RUN go mod download && \
    go build -o main -ldflags '-linkmode external -w -extldflags "-static"' -tags musl api/main.go

FROM alpine

RUN apk add --update bash curl

COPY --from=builder /app/api main

EXPOSE 3000 3001

CMD ["./main"]