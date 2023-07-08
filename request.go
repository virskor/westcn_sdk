package westcn

import "net/url"

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

// 发送请求
func (w *WestClient) Request() {
}
