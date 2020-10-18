package getui

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

type UserAliasReqParam struct {
	Cid   string `json:"cid"`
	Alias string `json:"alias"`
}

// 绑定别名
func (g *Getui) UserAlias(datas []UserAliasReqParam) (*Resp, error) {
	if len(datas) == 0 {
		return nil, errors.New("no data")
	}

	data := map[string]interface{}{
		"data_list": datas,
	}

	resp, err := Do("POST", g.url("user/alias"), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 根据Cid查询别名
func (g *Getui) QueryCidAlias(cid string) (alias string, resp *Resp, err error) {
	if len(cid) == 0 {
		return "", nil, errors.New("no data")
	}

	resp, err = Do("GET", g.url(fmt.Sprintf("user/alias/cid/%s", cid)), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
		return "", resp, g.hasError()

	}

	if v, ok := resp.Data["alias"]; ok {
		return cast.ToString(v), resp, nil
	}
	return "", resp, errors.New("The alias not find in resp ")
}

// 根据别名查cid
func (g *Getui) QueryAliasCid(alias string) (cids []string, resp *Resp, err error) {
	if len(alias) == 0 {
		return nil, nil, errors.New("no data")
	}

	resp, err = Do("GET", g.url(fmt.Sprintf("user/cid/alias/%s", alias)), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
		return nil, resp, g.hasError()

	}

	if v, ok := resp.Data["cid"]; ok {
		return cast.ToStringSlice(v), resp, nil
	}
	return nil, resp, errors.New("The cid not find in resp ")
}

// 批量解绑别名
func (g *Getui) DeleteUserAlias(datas []UserAliasReqParam) (*Resp, error) {
	if len(datas) == 0 {
		return nil, errors.New("no data")
	}

	data := map[string]interface{}{
		"data_list": datas,
	}

	resp, err := Do("DELETE", g.url("user/alias"), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 解绑所有别名
func (g *Getui) DeleteAllAlias(alias string) (*Resp, error) {
	if len(alias) == 0 {
		return nil, errors.New("no data")
	}

	resp, err := Do("DELETE", g.url(fmt.Sprintf("user/alias/%s", alias)), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 某用户绑多个标签
func (g *Getui) UserInCustomTags(cid string, tags ...string) (*Resp, error) {
	if len(cid) == 0 {
		return nil, errors.New("no data")
	}

	data := map[string]interface{}{
		"custom_tag": tags,
	}

	resp, err := Do("POST", g.url(fmt.Sprintf("user/custom_tag/cid/%s", cid)), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 一批用户绑一个标签
func (g *Getui) CustomTagsToUsers(tag string, cids ...string) (*Resp, error) {
	if len(tag) == 0 {
		return nil, errors.New("no data")
	}

	data := map[string]interface{}{
		"cid": cids,
	}

	resp, err := Do("PUT", g.url(fmt.Sprintf("user/custom_tag/batch/%s", tag)), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 解绑标签
func (g *Getui) DeleteCustomTags(tag string, cids ...string) (*Resp, error) {
	if len(tag) == 0 {
		return nil, errors.New("no data")
	}

	data := map[string]interface{}{
		"cid": cids,
	}

	resp, err := Do("DELETE", g.url(fmt.Sprintf("user/custom_tag/batch/%s", tag)), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 查询用户标签
func (g *Getui) QueryCustomTags(cid string) (tags []string, resp *Resp, err error) {
	if len(cid) == 0 {
		return nil, nil, errors.New("no data")
	}

	resp, err = Do("GET", g.url(fmt.Sprintf("user/custom_tag/cid/%s", cid)), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
		return nil, resp, g.hasError()

	}

	if v, ok := resp.Data[cid]; ok {
		return cast.ToStringSlice(v), resp, nil
	}
	return nil, resp, errors.New("The cid not find in resp ")
}

// 添加黑名单用户
func (g *Getui) AddBlackUser(cids ...string) (*Resp, error) {
	if len(cids) == 0 {
		return nil, errors.New("no data")
	}

	resp, err := Do("POST", g.url(fmt.Sprintf("user/black/cid/%s", strings.Join(cids, ","))), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 删除黑名单用户
func (g *Getui) DeleteBlackUser(cids ...string) (*Resp, error) {
	if len(cids) == 0 {
		return nil, errors.New("no data")
	}

	resp, err := Do("DELETE", g.url(fmt.Sprintf("user/black/cid/%s", strings.Join(cids, ","))), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 查询用户状态
// 这个没解析返回值，可自行从resp中解析
func (g *Getui) QueryUserStatus(cids ...string) (*Resp, error) {
	if len(cids) == 0 {
		return nil, errors.New("no data")
	}

	resp, err := Do("GET", g.url(fmt.Sprintf("user/status/%s", strings.Join(cids, ","))), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 设置角标（IOS）
func (g *Getui) SetIOSBadge(badge string, cids ...string) (*Resp, error) {
	if len(cids) == 0 {
		return nil, errors.New("no data")
	}

	data := map[string]interface{}{
		"badge": badge,
	}

	resp, err := Do("POST", g.url(fmt.Sprintf("user/badge/cid/%s", strings.Join(cids, ","))), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
}

// 查询用户总量
func (g *Getui) CountUser(req []AudienceTag) (count int, resp *Resp, err error) {
	if len(req) == 0 {
		return 0, nil, errors.New("no data")
	}

	data := map[string]interface{}{
		"tag": req,
	}

	resp, err = Do("POST", g.url("user/count"), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
		return 0, resp, g.hasError()

	}

	if v, ok := resp.Data["user_count"]; ok {
		return cast.ToInt(v), resp, nil
	}
	return 0, resp, errors.New("The cid not find in resp ")
}
