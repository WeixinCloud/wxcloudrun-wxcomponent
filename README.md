# wxcloudrun-wxcomponent
[![GitHub license](https://img.shields.io/github/license/WeixinCloud/wxcloudrun-wxcomponent)](https://github.com/WeixinCloud/wxcloudrun-wxcomponent)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/WeixinCloud/wxcloudrun-wxcomponent)

微信云托管 微信第三方平台管理工具模版

## 功能介绍
此项目提供第三方平台的后端服务以及第三方平台管理工具。该镜像可一键部署到微信云托管，分钟级别即可完成第三方平台开发环境搭建以及第三方平台管理工具部署。详情参考官方文档：[服务商微管家介绍](https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/product/management-tools.html)。

![index](https://res.wx.qq.com/op_res/BF2B0NQ2bKt-rJQL--cB3fUuCyllmnvJdFT57k786XuTE5UJQh4x8KjxiaGsg48qsqLtlP1kCZcr7E48DKq2xg)

#### 第三方平台推送消息
微信第三方平台需要填写两个URL用于接受官方推送的消息，详情参考官方文档：[创建与配置第三方平台准备工作](https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/operation/thirdparty/prepare.html)。
- 授权事件URL: 本项目提供了接受官方推送并存入数据库的服务，对推送ticket、授权、解除授权的事件都做了相应处理。
- 消息与事件URL: 本项目提供了接受官方推送并存入数据库的服务，开发者可以读数据库查看推送消息，也可以在此基础上进行二次开发。
#### 第三方平台管理工具
- 授权帐号管理：可查看授权给第三方平台的公众号/小程序帐号信息。
- 第三方token获取：可一键获取component_verify_ticket、component_access_token、authorizer_access_token以及微信令牌，便于开发者进行调试。
- 第三方消息查看：可获取推送至授权事件URL和消息与事件URL的消息，便于开发者进行调试。
- 第三方授权页面生成：可一键生成PC版和H5版的授权页面，商家可扫码或者直接访问授权页面完成授权。

## 目录结构
```
.
├── Dockerfile
├── README.md
├── api                                 // 后端api
│   ├── admin                           // 管理工具，需管理员登录
│   ├── authpage                        // 授权页面，无鉴权
│   ├── innerservice                    // 提供内部服务
│   ├── proxy                           // 代理
│   └── wxcallback                      // 接收微信消息
├── client                              // 前端
│   ├── dist                            // 打包结果
│   ├── index.html
│   ├── node_modules
│   ├── package.json
│   ├── src                             // 源代码
│   ├── tsconfig.json
│   ├── vite.config.ts
│   └── yarn.lock
├── comm                                // 后端公共模块
│   ├── config                          // 配置
│   ├── encrypt                         // 加密
│   ├── errno                           // 错误码
│   ├── httputils                       // http
│   ├── inits                           // 初始化
│   ├── log                             // 日志
│   ├── utils                           // 其他工具
│   └── wx                              // 微信相关
├── container.config.json               // 微信云托管初始化服务配置，二开不生效请忽略
├── db                                  // 数据库相关
│   ├── dao
│   ├── init.go
│   └── model
├── go.mod
├── go.sum
├── main.go
├── middleware                          // 中间件
│   ├── innerservice.go                 // 内部服务
│   ├── jwt.go                          // jwt
│   ├── log.go                          // 日志
│   └── wxsource.go                     // 判断是否为微信来源
└── routers                             // 路由
    └── routers.go

```

## 其他说明
#### 本地调试
服务启动前会从环境变量中读取数据库配置，自行写入环境变量后运行一下代码，即可在本地启动服务。
```
go run main
```

#### 判断微信来源
服务部署在微信云托管时，微信推送消息走内网，无需加解密，判断header中是否有x-wx-source即可。

#### 数据表
```
+-----------------------+
| Tables_in_wxcomponent |
+-----------------------+
| authorizers           |
| comm                  |
| counter               |
| user                  |
| wxcallback_biz        |
| wxcallback_component  |
| wxcallback_rules      |
| wxtoken               |
+-----------------------+
```
- authorizers: 授权账号信息
- comm: 存储ticket、第三方信息等
- user: 用户表
- wxcallback_biz: 推送给消息与事件URL的消息
- wxcallback_component: 推送给授权事件URL的消息
- wxcallback_rules: 消息转发规则
- wxtoken: component_access_token和authorizer_access_token
- counter: 登录失败计数
#### 命名格式
- 微信开放平台接口: 下划线
- 微管家前后端交互: 小驼峰
- 微信回调消息: 大驼峰

## 使用注意
如果不是通过微信云托管控制台部署模板代码，而是自行复制/下载模板代码后，手动新建一个服务并部署，需要在「服务设置」中补全以下环境变量，才可正常使用，否则会引发无法连接数据库，进而导致部署失败。

- MYSQL_ADDRESS
- MYSQL_PASSWORD
- MYSQL_USERNAME 
以上三个变量的值请按实际情况填写。如果使用云托管内MySQL，可以在控制台MySQL页面获取相关信息。


## License

[MIT](./LICENSE)
