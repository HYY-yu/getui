package getui

import "fmt"

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
