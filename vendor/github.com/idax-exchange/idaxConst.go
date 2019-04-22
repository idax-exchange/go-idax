package idax

// Constant structure

// REST Path Constant
const (
	REST_TIME          = "/api/v2/time"
	REST_PLACE_ORDER   = "/api/v2/placeOrder"
	REST_CANCEL_ORDER  = "/api/v2/cancelOrder"
	REST_ORDER_INFO    = "/api/v2/orderInfo"
	REST_ORDER_HISTORY = "/api/v2/orderHistory"
	REST_TRADES        = "/api/v2/trades"
	REST_MY_TRADES     = "/api/v2/myTrades"
	REST_USERINFO      = "/api/v2/userinfo"
	REST_TICKER        = "/api/v2/ticker"
	REST_DEPTH         = "/api/v2/depth"
	REST_KLINE         = "/api/v2/kline"
	REST_PAIRS         = "/api/v2/pairs"
	REST_PAIR_LIMITS   = "/api/v2/pairLimits"
	REST_GET_SIGN      = "/api/v2/GetSign"
)

// Websocket channel constant
// Dynamic product placeholder% V
const (
	WS_X_TICKER  = `{"event":"addChannel","channel":"idax_sub_%v_ticker"}`
	WS_X_DEPTH   = `{"event":"addChannel","channel":"idax_sub_%v_depth"}`
	WS_X_DEPTH_Y = `{"event":"addChannel","channel":"idax_sub_%v_depth_%v"}`
	WS_X_TRADES  = `{"event":"addChannel","channel":"idax_sub_%v_trades"}`
	WS_X_KLINE_Y = `{"event":"addChannel","channel":"idax_sub_%v_kline_%v"}`
	WS_MY_TRADE  = `{"event":"addChannel","channel":"idax_sub_mytrade","parameters":%v}`
	WS_MY_ORDER  = `{"event":"addChannel","channel":"idax_sub_myorder","parameters":%v}`
)
