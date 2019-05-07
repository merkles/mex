package mex

import (
	"github.com/bitly/go-simplejson"
	"huobiapi/client"
	"huobiapi/market"
)

type JSON = simplejson.Json

type ParamsData = client.ParamData
type Market = market.Market
type Listener = market.Listener
type Client = client.Client

// 创建WebSocket版Market客户端
func NewMarket() (*market.Market, error) {
	return market.NewMarket(market.Endpoint)
}

/// 创建WebSocket版Market客户端,指定endPoint
func NewMarketByEndPoint(endPoint string) (*market.Market, error) {
	return market.NewMarket(endPoint)
}

/// 创建RESTFul客户端
func NewClient(accessKeyId, accessKeySecret string) (*client.Client, error) {
	return client.NewClient(client.Endpoint, accessKeyId, accessKeySecret)
}
