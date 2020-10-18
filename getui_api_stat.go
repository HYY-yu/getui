package getui

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 获取推送结果
func (g *Getui) QueryTasks(taskIds []string) (*Resp, error) {
	if len(taskIds) == 0 {
		return nil, errors.New("no data")
	}

	resp, err := Do("GET", g.url(fmt.Sprintf("report/push/task/count/%s", strings.Join(taskIds, ","))), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}

// 任务组名查报表
func (g *Getui) QueryTaskGroup(taskGourpName string) (*Resp, error) {
	if len(taskGourpName) == 0 {
		return nil, errors.New("no data")
	}

	resp, err := Do("GET", g.url(fmt.Sprintf("report/push/task_group/%s", taskGourpName)), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}

// 获取单日推送数据
func (g *Getui) QueryPushDataByDay(date time.Time) (*Resp, error) {
	if date.IsZero() {
		return nil, errors.New("no data")
	}

	resp, err := Do("GET", g.url(fmt.Sprintf("report/push/date/%s", date.Format("2006-01-02"))), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}

// 获取单日用户数据
func (g *Getui) QueryUserDataByDay(date time.Time) (*Resp, error) {
	if date.IsZero() {
		return nil, errors.New("no data")
	}

	resp, err := Do("GET", g.url(fmt.Sprintf("report/user/date/%s", date.Format("2006-01-02"))), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}

// 获取24个小时在线用户数
func (g *Getui) QueryOnlineUserBy24h() (*Resp, error) {
	resp, err := Do("GET", g.url("report/online_user"), g.token(), nil)
	if err != nil {
		g.err = append(g.err, err)
	}

	return resp, g.hasError()
}
