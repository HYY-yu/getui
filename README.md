个推 golangSDK
---
[![Go Report Card](https://goreportcard.com/badge/github.com/HYY-yu/getui)](https://goreportcard.com/report/github.com/HYY-yu/getui)

封装个推Rest API版本V2，目前github上的sdk都比较老，而且不是v2版本。故此实现一个。   
`https://restapi.getui.com/v2`   

功能特色：

- 杜绝`map[string]interface{}`，请求、返回都封装结构体，使用上更清晰
- Cache的依赖倒置，可以兼容不同的缓存实现
- 对error的统一处理
- 易于使用的API设计



使用方法
---

```golang
    resp,err := NewGetui(GeTuiConig{}).ToSingleCid(
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
```

建议一定要设置token的缓存，只要实现Cache接口即可。

```golang
    resp,err := NewGetui(?).SetCache(?).ToSingleCid(
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
```


> getui_test.go 有示例