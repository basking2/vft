FROM golang:1.9.2-alpine as builder
WORKDIR /go/src/github.com/bbriggs/vft
COPY . .
RUN go get -d -v .
RUN apk update
RUN apk --no-cache add git
RUN apk add build-base
WORKDIR /go/src/github.com/bbriggs/vft/cmd/vft/
RUN go get ./...
RUN go build -o vft .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/bbriggs/vft/cmd/vft/vft .
ENTRYPOINT ["/root/vft"]
CMD ["--bind","0.0.0.0:9999","server"]
