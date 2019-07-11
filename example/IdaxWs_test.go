package idax_test

import (
	"fmt"
	"github.com/idax-exchange"
	"testing"
)

func TestWebsocket(test *testing.T) {
	// Initialize Websocket client service
	var wsConn = idax.NewIdaxWsService("wss://openws.idax.pro/ws", "https://openws.idax.pro/ws", "your apiKey", "your secret")

	// Initialize Websocket receives message channels
	rchan := wsConn.ReceiveMessage()

	symbol := "ETH_BTC"
	// Send subscriptions `symbol` Ticker. (idaxSubXTicker)
	wsConn.SendSubXTicker(symbol)

	// Send subscriptions `symbol` Depth. (SendSubXDepth)
	wsConn.SendSubXDepth(symbol)

	// Send subscriptions `symbol` Trades. (idaxSubXTrades)
	wsConn.SendSubXTrades(symbol)

	// Send subscriptions `symbol` Depth Count[20]. (idaxSubXDepthY)
	wsConn.SendSubXDepthY(symbol, 20)

	// Send subscriptions `symbol` Kline Interval[idax.Minute]. (SendSubXKlineY)
	wsConn.SendSubXKlineY(symbol, idax.Minute)

	// Send a subscription to my order message. (SendSubMyOrder)
	wsConn.SendSubMyOrder()

	// Send a subscription to my trade message. (SendSubMyTrade)
	wsConn.SendSubMyTrade()

	// 主线程阻塞
	for {
		fmt.Printf("Receive: %s\n", <-rchan)
	}
}
