# IDAX API

To read full documentation, specs and find out which request params are required/optional, please visit the official
[documentation](https://github.com/idax-exchange/idax-official-api-docs) page.

## Getting started

```go
// Initialize REST client service
var idaxRestService = idax.NewIdaxRestService(
	"https://qa-openapi.idax.mn",
	"key", 
	"secret"
)

// Initialize Websocket client service
var idaxWsConn = idax.NewIdaxWsService(
	"wss://openws.idax.pro/ws",
	"https://openws.idax.pro/ws",
	"key",
	"secret"
)
// Initialize Websocket receives message channels
rchan := wsConn.ReceiveMessage()
```

## Examples

Following provides list of main usages of library. See `example` package for testing application with more examples.

Each call has its own *Request* structure with data that can be provided. The library is not responsible for validating
the input and if non-zero value is used, the param is sent to the API server.

In case of an standard error, instance of `idaxStruct.Error` is returned with additional info.

The request structure and response structure of each interface are in place `idaxStruct`

### PlaceOrder

```go
// Request structure
reqParam := idax.PlaceOrderReq{
	Pair: "ETH_BTC", 
	Amount: 1.05, 
	OrderSide: "buy", 
	OrderType: "limit", 
	Price: 0.034775, 
	Timestamp: time.Now().UnixNano()
}

// Send request
placeOrderResp, err := idaxRestService.PlaceOrder(reqParam)
if err != nil {
    panic(err)
}
fmt.Println("PlaceOrder Resp：", placeOrderResp, err)
```

### CancelOrder

```go
// Request structure
reqParam := idax.CancelOrderReq{
	OrderId: "1667991",
	Timestamp: time.Now().UnixNano()
}

// Send request
cancelOrderResp, err := idaxRestService.CancelOrder(reqParam)
if err != nil {
    panic(err)
}
fmt.Println("CancelOrder Resp：", cancelOrderResp, err)
```

### Kline

```go
// Request structure
reqParam := idax.KlineReq{
	Pair: "ETH_BTC", 
	Period: idax.Minute
}
// Send request
klineResp, err := idaxRestService.Kline(reqParam)
if err != nil {
       panic(err)
   }
fmt.Println("Kline Resp：", klineResp, err)
```
    
### Trade Websocket

```go
// Initialize pairs
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


for {
	// Input and receive messages
	fmt.Printf("Websocket Receive: %s\n", <-rchan)
}
```

## Known issues

* Websocket error handling is not perfect and occasionally attempts to read from closed connection.