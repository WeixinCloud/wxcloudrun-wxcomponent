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

把.env.example重命名为.env， 自行修改对应配置项， 运行后会自动加载。

```
MYSQL_USERNAME="admin"          #数据库用户名
MYSQL_PASSWORD="Aa123456"       #数据库密码
MYSQL_ADDRESS="127.0.0.1:3306"  #数据库地址
WX_APPID="wx################"   #微信APPID
```

服务启动前会从环境变量中读取数据库配置，自行写入环境变量后运行一下代码， 即可在本地启动服务。
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
1. 需打开云托管开放接口服务开关后才能免鉴权调用微信开放api。
  ![开放接口服务](https://res.wx.qq.com/op_res/6rcrTi7fRr5LStuAxoI94RrXbKG5L7OAiRfliINbA4qcM73YIjl7sMUgvSQBycgeKMBmsj7mJ2l5gTj1uCMlkA)
  说明：如果需要自行下载代码包部署，请确保comm/config/server.conf里的UseCloudBaseAccessToken为true

2. 如果不是通过开放平台一键部署服务，而是自行复制/下载模板代码后，手动新建一个服务并部署，需要在「服务设置」中补全以下环境变量，才可正常使用，否则会引发无法连接数据库，进而导致部署失败。

   - MYSQL_ADDRESS
   - MYSQL_PASSWORD
   - MYSQL_USERNAME 
   - WX_APPID

   以上四个变量的值请按实际情况填写。如果使用云托管内MySQL，可以在控制台MySQL页面获取相关信息。

3. 自行复制/下载模板代码包部署时安装有问题，则先关闭「云调用-开放接口服务」，等部署后，再重新打开再触发一次部署。开放接口服务配置接口填如下：

```
/cgi-bin/component/api_get_authorizer_list
/cgi-bin/component/api_get_authorizer_info
/cgi-bin/component/api_create_preauthcode
/cgi-bin/component/api_authorizer_token
/cgi-bin/component/api_component_token
/cgi-bin/component/api_query_auth
/cgi-bin/media/upload
/cgi-bin/componentloginpage
/wxa/submit_audit
/wxa/get_latest_auditstatus
/wxa/getvisitstatus
/wxa/getversioninfo
/wxa/getwxacodeunlimit
/wxa/get_qrcode
/wxa/gettemplatelist
/wxa/undocodeaudit
/wxa/speedupaudit
/wxa/commit
/wxa/release
/wxa/change_visitstatus
/wxa/revertcoderelease
/wxa/get_page
/wxa/get_category
```
4. 运行成功后，需在「系统管理-Secret与密码管理」中配置第三方平台的 Secret信息，部署后需要在第三方平台中，配置相应推送授权路径，[效果图](https://res.wx.qq.com/op_res/-hcDJVTUTTq70gAEYhzdVM2LYyhcKDHuLRX_4rEpTNjjwkqGWnvfuCmbLmAtI2LXtqvu-PnIMyGeH8TzQf-u9Q)

- 登录授权的发起页域名：直接配置服务商微管家的运行域名
- 推送路径按照图中配置
- 云托管环境和服务名称写自己的


6. 如需拉取部署微管家以前的授权数据，可向微管家服务发起数据同步请求，请求说明如下：
    - method: POST
    - path: wxcomponent/admin/pull-authorizer-list
    - header: `Authorization: Bear <your jwt>`

## License

[MIT](./LICENSE)
