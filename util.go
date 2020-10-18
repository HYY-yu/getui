package getui

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Signature(appKey string, masterSecret string) (string, string) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10) //签名开始生成毫秒时间

	//sha256(appkey+timestamp+mastersecret),mastersecret为注册应用时生成
	original := appKey + timestamp + masterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
}

type Error []error

func (me Error) Error() string {
	if me == nil {
		return ""
	}

	strs := make([]string, len(me))
	for i, err := range me {
		strs[i] = fmt.Sprintf("[%d] %v", i, err)
	}
	return strings.Join(strs, " ")
}

func (me Error) AsError() error {
	if len([]error(me)) <= 0 {
		return nil
	}

	return me
}