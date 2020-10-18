package getui

import (
	"fmt"
)

type Cache interface {
	Save(token string, expireTime int64) error
	Get() string
	Delete() error
}

type MockCache struct {
	token string
}

func (m MockCache) Save(token string, expireTime int64) error {
	fmt.Println("token: ", token)
	m.token = token
	return nil
}

func (m MockCache) Get() string {
	return m.token
}

func (m MockCache) Delete() error {
	m.token = ""
	return nil
}
