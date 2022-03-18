FROM golang:1.9.2-alpine as builder
WORKDIR /go/src/github.com/madurosecurity/vft
COPY . .
RUN go get -d -v .
RUN apk update
RUN apk --no-cache add git

# Build server
RUN apk add build-base
WORKDIR /go/src/github.com/madurosecurity/vft/cmd/vft-server/
RUN go get ./...
RUN go build -o vft-server .

# Build client
WORKDIR /go/src/github.com/madurosecurity/vft/cmd/vft-client/
RUN go get ./...
RUN go build -o vft-client .

# Install binaries
FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/madurosecurity/vft/cmd/vft-server .
COPY --from=builder /go/src/github.com/madurosecurity/vft/cmd/vft-client .
ENTRYPOINT ["/root/vft-server"]
CMD ["--bind","0.0.0.0:9999"]
