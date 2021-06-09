## Intro

A pool server for the Chia pooling protocol, written in Golang.
Based on https://github.com/Chia-Network/pool-reference

### Usage

1. Start Redis

`docker run -p 6379:6379 --name local-redis -d redis`

2. Start server

`go run *.go`