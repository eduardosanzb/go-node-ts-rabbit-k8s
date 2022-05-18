# Unity challenge for Eduardo

## Overview
```

             ┌──────────────────┐
             │                  │
             │                  │
             │                  │
  POST       │                  │
  /to-queue  │       Producer   │
─────────────►                  ├─────┐      ┌──────────────────┐
             │  puts valid      │     │      │                  │
             │ messages into    │     │      │                  │
             │   the queue      │     │      │  rabbitMQ        │
             │                  │     │      │                  │
             └──────────────────┘     │      │                  │
                                      └──────►                  │
                                             │                  │
                                             └──────────┬───────┘
                                                        │
                                                        │
                                                        │
                                             ┌──────────▼───────┐
                                             │                  │
                                             │                  │
                                             │                  │
                                             │                  │
                                             │     consumer     │
                                             │                  │
                                             │                  │
                                             │                  │
                                             └─────────┬────────┘
                                                       │
                                                       │
                                             ┌─────────▼───┐
                       ┌──────────┐          │             │
                       │          │          │    postgres │
                       │ hasura   ├──────────►             │
                       └──────────┘          │             │
                                             │             │
                                             │             │
                                             └─────────────┘
```

## Structure of the project

The project have the next components:
- Postgres
- RabbitMQ
- Hasura (for quick access and querying the db)
- Producer. Built with typescript; which will validate the request and put in the broker
- Consumer. Built with Go; which will insert the messages into the DB

Potentially the technology for producer/consumer should be swap (given that Go is way superior with concurrent request to handle bazillion of requests; yes even better than the *blazingly fast* nodejs [example](https://www.youtube.com/watch?v=h7UEwBaGoVo&t=1s))
but given that the position required TS more than Go, I wanted to showcase my skills with TS.

## How to run the dev environment

I like to put all in the docker compose env.
e.g. eduardo-local-compose.yaml
open some ports and play.

You can stress the system with a small code in ./stress-test
just `go build` or `go run main.go`

### Docker compose

run the docker compose file

### Kubernetes

pre-requisits
You need kind. You can use docker desktop with kind/kind as cluster engine.

Just `sh ./k8s/init.sh`

## Conclusions

bla bla bla resulst about my stress tests
