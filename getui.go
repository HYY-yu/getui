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
	配置中心添加配置文件
	config_key_hooker加载好配置

	resp,err := NewGetui(GeTuiCon).ToSingleCid(
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

// cid单推
func (g *Getui) ToSingleCid(req *Req) (*Resp, error) {
	if g.checkRequestId(req.RequestId) {
		return nil, ErrorRequestIdLen
	}

	if len(req.Audience.Cid) != 1 {
		return nil, errors.New(" The length of cid must be 1 ")
	}

	resp, err := Do("POST", g.url("push/single/cid"), g.token(), req)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}

// 别名单推
func (g *Getui) ToSingleAlias(req *Req) (*Resp, error) {
	if g.checkRequestId(req.RequestId) {
		return nil, ErrorRequestIdLen
	}

	if len(req.Audience.Alias) != 1 {
		return nil, errors.New(" The length of cid must be 1")
	}

	resp, err := Do("POST", g.url("push/single/alias"), g.token(), req)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}

// cid批量单推
func (g *Getui) ToSingleBatchCid(req []Req, isAsync bool) (*Resp, error) {
	if len(req) <= 0 || len(req) > 200 {
		return nil, errors.New(" The length of req in (0,200] ")
	}

	for _, e := range req {
		if g.checkRequestId(e.RequestId) {
			return nil, ErrorRequestIdLen
		}

		if len(e.Audience.Cid) != 1 {
			return nil, errors.New(" The length of cid must be 1 ")
		}
	}

	data := map[string]interface{}{
		"is_async": isAsync,
		"msg_list": req,
	}

	resp, err := Do("POST", g.url("push/single/batch/cid"), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}

// 别名批量单推
func (g *Getui) ToSingleBatchAlias(req []Req, isAsync bool) (*Resp, error) {
	if len(req) <= 0 || len(req) > 200 {
		return nil, errors.New(" The length of req in (0,200] ")
	}

	for _, e := range req {
		if g.checkRequestId(e.RequestId) {
			return nil, ErrorRequestIdLen
		}

		if len(e.Audience.Alias) != 1 {
			return nil, errors.New(" The length of cid must be 1 ")
		}
	}

	data := map[string]interface{}{
		"is_async": isAsync,
		"msg_list": req,
	}

	resp, err := Do("POST", g.url("push/single/batch/alias"), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}

// toList创建消息
func (g *Getui) ToListMessage(req *Req) (taskId string, resp *Resp, err error) {
	if g.checkRequestId(req.RequestId) {
		return "", nil, ErrorRequestIdLen
	}

	resp, err = Do("POST", g.url("push/list/message"), g.token(), req)
	if err != nil {
		g.err = append(g.err, err)
		return "", nil, g.hasError()
	}

	if v, ok := resp.Data["taskid"]; ok {
		return cast.ToString(v), resp, nil
	}

	return "", nil, fmt.Errorf("not taskid: %v", resp)
}

// toList Cid批量推
func (g *Getui) ToListCid(taskId string, isAsync bool, audi *Audience) (cidMap map[string]interface{}, resp *Resp, err error) {
	if len(taskId) == 0 {
		return nil, nil, errors.New("must have taskId")
	}

	if audi == nil || len(audi.Cid) == 0 {
		return nil, nil, errors.New("must have Audience Cid ")
	}

	data := map[string]interface{}{
		"taskid":   taskId,
		"is_async": isAsync,
		"audience": audi,
	}

	resp, err = Do("POST", g.url("push/list/cid"), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
		return nil, nil, g.hasError()
	}

	if v, ok := resp.Data[taskId]; ok {
		if vv, ok := v.(map[string]interface{}); ok {
			return vv, resp, nil
		}
	}
	return nil, resp, fmt.Errorf("error format resp : %v", resp)
}

// toList 别名批量推
func (g *Getui) ToListAlias(taskId string, isAsync bool, audi *Audience) (cidMap map[string]interface{}, resp *Resp, err error) {
	if len(taskId) == 0 {
		return nil, nil, errors.New("must have taskId")
	}

	if audi == nil || len(audi.Alias) == 0 {
		return nil, nil, errors.New("must have Audience Alias")
	}

	data := map[string]interface{}{
		"taskid":   taskId,
		"is_async": isAsync,
		"audience": audi,
	}

	resp, err = Do("POST", g.url("push/list/alias"), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
		return nil, nil, g.hasError()
	}

	if v, ok := resp.Data[taskId]; ok {
		if vv, ok := v.(map[string]interface{}); ok {
			return vv, resp, nil
		}
	}
	return nil, resp, fmt.Errorf("error format resp : %v", resp)
}

// 执行群推
func (g *Getui) ToApp(req *Req) (taskId string, resp *Resp, err error) {
	if g.checkRequestId(req.RequestId) {
		return "", nil, ErrorRequestIdLen
	}

	if req.Audience.All != "all" {
		return "", nil, errors.New("audience.all must be <all> ")
	}

	resp, err = Do("POST", g.url("push/all"), g.token(), req)
	if err != nil {
		g.err = append(g.err, err)
		return "", nil, g.hasError()
	}

	if v, ok := resp.Data["taskid"]; ok {
		return cast.ToString(v), resp, err
	}
	return "", resp, fmt.Errorf("error format resp : %v", resp)

}

// 根据条件筛选用户推送
func (g *Getui) ToAppTag(req *Req) (taskId string, resp *Resp, err error) {
	if g.checkRequestId(req.RequestId) {
		return "", nil, ErrorRequestIdLen
	}

	if len(req.Audience.Tag) == 0 {
		return "", nil, errors.New("The length of audience.Tag is zero ")
	}

	resp, err = Do("POST", g.url("push/tag"), g.token(), req)
	if err != nil {
		g.err = append(g.err, err)
		return "", nil, g.hasError()
	}

	if v, ok := resp.Data["taskid"]; ok {
		return cast.ToString(v), resp, err
	}
	return "", resp, fmt.Errorf("error format resp : %v", resp)
}

// 使用标签快速推送
func (g *Getui) ToAppCustomTag(req *Req) (taskId string, resp *Resp, err error) {
	if g.checkRequestId(req.RequestId) {
		return "", nil, ErrorRequestIdLen
	}

	if len(req.Audience.FastCustomTag) == 0 {
		return "", nil, errors.New("The length of audience.Tag is zero ")
	}

	resp, err = Do("POST", g.url("push/fast_custom_tag"), g.token(), req)
	if err != nil {
		g.err = append(g.err, err)
		return "", nil, g.hasError()
	}

	if v, ok := resp.Data["taskid"]; ok {
		return cast.ToString(v), resp, err
	}
	return "", resp, fmt.Errorf("error format resp : %v", resp)
}

// 某用户绑多个标签
func (g *Getui) CustomTagsToUser(cid string, tags ...string) (*Resp, error) {
	data := map[string]interface{}{
		"custom_tag": tags,
	}

	resp, err := Do("POST", g.url(fmt.Sprintf("user/custom_tag/cid/%s", cid)), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
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
