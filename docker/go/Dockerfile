FROM golang:1.18-alpine3.16 as builder

RUN apk add --update --upgrade build-base

WORKDIR /app

COPY ../../.. .

RUN go mod download && \
    go build -o server -ldflags '-linkmode external -w -extldflags "-static"' -tags musl api/main.go

FROM alpine

RUN apk add --update bash curl

COPY --from=builder /app/api server

EXPOSE 8080 8080

CMD ["./server"]