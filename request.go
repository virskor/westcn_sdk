package westcn

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// 西部数码响应结构化数据
type Response struct {
	// 可能为空 返回代码 200 成功  其它为失败
	Result int `json:"result,omitempty"`

	// 必须 请求识别码
	ClientID string `json:"clientid"`

	// 可能为空 失败/成功返回的文本信息
	Msg string `json:"msg,omitempty"`

	// 可能为空 错误码（错误码含义详见 附录7.1 错误码说明）
	Errcode int `json:"errcode,omitempty"`

	// 可能为空 响应数据
	Data interface{} `json:"data,omitempty"`
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
		Time: now.UnixMilli(),
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
// 注意: Request返回和请求时，数据都被转换为GBK。返回的Response为原始的
func (w *WestClient) Request(option *RequestOption) (*Response, error) {

	// 计算Token
	token := w.getToken()

	// 重构请求的URL，将公共参数添加到Query中
	requestUrlStr := fmt.Sprintf("%s%s", w.Options.BaseApiPath, option.Path)
	requestUrl, err := url.Parse(requestUrlStr)
	if err != nil {
		return nil, err
	}

	/// 添加公共参数
	values := requestUrl.Query()
	values.Add("token", token.Value)
	values.Add("username", w.Options.Username)
	values.Add("time", fmt.Sprintf("%d", token.Time))
	requestUrl.RawQuery = values.Encode()

	httpClient := &http.Client{}

	// 默认为GET，不添加body
	var req *http.Request
	req, err = http.NewRequest("GET", requestUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	// 如果是POST，那么将req覆盖
	if option.Method == "POST" && option.FormData != nil {
		// 要将Post的数据全部转换为GBK才能进行提交
		for key, val := range option.FormData {
			str, _ := ToGBKString(val[0])
			option.FormData[key] = []string{
				str,
			}
		}

		// 再次构造请求参数
		formDataBytes := []byte(option.FormData.Encode())
		formBytesReader := bytes.NewReader(formDataBytes)
		req, err = http.NewRequest(option.Method, requestUrl.String(), formBytesReader)
		if err != nil {
			return nil, err
		}

		// 添加请求Header
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// 发送请求
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 取响应的数据
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 要将GBK转为UTF8
	respStr, err := ToUnicodeString(string(data))
	if err != nil {
		return nil, err
	}

	// json格式化数据
	var westResp Response
	err = json.Unmarshal([]byte(respStr), &westResp)
	if err != nil {
		return nil, err
	}

	// 发送请求，并将请求的结果转换为UTF-8
	return &westResp, nil
}
