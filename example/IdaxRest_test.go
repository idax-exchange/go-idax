package idax_test

import (
	"fmt"
	"github.com/idax-exchange"
	"testing"
	"time"
)

/** Initialize REST client service **/
var idaxRestService = idax.NewIdaxRestService(
	"https://openapi.idax.pro",
	"93d6db906e814ab3b0ad5c77aa69ebc2bed4390b1f87444bb040ab775d347858",
	"13896d803ff644d2a0033580b8f86b5abf10d2c9f2d9454da0541e586b6d77d4")

/** Test request server time **/
func TestIDAXTime(test *testing.T) {
	idaxTime, err := idaxRestService.IDAXTime()
	fmt.Println("Server Time Response：", idaxTime, err)
}

/** Test Request Place Order **/
func TestPlaceOrder(test *testing.T) {
	reqParam := idax.PlaceOrderReq{Pair: "ETH_BTC", Amount: 1.05, OrderSide: "buy", OrderType: "limit", Price: 0.034775, Timestamp: time.Now().UnixNano()}
	placeOrderResp, err := idaxRestService.PlaceOrder(reqParam)
	fmt.Println(" Place Order response：", placeOrderResp, err)
}

/** Test request Cancel Order **/
func TestCancelOrder(test *testing.T) {
	reqParam := idax.CancelOrderReq{OrderId: "1667991", Timestamp: time.Now().UnixNano()}
	cancelOrderResp, err := idaxRestService.CancelOrder(reqParam)
	fmt.Println("Cancel Order response：", cancelOrderResp, err)
}

/** Test Request Order Information **/
func TestOrderInfo(test *testing.T) {
	reqParam := idax.OrderInfoReq{Pair: "ETH_BTC", OrderId: "-1", PageIndex: 1, PageSize: 10, Timestamp: time.Now().UnixNano()}
	orderInfoResp, err := idaxRestService.OrderInfo(reqParam)
	fmt.Println("Order Information Response：", orderInfoResp, err)
}

/** Test Request History Order **/
func TestOrderHis(test *testing.T) {
	reqParam := idax.OrderHistoryReq{Pair: "ETH_BTC", OrderState: 0, CurrentPage: 1, PageLength: 10, Timestamp: time.Now().UnixNano()}
	orderHistoryResp, err := idaxRestService.OrderHistory(reqParam)
	fmt.Println("Order set response：", orderHistoryResp, err)
}

/** Test requests for the latest 60 transactions **/
func TestTrades(test *testing.T) {
	reqParam := idax.TradesReq{Pair: "ETH_BTC"}
	tradesResp, err := idaxRestService.Trades(reqParam)
	fmt.Println("Test requests for the latest 60 transactions：", tradesResp, err)
}

/** Request for my historical transaction information **/
func TestMyTrade(test *testing.T) {
	reqParam := idax.MyTradesReq{CurrentPage: 1, PageLength: 10, Timestamp: time.Now().UnixNano()}
	myTradesResp, err := idaxRestService.MyTrades(reqParam)
	fmt.Println("My Historical Transaction Information Response：", myTradesResp, err)
}

/** Test Request Account Information **/
func TestAccountInfo(test *testing.T) {
	reqParam := idax.AccountInfoReq{"", time.Now().UnixNano(), ""}
	placeOrderResp, err := idaxRestService.AccountInfo(reqParam)
	fmt.Println("Account Information Response：", placeOrderResp, err)
}

/** Test request to get the price of the transaction pair。 **/
func TestTicker(test *testing.T) {
	tickerResp, err := idaxRestService.Ticker(idax.TickerReq{})
	fmt.Println("Price Response to Transaction Pairs：", tickerResp, err)
}

/** Test Request Acquisition Market Depth **/
func TestDepth(test *testing.T) {
	reqParam := idax.DepthReq{Pair: "ETH_BTC"}
	depthResp, err := idaxRestService.Depth(reqParam)
	fmt.Println("Market Deep Response：", depthResp, err)
}

/** Test Request K-Line Data **/
func TestKline(test *testing.T) {
	reqParam := idax.KlineReq{Pair: "ETH_BTC", Period: idax.Minute}
	klineResp, err := idaxRestService.Kline(reqParam)
	fmt.Println("K-line data response：", klineResp, err)
}

/** Test all trading pairs that request exchange support **/
func TestPairs(test *testing.T) {
	pairsResp, err := idaxRestService.Pairs()
	fmt.Println("All transactions responded：", pairsResp, err)
}

/** Maximum, minimum, price, and number of transaction pairs supported by test requests **/
func TestPairLimits(test *testing.T) {
	reqParam := idax.PairLimitsReq{Pair: "ETH_BTC"}
	pairLimitsResp, err := idaxRestService.PairLimits(reqParam)
	fmt.Println("Maximum, Minimum, Price, and Quantity Responses of Acquiring Supported Transaction Pairs：", pairLimitsResp, err)
}

/** Test Request Signature **/
func TestGetSign(test *testing.T) {
	reqParam := idax.GetSignReq{NeedSignature: `{"key":"93d6db906e814ab3b0ad5c77aa69ebc2bed4390b1f87444bb040ab775d347858","sign":"958da9ec4c890841bc46adc11dc595f6b4666d80c97f36a489ac71b6dc8ccc5e","timestamp":1552531718000}`}
	getSignResp, err := idaxRestService.GetSign(reqParam)
	fmt.Println("Signature response：", getSignResp, err)
}
