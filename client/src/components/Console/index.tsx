import React, {useEffect, useMemo, useState} from 'react'
import styles from './index.module.less'
import Menu from '../Menu'
import {Outlet, useNavigate, useLocation, Link} from "react-router-dom";
import * as Icon from 'tdesign-icons-react'
import {Dropdown, Dialog} from 'tdesign-react';
import {checkLogin, initNav, logout} from "../../utils/login";

export const routes = {
    home: {
      label: '首页',
      path: '/home'
    },
    authorizedAccountManage: {
        label: '授权帐号管理',
        path: '/authorizedAccountManage'
    },
    authPageManage: {
        label: '授权链接生成器',
        path: '/authPageManage'
    },
    passwordManage: {
        label: 'Secret与密码管理',
        path: '/passwordManage'
    },
    systemVersion: {
        label: '系统版本',
        path: '/systemVersion'
    },
    login: {
        label: '登录',
        path: '/login'
    },
    authorize: {
        label: '授权页',
        path: '/authorize'
    },
    authorizeH5: {
        label: '授权页H5',
        path: '/authorizeH5'
    },
    developTools: {
        label: '开发调试',
        path: '/developTools'
    },
    thirdToken: {
        label: '第三方 Token',
        path: '/developTools/token',
        showPath: '/developTools'
    },
    thirdMessage: {
        label: '第三方消息查看',
        path: '/developTools/message',
        showPath: '/developTools',
    },
}

type IMenuItem = {
    label: string
    path: string
    item?: IMenuItem[]
    showPath?: string
    hideItem?: IMenuItem[]
}

type IMenuList = {
    label: string
    icon: JSX.Element
    path?: string
    item?: IMenuItem[]
    hideItem?: IMenuItem[]
}[]

const menuList: IMenuList = [{
    ...routes.home,
    icon: <Icon.HomeIcon />,
}, {
    label: '管家中心',
    icon: <Icon.AppIcon />,
    item: [routes.authPageManage, routes.authorizedAccountManage]
}, {
    label: '开发辅助',
    icon: <Icon.ViewListIcon />,
    item: [{
        ...routes.developTools,
        hideItem: [routes.thirdToken, routes.thirdMessage]
    }]
}, {
    label: '系统管理',
    icon: <Icon.SettingIcon />,
    item: [routes.passwordManage, routes.systemVersion]
}]

const options = [
    {
        content: '微信开放平台',
        value: 'https://open.weixin.qq.com/',
    },
    {
        content: '微信第三方平台',
        value: 'https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/product/Third_party_platform_appid.html',
    },
];

const noticeOptions = [
    {
        content: '查看通知',
    },
];

export default function Console() {

    const [showNotice, setShowNotice] = useState<boolean>(false)
    const navigate = useNavigate()
    const location = useLocation()
    const [username] = useState(localStorage.getItem('username') || '')

    useEffect(() => {
        initNav(navigate)
        if (checkLogin()) {
            if (location.pathname === '/' || location.pathname === routes.login.path) {
                navigate(routes.home.path)
            }
        }
    }, [])

    const findLabel: (path: string, menu: IMenuList | IMenuItem, father?: IMenuItem) => JSX.Element | undefined = (path: string, menu: IMenuList | IMenuItem, father?: IMenuItem) => {
        if (Array.isArray(menu)) {
            for (let i = 0; i < menu.length; i++) {
                const result = findLabel(path, menu[i] as IMenuItem)
                if (result) return result
            }
        } else {
            if (menu.path === path) {
                if (menu.showPath) {
                    return <p className={styles.detail_header_title}><a href={`#${menu.showPath}`} className={`${styles.detail_header_title} a`}>{father?.label}</a> / {menu.label}</p>
                }
                return <p className={styles.detail_header_title}>{menu.label}</p>
            }
            if (menu.item) {
                for (let i = 0; i < menu.item.length; i++) {
                    const result = findLabel(path, menu.item[i], menu)
                    if (result) return result
                }
            }
            if (menu.hideItem) {
                for (let i = 0; i < menu.hideItem.length; i++) {
                    const result = findLabel(path, menu.hideItem[i], menu)
                    if (result) return result
                }
            }
        }
    }

    const headerLabel = useMemo(() => {
        return findLabel(location.pathname, menuList)
    }, [location.pathname])

    return (
        <div className={styles.console}>
            <div style={{width: '232px'}} />
            <span className={styles.console_menu}>
                <Menu menuList={menuList} />
            </span>
            <div className={styles.detail}>
                <div className={styles.detail_header}>
                    {headerLabel}
                    <div className={styles.detail_header_notice}>
                        <Dropdown maxColumnWidth={200} options={noticeOptions}
                                  onClick={() => setShowNotice(true)}>
                            <div className={styles.detail_header_notice_item}>
                                <Icon.NotificationIcon />
                                <p>通知</p>
                                <Icon.ChevronDownIcon />
                            </div>
                        </Dropdown>
                        <div className={styles.detail_header_notice_line} />
                        <Dropdown maxColumnWidth={200} options={options}
                                  onClick={(data) => window.open(data.value as string)}>
                            <div className={styles.detail_header_notice_item}>
                                <p>快捷链接</p>
                                <Icon.ChevronDownIcon />
                            </div>
                        </Dropdown>
                        <div className={styles.detail_header_notice_line} />
                        <p style={{ marginLeft: '15px' }}>{username}</p>
                        <p onClick={logout} style={{ margin: '0 15px', cursor: 'pointer' }}>退出</p>
                    </div>
                </div>
                <div className={styles.content}>
                    <Outlet />
                </div>
            </div>
            <Dialog header="通知中心" visible={showNotice} onConfirm={() => setShowNotice(false)}
                    onClose={() => setShowNotice(false)}>
                <p>管理工具最新版本为V 1.1.1，详情可前往<a className="a" href={`#${routes.systemVersion.path}`}>系统版本</a>查看 2022-01-18</p>
            </Dialog>
        </div>
    )
}
