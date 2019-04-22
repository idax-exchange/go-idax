package idax

import "fmt"

// All data structures

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type IDAXTimeResp struct {
	Timestamp int64 `json:"timestamp"`
}

type PlaceOrderReq struct {
	Key       string  `json:"key"`       // apiKey of the user.
	Pair      string  `json:"pair"`      // IDAX supports trade pairs
	OrderType string  `json:"orderType"` // order type: (limit/market)
	OrderSide string  `json:"orderSide"` // order side:(buy/sell)
	Price     float32 `json:"price"`     // order price
	Amount    float32 `json:"amount"`    // order qty
	Timestamp int64   `json:"timestamp"` // Request timestamp (valid for 3 minutes)
	Sign      string  `json:"sign"`      // signature of request parameters
}

type PlaceOrderResp struct {
	OrderId string `json:"orderId"` // type String  order ID
}

type CancelOrderReq struct {
	OrderId   string `json:"orderId"`   // order ID (multiple orders are separated by a comma ',', Max of 5 orders are allowed per request)
	Timestamp int64  `json:"timestamp"` // Request timestamp (valid for 3 minutes)
	Key       string `json:"key"`       // apiKey of the user
	Sign      string `json:"sign"`      // signature of request parameters
}

type CancelOrderResp struct {
	Accepted string `json:"accepted"` // IDs(Accepted request for cancellation of order)
}

type OrderInfo struct {
	Quantity     string `json:"quantity"`     // order quantity
	AvgPrice     string `json:"avgPrice"`     // average transaction price
	Timestamp    int64  `json:"timestamp"`    // order time
	DealQuantity string `json:"dealQuantity"` // filled quantity
	OrderId      int64  `json:"orderId"`      // order ID
	Price        string `json:"price"`        // order price
	OrderState   int    `json:"orderState"`   // 1 = unfilled,2 = partially filled, 9 = fully filled, 19 = cancelled
	OrderSide    string `json:"orderSide"`    // buy/sell
}

type OrderInfoReq struct {
	Pair      string `json:"pair"`      // IDAX supports trade pairs
	OrderId   string `json:"orderId"`   // if order_id is -1, then return all unfilled orders, otherwise return the order specified
	PageIndex int    `json:"pageIndex"` // current page number
	PageSize  int    `json:"pageSize"`  // number of orders returned per page
	Timestamp int64  `json:"timestamp"` // request timestamp (valid for 3 minutes)
	Key       string `json:"key"`       // apiKey of the user
	Sign      string `json:"sign"`      // signature of request parameters
}

type OrderInfoResp struct {
	Orders []OrderInfo `json:"orders"`
	Total  int         `json:"total"`
}

type OrderListReq struct {
	Pair      string `json:"pair"`      // IDAX supports trade pairs
	OrderId   string `json:"orderId"`   // if order_id is -1, then return all unfilled orders, otherwise return the order specified
	Timestamp int64  `json:"timestamp"` // request timestamp (valid for 3 minutes)
	Key       string `json:"key"`       // apiKey of the user
	Sign      string `json:"sign"`      // signature of request parameters
}

type OrderListResp struct {
	Orders []OrderInfo `json:"orders"`
}

type OrderHistoryReq struct {
	Key         string `json:"key"`         // apiKey of the user
	Pair        string `json:"pair"`        // IDAX supports trade pairs
	OrderState  int    `json:"orderState"`  // query status: 0 for all orders,query status: 0 for unfilled orders, 1 for filled orders (only the data of the last two days are returned)
	CurrentPage int    `json:"currentPage"` // current page number
	PageLength  int    `json:"pageLength"`  // number of orders returned per page, maximum 100
	Timestamp   int64  `json:"timestamp"`   // request timestamp (valid for 3 minutes)
	Sign        string `json:"sign"`        // signature of request parameters
}

type OrderHistoryResp struct {
	CurrentPage int         `json:"currentPage"` // current page number
	Orders      []OrderInfo `json:"orders"`
	PageLength  int         `json:"pageLength"` // number of orders per page
	Total       int         `json:"total"`      // The total number of records
}

type Trades struct {
	Timestamp int64  `json:"timestamp"` //trade time
	Price     string `json:"price"`     //deal price
	Quantity  string `json:"quantity"`  //qty in base coin
	Id        string `json:"id"`        //trade id
	Maker     string `json:"maker"`     //deal direction Buy/Sell
}

type TradesReq struct {
	Pair string `json:"pair"` // IDAX supports trade pairs.
}

