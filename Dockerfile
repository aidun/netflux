FROM golang:1.12.0-stretch as builder

COPY ./go.mod ./go.sum /netflux/
RUN go mod download

COPY . /netflux
WORKDIR /netflux
RUN CGO_ENABLED=0 GOOS=linux go build
#second stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /netflux/netflux .
CMD ["./netflux"]