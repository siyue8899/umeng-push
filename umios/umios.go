package umios

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/guonaihong/gout"
)

type Api struct {
	iosAppKey          string
	iosAppMasterSecret string
}

func NewPush(iosAppKey, iosAppMasterSecret string) *Api {
	return &Api{
		iosAppKey:          iosAppKey,
		iosAppMasterSecret: iosAppMasterSecret,
	}
}

// Push IOS推送
func (a *Api) Push(param PushTemplate) (*PushResp, error) {
	b, _ := json.Marshal(param)
	method := "POST"
	url := "https://msgapi.umeng.com/api/send"
	signStr := fmt.Sprintf("%s%s%s%s", method, url, b, a.iosAppMasterSecret)
	sign := GetMd5Encode(signStr)
	resp := PushResp{}
	err := gout.POST("https://msgapi.umeng.com/api/send?sign=" + sign).SetJSON(b).BindJSON(&resp).Do()
	return &resp, err
}
func GetMd5Encode(data string) string {
	h := md5.New()
	_, _ = h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
