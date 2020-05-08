# feeds-encoder
encoder script for swarm feeds content update.

## use
call the script with comma-separated hex prefixes as its only argument.

## example
if your pinner addresses have the following prefixes:
- `3978`
- `d19c`
- `9c26`

then run:
```sh
go run encode.go "3978,d19c,9c26"
```
which will result in:
```sh
5b224f58673d222c22305a773d222c226e43593d225d
```

this resulting hex string should be posted to the recovery feed as its content so that the swarm client correctly derives the trojan targets.

## posting to a swarm feed
for instructions on how to submit content to a feed, please refer to the official [swarm documentation](https://swarm-guide.readthedocs.io/en/latest/features/feed.html#id5).