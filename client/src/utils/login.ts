import moment from 'moment'
import {request} from "./axios";
import { refreshTokenRequest } from './apis'

let nav: (to: string) => void

// 这个 hook 只能在函数组件里使用，在console初始化的时候绑定过来一份
export const initNav = (navigate: (to: string) => void) => {
    nav = navigate
}

export const checkLogin = () => {
    const token = localStorage.getItem('token')
    const expiresTime = localStorage.getItem('expiresTime')
    if (!token || !expiresTime) {
        if (nav) {
            nav('/login')
        } else {
            window.location.hash = '#/login'
        }
        return false
    }
    const now = moment().valueOf()
    const timeDiff = Number(expiresTime) - now
    if (timeDiff) {
        // 少于3小时刷新token 3 * 60 * 60 * 1000
        if (timeDiff < (3 * 60 * 60 * 1000)) {
            refreshToken()
        }
        return true
    } else {
        localStorage.removeItem('token')
        localStorage.removeItem('expiresTime')
        if (nav) {
            nav('/login')
        } else {
            window.location.hash = '#/login'
        }
        return false
    }
}

export const refreshToken = async () => {
    const resp = await request({
        request: refreshTokenRequest,
        noNeedCheckLogin: true,
    }, (code) => {
        if (code === 1004) {
            localStorage.removeItem('token')
            localStorage.removeItem('expiresTime')
            if (nav) {
                nav('/login')
            } else {
                window.location.hash = '#/login'
            }
        }
    })
    if (resp.code === 0) {
        localStorage.setItem('token', resp.data.jwt)
        localStorage.setItem('expiresTime', String(moment().add(7, 'hours').valueOf()))
    }
}

export const logout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('expiresTime')
    localStorage.removeItem('username')
    if (nav) {
        nav('/login')
    } else {
        window.location.hash = '#/login'
    }
}
