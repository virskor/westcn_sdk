# 西数 SDK

West.cn API Request SDK for go
西部数码代理 API SDK，本 SDK 完成 Go 对 West.cn 部分 API 的封装，但仅限于开发者使用的部分，其他部分需要使用者自行封装。

## 安装和简单使用

使用 go get 命令进行安装

```
go get github.com/virskor/westcn_sdk
```

创建西部数码 SDK 客户端

```go
  client, err := westcn.NewClient(&westcn.ClientOptions{
    BaseApiPath: "https://api.west.cn/api/v2",
    Username: "username",
    ApiPassword: "api_password",
  })
  if err!= nil{
    panic(err)
  }

  // 演示，仅使用client查询域名
  resp, err := client.Request(&westcn.RequestOption{
    Path: "/domain/query/",
    Method: "POST",
    Header: map[string]interface{}{
      "Content-Type": "application/x-www-form-urlencoded"
    },
    FormData: url.Values{
      "domain": {"west"},
      "suffix": {".com/.net"},
    },
  })
  if err!= nil{
    panic(err)
  }
```
