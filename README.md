# 西数 SDK
West.cn API Request SDK for go
西部数码代理 API SDK，本SDK完成对West.cn部分API的封装，但仅限于开发者使用的部分，其他部分需要使用者自行封装。
内置GBK转UTF8，和Token自动计算，实名认证接口需自行拓展。
## 安装和简单使用

使用 go get 命令进行安装

```
go get github.com/virskor/westcn_sdk
```

创建西部数码 SDK 客户端，直接请求

```go
  client, err := westcn.NewClient(&westcn.ClientOptions{
    BaseApiPath: "https://api.west.cn/api/v2",
    Username: "username",
    ApiPassword: "api_password",
  })
  if err!= nil{
    panic(err)
  }

  // 准备提交的参数
  formData := url.Values{}
  formData.Add("domain", "west")
  formData.Add("suffix", ".com/.net")

  resp, err := client.Request(&westcn.RequestOption{
    Path:     westcn.PATH_DOMAIN_QUERY,
    Method:   "POST",
    FormData: formData,
  })
  if err != nil {
    panic(err)
  }
```
