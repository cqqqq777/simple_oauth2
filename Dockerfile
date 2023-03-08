FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o oauth2 .

FROM alpine:3.17

COPY /config /

COPY --from=builder /build/oauth2 /

ENTRYPOINT ["./oauth2"]