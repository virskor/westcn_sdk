package westcn

import "errors"

const (
	// 默认的API请求地址
	DEFAULT_BASEAPIPATH = "https://api.west.cn/api/v2"
)

type WestClient struct {
	// 客户端选项
	Options *ClientOptions
}

type ClientOptions struct {
	// 客户端请求基础地址，默认为 https://api.west.cn/api/v2
	BaseApiPath string `json:"baseApiPath"`

	// 用户名
	Username string `json:"username"`

	// 西部数码ApiPassword
	ApiPassword string `json:"apiPassword"`
}

// 创建新的客户端
func NewClient(options *ClientOptions) (*WestClient, error) {
	if options == nil {
		return nil, errors.New("未传入客户端必要参数")
	}

	// 校验Username
	if len(options.Username) == 0 {
		return nil, errors.New("缺少Username")
	}

	// 校验ApiPassword
	if len(options.ApiPassword) == 0 {
		return nil, errors.New("缺少ApiPassword")
	}

	// 默认帮填写BaseApiPath
	if len(options.BaseApiPath) == 0 {
		options.BaseApiPath = DEFAULT_BASEAPIPATH
	}

	return &WestClient{
		Options: options,
	}, nil
}
