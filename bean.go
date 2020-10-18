package getui


type Resp struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
}

//http://docs.getui.com/getui/server/rest_v2/push/
// 使用时对照个推文档

// omitempty 代表此参数是可选的

type Req struct {
	RequestId   string       `json:"request_id"`
	GroupName   string       `json:"group_name,omitempty"`
	Audience    Audience     `json:"audience"`
	Setting     *Setting     `json:"setting,omitempty"`
	PushMessage PushMessage  `json:"push_message"`
	PushChannel *PushChannel `json:"push_channel,omitempty"`
}

type Audience struct {
	Cid           []string      `json:"cid,omitempty"`
	Alias         []string      `json:"alias,omitempty"`
	FastCustomTag string        `json:"fast_custom_tag,omitempty"`
	Tag           []AudienceTag `json:"tag,omitempty"`
	All           string        `json:"all,omitempty"`
}

type AudienceTag struct {
	Key     string   `json:"key"`
	Values  []string `json:"values"`
	OptType string   `json:"opt_type"`
}

type Setting struct {
	Ttl          int                    `json:"ttl,omitempty"`
	Strategy     map[string]interface{} `json:"strategy,omitempty"`
	Speed        int                    `json:"speed,omitempty"`
	ScheduleTime int                    `json:"schedule_time,omitempty"`
}

type PushMessage struct {
	Duration     string        `json:"duration,omitempty"`
	Notification *Notification `json:"notification,omitempty"` // 注意，此字段只适用Android系统
	Transmission string        `json:"transmission,omitempty"` // Android\IOS 均可
}

type Notification struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	ClickType    string `json:"click_type"`
	BigText      string `json:"big_text,omitempty"`
	BigImage     string `json:"big_image,omitempty"`
	Logo         string `json:"logo,omitempty"`
	LogoUrl      string `json:"logo_url,omitempty"`
	ChannelId    string `json:"channel_id,omitempty"`
	ChannelName  string `json:"channel_name,omitempty"`
	ChannelLevel string `json:"channel_level,omitempty"`
	Intent       string `json:"intent,omitempty"`
	Url          string `json:"url,omitempty"`
	Payload      string `json:"payload,omitempty"`
	NotifyId     int    `json:"notify_id,omitempty"`
}

// 厂商通道消息内容
type PushChannel struct {
	Ios     *IOS     `json:"ios,omitempty"`
	Android *Android `json:"android,omitempty"`
}

// TODO 可以继续展开
type Android struct {
	Ups map[string]interface{} `json:"ups,omitempty"`
}

type IOS struct {
	Type           string       `json:"type,omitempty"`
	Aps            *IOSAps      `json:"aps,omitempty"`
	AutoBadge      string       `json:"auto_badge,omitempty"`
	Payload        string       `json:"payload,omitempty"`
	Multimedia     []Multimedia `json:"multimedia,omitempty"`
	ApnsCollapseId string       `json:"apns_collapse_id,omitempty"`
}

type IOSAps struct {
	Alert            Alert  `json:"alert,omitempty"`
	ContentAvailable int    `json:"content_available,omitempty"`
	Sound            string `json:"sound,omitempty"`
	Category         string `json:"category,omitempty"`
	ThreadId         string `json:"thread_id,omitempty"`
}

type Alert struct {
	Title           string   `json:"title,omitempty"`
	Body            string   `json:"body,omitempty"`
	ActionLocKey    string   `json:"action_loc_key,omitempty"`
	LocKey          string   `json:"loc_key,omitempty"`
	LocArgs         []string `json:"loc_args,omitempty"`
	LaunchImage     string   `json:"launch_image,omitempty"`
	TitleLocKey     string   `json:"title_loc_key,omitempty"`
	TitleLocArgs    []string `json:"title_loc_args,omitempty"`
	Subtitle        string   `json:"subtitle,omitempty"`
	SubtitleLocKey  string   `json:"subtitle_loc_key,omitempty"`
	SubtitleLocArgs []string `json:"subtitle_loc_args,omitempty"`
}

type Multimedia struct {
	Url      string `json:"url"`
	Type     int    `json:"type"`
	OnlyWifi bool   `json:"only_wifi,omitempty"`
}