type TradesResp struct {
	Trades []Trades `json:"trades"`
}

type TradesHistoryReq struct {
	Key       string `json:"key"`       //apiKey of the user
	Pair      string `json:"pair"`      //IDAX supports trade pairs
	Since     int64  `json:"since"`     //Get the latest 600 pieces of data from a given ID(Since fetches the returned trade ID)
	Timestamp int64  `json:"timestamp"` //Request timestamp (valid for 3 minutes)
	Sign      string `json:"sign"`      //signature of request parameters
}

type TradesHistoryResp struct {
	Trades []Trades `json:"trades"`
}

type MyTradesReq struct {
	Key         string `json:"key"`         // apiKey of the user
	Pair        string `json:"pair"`        // IDAX supports trade pairs
	OrderSide   string `json:"orderSide"`   // buy，sell
	CurrentPage int    `json:"currentPage"` // current page number
	PageLength  int    `json:"pageLength"`  // number of trade returned per page, maximum 1000
	StartDate   int64  `json:"startDate"`   // start date and timestamp (Millisecond)
	EndDate     int64  `json:"endDate"`     // end date and timestamp (Millisecond)
	Timestamp   int64  `json:"timestamp"`   // request timestamp (valid for 3 minutes)
	Sign        string `json:"sign"`        // signature of request parameters
}

type MyTradesResp struct {
	Trades []Trades `json:"trades"`
	Total  int      `json:"total"`
}

type AccountInfoReq struct {
	Key       string `json:"key"`       // apiKey of the user.
	Timestamp int64  `json:"timestamp"` // Request timestamp (valid for 3 minutes)
	Sign      string `json:"sign"`      // signature of request parameters
}

type AccountInfoResp struct {
	Total   map[string]string `json:"total"`   // total fund
	Free    map[string]string `json:"free"`    // available fund
	Freezed map[string]string `json:"freezed"` // frozen fund
}

type Ticker struct {
	Pair string `json:"pair"` //pair
	Open string `json:"open"` //open price
	High string `json:"high"` //high price
	Low  string `json:"low"`  //low price
	Last string `json:"last"` //last price
	Vol  string `json:"vol"`  //volume(in the last 24hours sliding window)
}

type TickerReq struct {
	Pair string `json:"pair"` // IDAX supports trade pairs.
}

type TickerResp struct {
	Ticker    []Ticker `json:"ticker"`
	Timestamp int64    `json:"timestamp"`
}

type DepthReq struct {
	Pair  string `json:"pair"`  // IDAX supports trade pairs. Returns a specific trade against the market when specifying a pair,Returns all trading to market prices without specifying pair;
	Size  int    `json:"size"`  // how many price level should be response. must be between 1 - 200
	Merge int    `json:"merge"` // price decimal price should be merged.
}

type DepthResp struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

type KlineReq struct {
	Pair   string `json:"pair"`   // IDAX supports trade pairs.
	Period string `json:"period"` // 1min,5min,15min,30min,1hour,2hour,4hour,6hour,12hour,1day,1week
	Size   int    `json:"size"`   // specify data size to be acquired
	Since  int64  `json:"since"`  // timestamp(eg:1417536000000). Data before returning timestamp
}

type KlineResp struct {
	Kline [][]interface{} `json:"kline"`
}

type PairsResp struct {
	Pairs []string `json:"pairs"`
}

type PairRuleVo struct {
	PairName          string  `json:"pairName"`          // pair
	MaxAmount         string  `json:"maxAmount"`         // max amount
	MinAmount         string  `json:"minAmount"`         // min amount
	PriceDecimalPlace float32 `json:"priceDecimalPlace"` // price decimal
	QtyDecimalPlace   float32 `json:"qtyDecimalPlace"`   // quantity decimal
}

type PairLimitsReq struct {
	Pair string `json:"pair"` // IDAX supports trade pairs.
}

type PairLimitsResp struct {
	PairRuleVo []PairRuleVo `json:"pairRuleVo"`
}

type GetSignReq struct {
	NeedSignature string `json:"needSignature"` //	The string to be signed must be in JSON format
}

type GetSignResp struct {
	Sign string `json:"sign"`
}

type WsMyReq struct {
	Key       string `json:"key"`       // apiKey of the user.
	Timestamp int64  `json:"timestamp"` // Request timestamp (valid for 3 minutes)
	Sign      string `json:"sign"`      // signature of request parameters
}
