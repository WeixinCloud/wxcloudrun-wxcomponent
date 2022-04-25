# 微管家前端

### 项目技术栈相关文档
vite：https://vitejs.cn/guide/ \
React：https://react.docschina.org/ \
TDesign-React：https://tdesign.tencent.com/react/components/overview

## 如何进行二次开发
只新增页面、逻辑的二次开发务必阅读下面介绍（后续升级微管家版本会相对方便）

最低限度的二次开发不需要去阅读理解大量官方本身的代码，项目结构已分割出自定义代码区并暴露相关能力位于 [src/custom](./src/custom)

用户自己开发的新增页面 建议都在 src/custom 下进行相关开发，这样在升级微管家版本时可无痛merge代码

在 custom 下编写相关的页面文件之后，修改导航栏和路由相关配置即可生效

左侧导航栏和页面配置的相关声明文件位于 [src/commonType.ts](./src/commonType.ts)

其中页面路由的配置在 [custom/config/route.tsx](./src/custom/config/route.tsx) 用法如下：
```typescript jsx
export const customRoute: IRoute = {
    demo: {
        label: '一个demo',
        path: '/demo',
        element: <Demo />
    },
    demo1: {
        label: '一个demo1',
        path: '/demo1',
        element: <Demo1 />
    },
    demo2: {
        label: '一个demo2',
        path: '/demo2',
        showPath: '/demo1',
        element: <Demo2 />
    },
    demo3: {
        label: '一个demo',
        path: '/demo',
        element: <Demo />,
        dontNeedMenu: true // 左侧不展示menu，例如登录页，即也不需要登录态
    },
}
```
其中左侧导航栏的配置在 [custom/config/menu.tsx](./src/custom/config/menu.tsx) 用法如下：
```typescript jsx
import { customRoute as routes } from './routes'
export const customMenuList: IMenuList = [{
    label: '测试路由',
    icon: <Icon.AppIcon />,
    item: [routes.demo, {
        ...routes.demo1,
        hideItem: [routes.demo2],// demo1的子级路由，但在导航menu中不显示
    }]
}]
```
项目中用到的所有接口 [custom/utils/apis.ts](custom/utils/apis.ts)

官方也同样暴露了二开可能用到的相关能力的函数方法，位于 [custom/utils/common.ts](./src/custom/utils/common.ts) 其中有以下方法，用法如下：
```typescript
copyMessage('要复制的文本') // 复制文本

checkLogin() // 判断是否登录 true 已登录 false 未登录

logout() // 退出登录

refreshToken() // 刷新登录token

import { apis } from './apis.ts'
request({
    request: apis.getTicketRequest,
    noNeedCheckLogin: true,
    data: {}
}) // 发起请求方法

request({
    request: {
        method: 'post',
        url: 'xxxxx',
    },
    data: {}
}) // 发起请求方法
```


## 项目开发
#### 项目页面路由及对应组件见[route 配置](./src/config/route.tsx)
#### 项目左侧导航栏配置见[menu 配置](./src/config/menu.tsx)
#### 安装依赖
```shell
 yarn 或 npm install
```
#### 修改 proxy 设置：将[vite 配置文件](./vite.config.ts)中的 proxy.target 改为自己的地址
#### 启动项目
```shell
 yarn dev 或 npm run dev
```
#### 项目打包
```shell
 yarn build 或 npm run build
```
#### 项目部署
见项目根目录 [README.md](../README.md)

## 目录结构
```
.
├── dist                                // 打包编译产物
├── scripts                             // 一些脚本文件
│   └── checkHost.mjs                   // 校验push代码是否携带敏感信息
├── src                                 // 具体前端资源
│   ├── assets                          // 项目静态资源
│   │   └── icons                       // icon类文件存放
│   │      └── xxxx.png                 // 具体的icon文件
│   ├── config                          // 配置存放
│   │   ├── menu.tsx                    // 左侧导航栏配置
│   │   └── route.tsx                   // 路由配置
│   ├── custom                          // 二开代码放置位置
│   │   ├── config                      // 二开配置文件
│   │   │   ├── menu.tsx                // 二开左侧导航栏配置
│   │   │   └── route.tsx               // 二开路由配置
│   │   └── utils                       // 二开工具函数
│   │       ├── apis.ts                 // 二开接口信息
│   │       └── common.ts               // 二开通用函数
│   ├── components                      // 公共组件
│   │   ├── Console                     // 控制台组件，各个页面的共用部分
│   │   │   ├── index.module.less       // 组件级样式文件
│   │   │   └── index.tsx               // 组件具体代码逻辑文件
│   │   └── Menu                        // 左侧导航栏组件
│   │         └── .....
│   ├── pages                           // 具体页面文件
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
│   │   ├── common.ts                   // 一些常用的方法函数
│   │   └── login.ts                    // 一些登录相关逻辑函数
│   ├── commonType.ts                   // 公共声明文件
│   ├── icon.ico                        // 项目 ico 文件
│   ├── main.less                       // 一些项目的全局样式
│   ├── main.tsx                        // 根文件 main.tsx
│   └── vite-env.d.ts                   // 环境变量配置
├── .editorconfig                       // 编辑器配置
├── .gitignore                          // git 忽略上传文件
├── index.html                          // 入口文件 index.html
├── package.json                        // 项目依赖
├── README.md
├── tsconfig.json                       // ts 配置文件
├── vite.config.ts                      // vite 配置文件
└── 接口文档.md

```
