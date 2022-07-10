FROM golang:1.17.3-alpine3.13 AS builder
WORKDIR /app
RUN apk add --no-cache git build-base tzdata

COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY . .

RUN go build -o cli cmd/cli/main.go

FROM alpine:3.15 as runner
WORKDIR /app

COPY --from=builder /app/cli .

CMD /app/cli
