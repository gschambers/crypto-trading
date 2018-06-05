# Crypto Trading
_A demo trading visualisation using Go and React_

This app comprises a backend price stream service and WebSocket interface - written in Go - to power a simple UI visualisation - written in React.

The backend service connects to a firehose of currency exchange prices (synthesised here in a tight loop, to better simulate load) and broadcasts those ticks over a websocket to any subscribed client. Clients receive and buffer ticks in order to render a simple sparkline.

To subscribe to a currency pair, the client sends a message in the following format:

```
client.subscribe({ from: "BTC", to: "USD" });
```

NB. the synthetic price feed is currently hardcoded to generate prices _only_ for BTC/USD, in the range 100-105.

## Getting started

0. Ensure `yarn` is installed (https://yarnpkg.com)
1. Fetch project: `go get gitlab.com/gschambers/crypto-trading`
2. Install client dependencies: `yarn install`
3. Build client: `yarn build`
4. Build server: `go build`
5. Run: `./crypto-trading` and visit `http://localhost:3000`

## TODO

* Integrate with GDAX WebSocket feed for live market data
* Implement visualisations for historical and live market data
* Simulate market, limit and stop order placing
