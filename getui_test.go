package getui

import (
	"fmt"
	"testing"
)

func TestGetui_ToSingleCid(t *testing.T) {
	g := &GeTuiConfig{
		RequestTimeout: 2,
	}

	resp, err := NewGetui(g).SetCache(MockCache{}).ToSingleCid(&Req{
		RequestId: "asdfghjkl123456",
		Audience: Audience{
			Cid: []string{"5c0bce442cb4d708da0501ce6e2c0d2a"},
		},
		PushMessage: PushMessage{
			Notification: &Notification{
				Title:     "测试",
				Body:      "测试消息",
				ClickType: "none",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success " + fmt.Sprintf("%v", resp))
}

func TestGetui_ToSingleAlias(t *testing.T) {
	g := &GeTuiConfig{
		RequestTimeout: 2,
	}

	resp, err := NewGetui(g).SetCache(MockCache{}).ToSingleAlias(&Req{
		RequestId: "asdfghjkl123459",
		Audience: Audience{
			Alias: []string{"yf"},
		},
		PushMessage: PushMessage{
			Notification: &Notification{
				Title:     "测试",
				Body:      "测试消息",
				ClickType: "none",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success " + fmt.Sprintf("%v", resp))
}

func TestGetui_ToSingleBatchCid(t *testing.T) {
	g := &GeTuiConfig{
		RequestTimeout: 2,
	}

	resp, err := NewGetui(g).SetCache(MockCache{}).ToSingleBatchCid([]Req{
		{
			RequestId: "asdfghjkl1234513",
			Audience: Audience{
				Cid: []string{"5c0bce442cb4d708da0501ce6e2c0d2a"},
			},
			PushMessage: PushMessage{
				Notification: &Notification{
					Title:     "测试ToSingleBatchCid",
					Body:      "测试消息",
					ClickType: "none",
				},
			},
		},
	}, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success " + fmt.Sprintf("%v", resp))
}

func TestGetui_ToSingleBatchAlias(t *testing.T) {
	g := &GeTuiConfig{
		RequestTimeout: 2,
	}

	resp, err := NewGetui(g).SetCache(MockCache{}).ToSingleBatchAlias([]Req{
		{
			RequestId: "asdfghjk211234513",
			Audience: Audience{
				Alias: []string{"yf"},
			},
			PushMessage: PushMessage{
				Notification: &Notification{
					Title:     "测试ToSingleBatchAlias",
					Body:      "测试消息",
					ClickType: "none",
				},
			},
		},
	}, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success " + fmt.Sprintf("%v", resp))
}

func TestGetui_ToAppTag(t *testing.T) {
	g := &GeTuiConfig{
		RequestTimeout: 2,
	}

	taskId, resp, err := NewGetui(g).SetCache(MockCache{}).ToAppTag(&Req{
		RequestId: "psdfghjk211234513",
		Audience: Audience{
			Tag: []AudienceTag{
				{
					Key:     "custom_tag",
					Values:  []string{"yufeng"},
					OptType: "and",
				},
			},
		},
		PushMessage: PushMessage{
			Notification: &Notification{
				Title:     "测试ToAppTag",
				Body:      "测试消息",
				ClickType: "none",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success " + fmt.Sprintf("%v", resp) + " _ " + taskId)
}

func TestGetui_CustomTagsToUser(t *testing.T) {
	g := &GeTuiConfig{
		RequestTimeout: 2,
	}

	resp, err := NewGetui(g).SetCache(MockCache{}).
		CustomTagsToUser("5c0bce442cb4d708da0501ce6e2c0d2a", "yufeng")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success " + fmt.Sprintf("%v", resp))
}
