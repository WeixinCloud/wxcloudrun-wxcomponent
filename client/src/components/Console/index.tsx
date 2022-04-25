import React, {useEffect, useMemo, useState} from 'react'
import styles from './index.module.less'
import Menu from '../Menu'
import {Outlet, useNavigate, useLocation} from "react-router-dom";
import * as Icon from 'tdesign-icons-react'
import {Dropdown, Dialog} from 'tdesign-react';
import {checkLogin, logout} from "../../utils/login";
import {menuList} from "../../config/menu";
import {IMenuItem, IMenuList} from "../../commonType";
import {routes} from "../../config/route";

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
                    return <div className="normal_flex" style={{ alignItems: 'center' }}>
                        <Icon.ArrowLeftIcon style={{ marginRight: '12px' }} className="a" size="26px" onClick={() => window.history.back()} />
                        {/*<p style={{ lineHeight: '28px' }} className={styles.detail_header_title}><a href={`#${menu.showPath}`} className={`${styles.detail_header_title} a`}>{father?.label}</a> / {menu.label}</p>*/}
                        <p style={{ lineHeight: '28px' }} className={styles.detail_header_title}>{father?.label} / {menu.label}</p>
                    </div>
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
            <div style={{width: '232px', backgroundColor: '#f5f6f7'}} />
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
                <p>管理工具最新版本为V 2.2.0，详情可前往<a className="a" href={`#${routes.systemVersion.path}`}>系统版本</a>查看</p>
            </Dialog>
        </div>
    )
}
