import {useEffect, useState} from 'react'
import { useNavigate, useLocation } from 'react-router-dom'
import { FullscreenExitIcon, FullscreenIcon } from 'tdesign-icons-react'
import { Menu } from 'tdesign-react/'
import {routes} from "../../config/route";

const {SubMenu, MenuItem} = Menu;

interface IProps {
    menuList: {
        label?: string | JSX.Element
        icon: JSX.Element
        path?: string
        item?: {
            label: string | JSX.Element
            path: string
            showPath?: string
        }[]
    }[]
}

export default function MyMenu(props: IProps) {

    const {menuList} = props

    const [activePath, setActivePath] = useState<string | number>('')
    const [expandsMenu, setExpandsMenu] = useState<Array<string | number>>([])
    const [collapsed, setCollapsed] = useState(false)
    const navigate = useNavigate()
    const location = useLocation()

    useEffect(() => {
        if (location.pathname === activePath) return
        // 没想到太好的解法只能先这样写了
        const keys = Object.keys(routes)
        for (let i = 0; i < keys.length; i++) {
            // @ts-ignore
            if (routes[keys[i]].path === location.pathname) {
                // @ts-ignore
                setActivePath(routes[keys[i]].showPath ?? location.pathname)
                return;
            }
        }
    }, [location.pathname])

    const changePath = (path: string | number) => {
        path = String(path)
        setActivePath(path)
        navigate(path)
    }

    return (
        <Menu
            theme="dark"
            collapsed={collapsed}
            value={activePath}
            expandMutex={false}
            expanded={expandsMenu}
            onExpand={(values) => setExpandsMenu(values)}
            onChange={changePath}
            style={{height: '100%'}}
            operations={collapsed ? <FullscreenIcon className="t-menu__operations-icon" onClick={() => setCollapsed(!collapsed)} /> : <FullscreenExitIcon className="t-menu__operations-icon" onClick={() => setCollapsed(!collapsed)} />}
            logo={<h3 style={{margin: '0 auto', color: 'white'}}>{collapsed ? '' : '服务商微管家'}</h3>}
        >
            {
                (menuList || []).map((i, index) => {
                    if (i.item) {
                        return (
                            <SubMenu key={`menu_father_${index}`} value={`menu_father_${index}`} title={i.label} icon={i.icon}>
                                {
                                    (i.item || []).map(item => {
                                        return (
                                            <MenuItem key={`menu_item_${item.path}`} value={item.path}>
                                                {item.label}
                                            </MenuItem>
                                        )
                                    })
                                }
                            </SubMenu>
                        )
                    } else {
                        return (
                            <MenuItem value={i.path} key={`menu_${i.path}`} icon={i.icon}>
                                {i.label}
                            </MenuItem>
                        )
                    }
                })
            }
        </Menu>
    )
}
