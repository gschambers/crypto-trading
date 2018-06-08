# Crypto Trading
_A demo trading visualisation using Go and React_

This app comprises a backend order book, price ticker and WebSocket interface - written in Go - to power a simple UI visualisation - written in React.

The backend service connects to a firehose of currency exchange prices (via GDAX public price feed) and broadcasts those ticks over a websocket to any subscribed client. Clients receive and buffer ticks in order to render a simple sparkline.

To subscribe to a currency pair, the client sends a message in the following format:

```
client.subscribe("BTC-USD");
```

## Getting started

0. Ensure `yarn` is installed (https://yarnpkg.com)
1. Fetch project: `go get gitlab.com/gschambers/crypto-trading`
2. Install client dependencies: `yarn install`
3. Build client: `yarn build`
4. Build server: `go build`
5. Run: `./crypto-trading` and visit `http://localhost:3000`

## TODO

* Update live visualisation to represent current market price and volume
* Implement visualisations for historical market data
* Simulate market, limit and stop order placing
