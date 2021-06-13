# Goal

An event driven system that will connect to the BTC and ETH blockchains,
and emit an event when a large transaction takes place.

I can use Server-Sent Events (SSE) to send the data to any clients that register.

# RabbitMQ
Run it from [Docker](https://registry.hub.docker.com/_/rabbitmq/)

The below command also has the management portal enabled by default (on port 15672).
```
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```


# Interfacing with bitcoind

Bitcoind has the ability to run certain commands on some triggers.
For example, to run a command at startup:

```
bitcoind -startupnotify=<cmd>
bitcoind -startupnotify="/home/kgb33/Code/CryptoWhaleWatchers/bitcoind_CMD/bitcoindCMD startup"
```

Their is also `-blocknotify=<cmd>`, which can inject the hash of the new block too!


Full Example:
```
bitcoind \
-startupnotify="/home/kgb33/Code/CryptoWhaleWatchers/BitcoindCMD/BitcoindCMD startup" \
-blocknotify="/home/kgb33/Code/CryptoWhaleWatchers/BitcoindCMD/BitcoindCMD block --hash %s"
```
