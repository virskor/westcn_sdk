package westcn

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"
)

// 发送请求的选项
type RequestOption struct {
	// 请求的路径
	Path string `json:"path"`

	// 请求的方式，默认为GET
	Method string `json:"method"`

	// 请求的Header，将在请求时同时发送
	Header map[string]interface{} `json:"header"`

	// FormData，传入
	FormData url.Values `json:"body"`
}

// 用于请求的Token
type RequestToken struct {
	// 请求的方式，默认为GET
	Value string `json:"value"`

	// token生成的时间
	Time int64 `json:"time"`
}

// 取得token
// 将字符串username与字符串api_password连接，再与timestamp连接，然后将生成的字符串进行md5求值，md5算法要求为：
// 32位16进制字符串，小写格式。
// 身份验证串有效期10分钟。
// 比如，您的西部数码用户名为：zhangsan，您的API密码为：5dh232kfg!* ,当前毫秒时间戳为：1554691950854，则：
// token = md5(zhangsan + 5dh232kfg!* + 1554691950854) = f17581fb2535b2a7ee4468eb3f96a2a9
func (w *WestClient) getToken() *RequestToken {
	now := time.Now()
	token := &RequestToken{
		Time: now.UnixMicro(),
	}

	// 计算token值
	tokenValue := fmt.Sprintf("%s%s%d", w.Options.Username, w.Options.ApiPassword, token.Time)
	token.Value = w.md5(tokenValue)

	return token
}

// 将字符串转换为md5
func (w *WestClient) md5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 发送请求
func (w *WestClient) Request() {
	token := w.getToken()
	println(token)

}
