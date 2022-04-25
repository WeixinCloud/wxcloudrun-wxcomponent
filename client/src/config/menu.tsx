import * as Icon from "tdesign-icons-react";
import {routes} from "./route";
import {IMenuList} from "../commonType";
import {customMenuList} from "../custom/config/menu";

export const menuList: IMenuList = [{
    ...routes.home,
    icon: <Icon.HomeIcon />,
}, {
    label: '管家中心',
    icon: <Icon.AppIcon />,
    item: [routes.authPageManage, {
        ...routes.authorizedAccountManage,
        hideItem: [routes.submitAudit, routes.miniProgramVersion],
    }]
}, {
    label: '开发辅助',
    icon: <Icon.ViewListIcon />,
    item: [{
        ...routes.developTools,
        hideItem: [routes.thirdToken, routes.thirdMessage],
    }, routes.forwardMessage, routes.proxyConfig]
}, {
    label: '系统管理',
    icon: <Icon.SettingIcon />,
    item: [routes.passwordManage, routes.systemVersion]
}, ...customMenuList]
