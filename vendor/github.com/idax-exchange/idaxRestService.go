package idax

// IDAX is wrapper for Binance API.
//
// Read web documentation for more endpoints descriptions and list of
// mandatory and optional params. Wrapper is not responsible for client-side
// validation and only sends requests further.
//
// For each API-defined enum there's a special type and list of defined
// enum values to be used.
type IdaxRestService interface {
	IdaxHttp(method string, httpUrl string, reqStruct interface{}) ([]byte, Error)
	SendCommonRequest(url string, reqParam interface{}, respStruct interface{}, error *Error)
	// Return server time
	IDAXTime() (IDAXTimeResp, Error)
	// Place an order and return the result of the order
	PlaceOrder(reqParam PlaceOrderReq) (PlaceOrderResp, Error)
	// Withdrawal and return the result of withdrawal
	CancelOrder(reqParam CancelOrderReq) (CancelOrderResp, Error)
	// Return order information
	OrderInfo(reqParam OrderInfoReq) (OrderInfoResp, Error)
	// Return historical order information
	OrderHistory(reqParam OrderHistoryReq) (OrderHistoryResp, Error)
	// Return transaction information
	Trades(reqParam TradesReq) (TradesResp, Error)
	// Return my transaction information
	MyTrades(reqParam MyTradesReq) (MyTradesResp, Error)
	// Return account funds information
	AccountInfo(reqParam AccountInfoReq) (AccountInfoResp, Error)
	// Return price change statistics.
	Ticker(reqParam TickerReq) (TickerResp, Error)
	// DepositHistory lists deposit data.
	Depth(reqParam DepthReq) (DepthResp, Error)
	// Return Kline.
	Kline(reqParam KlineReq) (KlineResp, Error)
	// Return all transaction pairs
	Pairs() (PairsResp, Error)
	PairLimits(reqParam PairLimitsReq) (PairLimitsResp, Error)
	// Returns the signature
	GetSign(reqParam GetSignReq) (GetSignResp, Error)
}

// Return the new IdaxRestService interface
func NewIdaxRestService(url string, apiKey string, secret string) IdaxRestService {
	return &IdaxRestConn{url, apiKey, secret}
}
