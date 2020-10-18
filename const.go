package getui

import (
	"time"
)

const (
	BaseUrl = "https://restapi.getui.com/v2/%s/%s"
)

const (
	// defaultRequestTimeout 默认请求超时时间
	defaultRequestTimeout = 3 * time.Second
)

type StrategyNum int

const (
	_               = iota
	SN1 StrategyNum = iota //表示该消息在用户在线时推送个推通道，用户离线时推送厂商通道;
	SN2                    //表示该消息只通过厂商通道策略下发，不考虑用户是否在线;
	SN3                    //表示该消息只通过个推通道下发，不考虑用户是否在线；
	SN4                    //表示该消息优先从厂商通道下发，若消息内容在厂商通道代发失败后会从个推通道下发。
)

