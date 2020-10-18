package getui

import (
	"errors"
	"fmt"
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

	resp, err := Do("POST", g.url(fmt.Sprintf("user/alias")), g.token(), data)
	if err != nil {
		g.err = append(g.err, err)
	}
	return resp, g.hasError()
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
