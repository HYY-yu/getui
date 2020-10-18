package getui

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
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
