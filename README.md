
#  CryptoWhaleWatchers

An event driven notification system for the Bitcoin Blockchain.


## Deployment

Unfortunately, to run the project you will set up a Bitcoin "Full Node".
This setup involves downloading the entire blockchain, about 370 GB.
You will also need Docker and Docker Compose.

  - [Bitcoind](https://bitcoin.org/en/full-node#what-is-a-full-node)
  - [Docker](https://docs.docker.com/get-docker/)
  - [Docker Compose](https://github.com/docker/compose)

To run the code:

```bash
git clone git@github.com:KGB33/CryptoWhaleWatchers.git
cd CryptoWhaleWatchers
docker-compose up [--build]
```

## Resources

[Working with RabbitMQ in Golang by examples](https://dev.to/koddr/working-with-rabbitmq-in-golang-by-examples-2dcn#toc)
