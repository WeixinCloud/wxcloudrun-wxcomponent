import {
    copyMessage as _copyMessage
} from '../../utils/common'
import {
    checkLogin as _checkLogin,
    logout as _logout,
    refreshToken as _refreshToken
} from '../../utils/login'

import { request as _request } from "../../utils/axios";

// 复制文字
export const copyMessage = _copyMessage

// 检查是否登录
export const checkLogin = _checkLogin

// 退出登录
export const logout = _logout

// 刷新登录token
export const refreshToken = _refreshToken

// 请求方法
export const request = _request
