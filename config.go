package getui

type GeTuiConfig struct {
	appId        string
	appKey       string
	masterSecret string

	RequestTimeout int64 // 单位：秒
}