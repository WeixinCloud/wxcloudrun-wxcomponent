# 第三方平台管理工具前端

## 项目技术栈相关文档
### vite：https://vitejs.cn/guide/
### React：https://react.docschina.org/
### TDesign-React：https://tdesign.tencent.com/react/components/overview

## 目录结构
```
.
├── dist                                // 打包编译产物
├── src                                 // 具体前端资源
│   ├── components                      // 项目公共组件
│   │   ├── Console                     // 控制台组件，各个页面的共用部分
│   │   │   ├── index.module.less       // 组件级样式文件
│   │   │   └── index.tsx               // 组件具体代码逻辑文件
│   │   └── Menu                        // 左侧导航栏组件
│   │         └── .....
│   ├── pages                           // 项目具体页面文件
│   │   ├── authorizedAccount           // 授权账号管理页面
│   │   │   ├── index.module.less       // 页面级样式文件
│   │   │   └── index.tsx               // 页面具体代码逻辑文件
│   │   ├── authPage                    // pc 用户授权页面
│   │   │   └── .....
│   │   ├── authPageH5                  // h5 用户授权页面
│   │   │   └── .....
│   │   ├── authPageManage              // 授权页面管理页面
│   │   │   └── .....
│   │   ├── login                       // 登录页面
│   │   │   └── .....
│   │   ├── passwordManage              // 密码 secret 及相关信息管理页面
│   │   │   └── .....
│   │   ├── systemVersion               // 系统版本
│   │   │   └── .....
│   │   ├── thirdMessage                // 第三方信息查看页面
│   │   │   └── .....
│   │   └── thirdToken                  // 第三方 token 查看页面
│   │       └── .....
│   ├── utils                           // 项目一些通用工具类函数
│   │   ├── apis.ts                     // 接口请求地址统一管理
│   │   ├── axios.ts                    // 封装相关请求函数
│   │   └── login.ts                    // 一些登录相关逻辑函数
│   ├── icon.ico                        // 项目 icon 文件
│   ├── main.less                       // 一些项目的全局样式
│   ├── main.tsx                        // 项目根文件 main.tsx
│   └── vite-env.d.ts                   // 项目环境变量配置
├── .editorconfig                       // 编辑器配置
├── .gitignore                          // git 忽略上传文件
├── index.html                          // 项目入口文件 index.html
├── package.json                        // 项目依赖
├── README.md
├── tsconfig.json                       // ts 配置文件
├── vite.config.ts                      // vite 配置文件
├── yarn.lock
└── 接口文档.md

```

## 项目开发
### 项目具体页面路由见[Console 组件 routes](./src/components/Console/index.tsx)
### 路由对应页面组件见[main.tsx](./src/main.tsx)
### 安装依赖
```shell
 yarn
```
### 修改 proxy 设置：将[vite 配置文件](./vite.config.ts)中的 proxy.target 改为自己的地址
### 启动项目
```shell
 yarn dev
```
### 项目打包
```shell
 yarn build
```
### 项目部署
见项目 [wxcloudrun-wxcomponent](https://github.com/WeixinCloud/wxcloudrun-wxcomponent) README.md
