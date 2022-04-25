import { IRoute } from "../commonType";
import { customRoute } from "../custom/config/route";
import Home from "../pages/home";
import AuthorizedAccountManage from "../pages/authorizedAccountManage";
import AuthPageManage from "../pages/authPageManage";
import PasswordManage from "../pages/passwordManage";
import SystemVersion from "../pages/systemVersion";
import Login from "../pages/login";
import AuthPage from "../pages/authPage";
import AuthPageH5 from "../pages/authPageH5";
import DevelopTools from "../pages/developTools";
import ThirdToken from "../pages/developTools/thirdToken";
import ThirdMessage from "../pages/developTools/thirdMessage";
import ForwardMessage from "../pages/forwardMessage";
import ProxyConfig from "../pages/proxyConfig";
import RedirectPage from "../pages/redirectPage";
import MiniProgramVersion from "../pages/authorizedAccountManage/miniProgramVersion";
import SubmitAudit from "../pages/authorizedAccountManage/submitAudit";

export const routes: IRoute = {
    home: {
        label: '首页',
        path: '/home',
        element: <Home />,
    },
    authorizedAccountManage: {
        label: '授权帐号管理',
        path: '/authorizedAccountManage',
        element: <AuthorizedAccountManage />
    },
    authPageManage: {
        label: '授权链接生成器',
        path: '/authPageManage',
        element: <AuthPageManage />
    },
    passwordManage: {
        label: 'Secret与密码管理',
        path: '/passwordManage',
        element: <PasswordManage />
    },
    systemVersion: {
        label: '系统版本',
        path: '/systemVersion',
        element: <SystemVersion />
    },
    login: {
        label: '登录',
        path: '/login',
        element: <Login />,
        dontNeedMenu: true
    },
    authorize: {
        label: '授权页',
        path: '/authorize',
        dontNeedMenu: true,
        element: <AuthPage />
    },
    authorizeH5: {
        label: '授权页H5',
        path: '/authorizeH5',
        dontNeedMenu: true,
        element: <AuthPageH5 />
    },
    developTools: {
        label: '开发调试',
        path: '/developTools',
        element: <DevelopTools />
    },
    thirdToken: {
        label: '第三方 Token',
        path: '/developTools/token',
        showPath: '/developTools',
        element: <ThirdToken />
    },
    thirdMessage: {
        label: '第三方消息查看',
        path: '/developTools/message',
        showPath: '/developTools',
        element: <ThirdMessage />
    },
    forwardMessage: {
        label: '消息转发器',
        path: '/forwardMessage',
        element: <ForwardMessage />
    },
    proxyConfig: {
        label: 'proxy 配置',
        path: '/proxyConfig',
        element: <ProxyConfig />
    },
    redirectPage: {
        label: '授权回调跳转页',
        path: '/redirectPage',
        dontNeedMenu: true,
        element: <RedirectPage />
    },
    miniProgramVersion: {
        label: '版本管理',
        path: '/authorizedAccountManage/miniProgramVersion',
        showPath: '/authorizedAccountManage',
        element: <MiniProgramVersion />
    },
    submitAudit: {
        label: '提交审核',
        path: '/authorizedAccountManage/submitAudit',
        showPath: '/authorizedAccountManage',
        element: <SubmitAudit />
    },
    ...customRoute
}
