package getui

import (
	"fmt"
)

// 实现Cache接口保存个推 token
type Cache interface {
	// 保存token
	// token: 个推token
	// expireTime: token过期时间（毫秒）
	Save(token string, expireTime int64) error
	// 获取token
	// 获取失败则为 ""
	Get() string
	// 删除缓存里的token
	Delete() error
}

// MockCache 提供给test文件使用
type MockCache struct {
	token string
}

// Mock save
// 内存缓存
func (m MockCache) Save(token string, expireTime int64) error {
	fmt.Println("token: ", token)
	m.token = token
	return nil
}

// Mock get
func (m MockCache) Get() string {
	return m.token
}

// Mock delete
func (m MockCache) Delete() error {
	m.token = ""
	return nil
}
