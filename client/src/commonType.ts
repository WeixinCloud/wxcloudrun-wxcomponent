import React from "react";

export type IRoute = Record<string, {
    label: string // 标题
    path: string // 路径
    showPath?: string // 实际展示的对应menu的path，在menu中可能是hide的然后展示其他menu active状态
    element: React.ReactNode // 路由对应的页面组件
    dontNeedMenu?: boolean // 不需要左侧menu
}>

export type IMenuItem = {
    label: string
    path: string
    item?: IMenuItem[]
    showPath?: string
    hideItem?: IMenuItem[]
}

export type IMenuList = {
    label: string
    icon: JSX.Element
    path?: string
    item?: IMenuItem[]
    hideItem?: IMenuItem[]
}[]
