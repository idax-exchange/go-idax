package idax

import (
	"encoding/json"
	"fmt"
)

type IdaxRestConn struct {
	Url    string
	Key    string
	Secret string
}

// SendCommonRequest
func (irc *IdaxRestConn) SendCommonRequest(url string, reqParam interface{}, respStruct interface{}, err *Error) {
	respByte, error := irc.IdaxHttp("POST", irc.Url+url, reqParam)
	json.Unmarshal(respByte, &respStruct)
	fmt.Print(error)
}

// Get Server timestamp
func (irc *IdaxRestConn) IDAXTime() (IDAXTimeResp, Error) {
	respByte, err := irc.IdaxHttp("GET", irc.Url+REST_TIME, nil)
	var idaxTimeResp IDAXTimeResp
	json.Unmarshal(respByte, &idaxTimeResp)
	return idaxTimeResp, err
}

// Create new order
func (irc *IdaxRestConn) PlaceOrder(reqParam PlaceOrderReq) (PlaceOrderResp, Error) {
	respByte, err := irc.IdaxHttp("POST", irc.Url+REST_PLACE_ORDER, reqParam)
	var placeOrderResp PlaceOrderResp
	json.Unmarshal(respByte, &placeOrderResp)
	return placeOrderResp, err
}

// Cancel orders
func (irc *IdaxRestConn) CancelOrder(reqParam CancelOrderReq) (CancelOrderResp, Error) {
	respByte, err := irc.IdaxHttp("POST", irc.Url+REST_CANCEL_ORDER, reqParam)
	var cancelOrderResp CancelOrderResp
	json.Unmarshal(respByte, &cancelOrderResp)
	return cancelOrderResp, err
}

// Get Order Info
func (irc *IdaxRestConn) OrderInfo(reqParam OrderInfoReq) (OrderInfoResp, Error) {
	respByte, err := irc.IdaxHttp("POST", irc.Url+REST_ORDER_INFO, reqParam)
	var orderInfoResp OrderInfoResp
	json.Unmarshal(respByte, &orderInfoResp)
	return orderInfoResp, err
}

// Get historical order information and return information only for the last two days
func (irc *IdaxRestConn) OrderHistory(reqParam OrderHistoryReq) (OrderHistoryResp, Error) {
	respByte, err := irc.IdaxHttp("POST", irc.Url+REST_ORDER_HISTORY, reqParam)
	var orderHistoryResp OrderHistoryResp
	json.Unmarshal(respByte, &orderHistoryResp)
	return orderHistoryResp, err
}

// Get Recently Trades
func (irc *IdaxRestConn) Trades(reqParam TradesReq) (TradesResp, Error) {
	respByte, err := irc.IdaxHttp("GET", irc.Url+REST_TRADES, reqParam)
	var tradesResp TradesResp
	json.Unmarshal(respByte, &tradesResp)
	return tradesResp, err
}

// Get my historical trading information
func (irc *IdaxRestConn) MyTrades(reqParam MyTradesReq) (MyTradesResp, Error) {
	respByte, err := irc.IdaxHttp("POST", irc.Url+REST_MY_TRADES, reqParam)
	var myTradesResp MyTradesResp
	json.Unmarshal(respByte, &myTradesResp)
	return myTradesResp, err
}

// Get account info
func (irc *IdaxRestConn) AccountInfo(reqParam AccountInfoReq) (AccountInfoResp, Error) {
	respByte, err := irc.IdaxHttp("POST", irc.Url+REST_USERINFO, reqParam)
	var accountInfoResp AccountInfoResp
	json.Unmarshal(respByte, &accountInfoResp)
	return accountInfoResp, err
}

// Get the price of specific ticker
func (irc *IdaxRestConn) Ticker(reqParam TickerReq) (TickerResp, Error) {
	respByte, err := irc.IdaxHttp("GET", irc.Url+REST_TICKER, reqParam)
	var tickerResp TickerResp
	json.Unmarshal(respByte, &tickerResp)
	return tickerResp, err
}

// Get the market depth for specific market.
func (irc *IdaxRestConn) Depth(reqParam DepthReq) (DepthResp, Error) {
	respByte, err := irc.IdaxHttp("GET", irc.Url+REST_DEPTH, reqParam)
	var depthResp DepthResp
	json.Unmarshal(respByte, &depthResp)
	return depthResp, err
}

// Get kline data
func (irc *IdaxRestConn) Kline(reqParam KlineReq) (KlineResp, Error) {
	respByte, err := irc.IdaxHttp("GET", irc.Url+REST_KLINE, reqParam)
	var klineResp KlineResp
	json.Unmarshal(respByte, &klineResp)
	return klineResp, err
}

// All trading pairs supported by exchanges
func (irc *IdaxRestConn) Pairs() (PairsResp, Error) {
	respByte, err := irc.IdaxHttp("GET", irc.Url+REST_PAIRS, nil)
	var pairsResp PairsResp
	json.Unmarshal(respByte, &pairsResp)
	return pairsResp, err
}

// Gets the maximum, minimum, price, and quantity of the supported transaction pairs
func (irc *IdaxRestConn) PairLimits(reqParam PairLimitsReq) (PairLimitsResp, Error) {
	respByte, err := irc.IdaxHttp("GET", irc.Url+REST_PAIR_LIMITS, reqParam)
	var pairLimitsResp PairLimitsResp
	json.Unmarshal(respByte, &pairLimitsResp)
	return pairLimitsResp, err
}

// Get sign
func (irc *IdaxRestConn) GetSign(reqParam GetSignReq) (GetSignResp, Error) {
	respByte, err := irc.IdaxHttp("GET", irc.Url+REST_GET_SIGN, reqParam)
	var getSignResp GetSignResp
	json.Unmarshal(respByte, &getSignResp)
	return getSignResp, err
}
