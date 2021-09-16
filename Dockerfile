FROM golang:1.16-alpine3.13 as builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/karamaru-alpha/test

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o app

FROM alpine:3.13
WORKDIR /root/
COPY --from=builder /go/src/github.com/karamaru-alpha/test .

EXPOSE 8080
ENTRYPOINT ["./app"]
