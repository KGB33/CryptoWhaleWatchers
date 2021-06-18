FROM golang:latest as BUILDER

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .

RUN go build .

FROM debian:latest

RUN apt-get update && apt-get install wget -y

COPY --from=BUILDER /build/BlkProcessor .

CMD ["./BlkProcessor"]
