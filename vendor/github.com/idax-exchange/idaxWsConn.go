package idax

import (
	"fmt"
	"github.com/idax-exchange/websocket"
	"log"
	"time"
)

//
// idax websocket connection
//

// web socket connection structure
type IdaxWsConn struct {
	Conn   websocket.Conn
	Key    string
	Secret string
}

// websocket sends messages
func (iwc *IdaxWsConn) SendMessage(jsonStr string) error {
	_, err := iwc.Conn.Write([]byte(jsonStr))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Send: %s\n", jsonStr)
	return nil
}

// websocket receives messages
func (iwc *IdaxWsConn) ReceiveMessage() chan interface{} {
	// Initialization channel
	rchan := make(chan interface{}, 1)
	go func() {
		for {
			// Get response data
			var msg = make([]byte, 512)
			_, err := iwc.Conn.Read(msg)
			if err != nil {
				log.Fatal(err)
			}
			// Send the received data to the channel
			rchan <- msg
		}
	}()
	// Return to a receiving channel
	return rchan
}

// Send subscriptions `symbol` Ticker.
func (iwc *IdaxWsConn) SendSubXTicker(symbol string) error {
	return iwc.SendMessage(fmt.Sprintf(WS_X_TICKER, symbol))
}

// Send subscriptions `symbol` Depth.
func (iwc *IdaxWsConn) SendSubXDepth(symbol string) error {
	return iwc.SendMessage(fmt.Sprintf(WS_X_DEPTH, symbol))
}

// Send subscriptions `symbol`  Depth Count[N]
func (iwc *IdaxWsConn) SendSubXDepthY(symbol string, number int) error {
	return iwc.SendMessage(fmt.Sprintf(WS_X_DEPTH_Y, symbol, number))
}

// Send subscriptions `symbol` Kline Interval
func (iwc *IdaxWsConn) SendSubXKlineY(symbol string, lineType string) error {
	return iwc.SendMessage(fmt.Sprintf(WS_X_KLINE_Y, symbol, lineType))
}

// Send subscriptions `symbol` Trades
func (iwc *IdaxWsConn) SendSubXTrades(symbol string) error {
	return iwc.SendMessage(fmt.Sprintf(WS_X_TRADES, symbol))
}

// Send a subscription to my trade message
func (iwc *IdaxWsConn) SendSubMyTrade() error {
	if iwc.Secret == "" || iwc.Key == "" {
		return Error{Code: WS_NOT_KS}
	}
	req := WsMyReq{Timestamp: time.Now().UnixNano()}
	return iwc.SendMessage(fmt.Sprintf(WS_MY_TRADE, AddSignToJsonStr(req, iwc.Key, iwc.Secret)))
}

// Send a subscription to my order message
func (iwc *IdaxWsConn) SendSubMyOrder() error {
	req := WsMyReq{Timestamp: time.Now().UnixNano()}
	return iwc.SendMessage(fmt.Sprintf(WS_MY_ORDER, AddSignToJsonStr(req, iwc.Key, iwc.Secret)))
}
