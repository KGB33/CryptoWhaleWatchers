FROM golang:latest as BUILDER

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .

RUN go build -o BTCd-CMD .

FROM debian:latest

RUN apt-get update && apt-get install wget -y

COPY --from=BUILDER /build/BTCd-CMD .

# Install Bitcoin Core
ENV BTC_CORE_VER=0.21.1
RUN wget https://bitcoincore.org/bin/bitcoin-core-$BTC_CORE_VER/bitcoin-$BTC_CORE_VER-x86_64-linux-gnu.tar.gz
RUN tar xzf bitcoin-${BTC_CORE_VER}-x86_64-linux-gnu.tar.gz
RUN install -m 0755 -o root -g root -t /usr/local/bin bitcoin-$BTC_CORE_VER/bin/*

CMD ["bitcoind", "-datadir=/bitcoin","-startupnotify='/BTCd-CMD startup'", "-blocknotify='/BTCd-CMD block --hash %s'"]
