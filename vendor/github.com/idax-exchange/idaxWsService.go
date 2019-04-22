package idax

import (
	"github.com/idax-exchange/websocket"
	"log"
)

//
//	IDAX websocket interface service
//
type IdaxWsService interface {
	// websocket sends messages
	SendMessage(jsonStr string) error
	// websocket receives messages and returns a message channel
	ReceiveMessage() chan interface{}
	// Send subscriptions `symbol` Ticker.
	SendSubXTicker(symbol string) error
	// Send subscriptions `symbol` Depth.
	SendSubXDepth(symbol string) error
	// Send subscriptions `symbol`  Depth Count[N]
	SendSubXDepthY(symbol string, number int) error
	// Send subscriptions `symbol` Kline Interval
	SendSubXKlineY(symbol string, lineType string) error
	// Send subscriptions `symbol` Trades
	SendSubXTrades(symbol string) error
	// Send a subscription to my trade message
	SendSubMyTrade() error
	// Send a subscription to my order message
	SendSubMyOrder() error
}

// Return the new IdaxWsService interface
func NewIdaxWsService(url string, origin string, apiKey string, secret string) IdaxWsService {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	return &IdaxWsConn{*ws, apiKey, secret}
}
