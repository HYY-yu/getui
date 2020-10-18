package getui

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/spf13/cast"
)

//http://docs.getui.com/getui/server/rest_v2/push/
// 使用时对照个推文档
/*
	加载好配置

	resp,err := NewGetui(GeTuiConing{}).ToSingleCid(
			&Req{
				RequestId: "xxxxx",
				Audience: Audience{
					Cid: []string{"cid"},
				},
				PushMessage: PushMessage{
					Transmission: "透传消息",
				},
			}
		)
	if err != nil{
		// 只要个推没返回 HTTPCode=200 且 resp.Code=0 ,err就有值
	}

*/

var (
	ErrorRequestIdLen = errors.New(" must have request_id, and the length of request_id in [10,32] ")
)

type Getui struct {
	con *GeTuiConfig
	err Error

	cache Cache
}

func NewGetui(con *GeTuiConfig) *Getui {
	// 检查配置
	if len(con.appId) == 0 {
		panic("请先加载配置")
	}

	// 设置超时
	SetTimeout(time.Duration(con.RequestTimeout) * time.Second)

	return &Getui{
		con: con,
	}
}

func (g *Getui) SetCache(c Cache) *Getui {
	g.cache = c
	return g
}

// 删除Token
func (g *Getui) DeleteToken(token string) error {
	if g.cache != nil {
		err := g.cache.Delete()
		if err != nil {
			g.err = append(g.err, err)
		}
	}

	_, err := Do("DELETE", g.url(fmt.Sprintf("auth/%s", token)), "", nil)
	if err != nil {
		g.err = append(g.err, err)
	}

	return g.hasError()
}

func (g *Getui) checkStringLen(s string, len int) bool {
	return utf8.RuneCountInString(s) <= len
}

func (g *Getui) checkRequestId(r string) bool {
	l := utf8.RuneCountInString(r)
	return l < 10 || l > 32
}

func (g *Getui) url(path string) string {
	return fmt.Sprintf(BaseUrl, g.con.appId, path)
}

func (g *Getui) token() string {
	if g.cache != nil {
		token := g.cache.Get()
		if len(token) > 0 {
			return token
		}
	}

	sign, tp := Signature(g.con.appKey, g.con.masterSecret)

	body := map[string]string{
		"sign":      sign,
		"timestamp": tp,
		"appkey":    g.con.appKey,
	}

	resp, err := Do("POST", g.url("auth"), "", body)
	if err != nil || resp.Data == nil {
		g.err = append(g.err, fmt.Errorf("token err: %v", err))
		return ""
	}

	expireTime := cast.ToInt64(resp.Data["expire_time"])
	token := cast.ToString(resp.Data["token"])

	if g.cache != nil {
		err := g.cache.Save(token, expireTime)
		if err != nil {
			g.err = append(g.err, err)
		}
	}
	return token
}

func (g *Getui) hasError() error {
	if g.err.AsError() != nil {
		return g.err.AsError()
	}
	return nil
}
