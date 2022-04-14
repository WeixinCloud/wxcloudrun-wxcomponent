import axios, {AxiosRequestConfig} from 'axios'
import {checkLogin, logout} from "./login";
import { MessagePlugin } from "tdesign-react";
import {objToQueryString} from "./common";

type IAxiosParams = {
    url: string
    data?: Record<string, any>
    config?: AxiosRequestConfig
    noNeedCheckLogin?: boolean
}

type IAxiosParams1 = {
    request: {
        method: 'get' | 'post' | 'put' | 'delete'
        url: string
    }
    data?: Record<string, any>
    config?: AxiosRequestConfig
    noNeedCheckLogin?: boolean
}

type IAxiosResp = {
    code: number
    errorMsg: string
    data: any
}

const noLoginError: {
    code: -1
    errorMsg: string
} = {
    code: -1,
    errorMsg: 'should login'
}

const errorMsg: Record<number, string> = {
    1000: '访问未授权',
    1001: '参数格式错误',
    1003: '登录超时',
    1004: 'Token 错误',
    1005: '用户更新错误',
    1006: '登录失败',
    1007: 'Ticket 为空',
    41004: '尚未配置 Secret，无法生成 token，需前往"系统管理-Secret与密码管理"完成配置后再使用'
}

type IErrHandle = (code: number, data?: any) => void;

const mustErrHandle = (data: {
    code: number
    errorMsg: string
    data: string
}) => {
    if (data.code === 1003 || data.code === 1004) {
        logout()
        MessagePlugin.error(errorMsg[data.code], 4000)
        return true
    }
    return false
}

const defaultErrHandle = (data: {
    code: number
    errorMsg: string
    data: string
}) => {
    switch (data.code) {
        case 1002: {
            return MessagePlugin.error(`${data.data}`, 4000)
        }
        default: {
            return MessagePlugin.error(errorMsg[data.code] || `系统错误，请稍后重试 code：${data.code}`, 4000)
        }
    }
}

const dataFormat = (data: Record<string, any>) => {
    Object.keys(data).forEach(i => {
        if (typeof data[i] === 'number' && String(data[i] ).length === 10) {
            // 防止精度问题
            data[i]  = Number(String(data[i]) + '000')
        }
        if (typeof data[i] === 'object') {
            // typeof null === 'object'
            if (!data[i]) return
            if (Array.isArray(data[i])) {
                arrayFormat(data[i])
            } else {
                dataFormat(data[i])
            }
        }
    })
    return data
}

const arrayFormat = (data: any[]) => {
    data.forEach((i, index) => {
        if (typeof i === 'number' && String(i).length === 10) {
            // 防止精度问题
            data[index] = Number(String(i) + '000')
        }
        if (typeof i === 'object') {
            if (Array.isArray(i)) {
                arrayFormat(i)
            } else {
                dataFormat(i)
            }
        }
    })
    return data
}

export const get = async (params: IAxiosParams, errHandle?: IErrHandle) => {
    if (!params.noNeedCheckLogin) {
        if (!checkLogin()) {
            return noLoginError
        }
    }
    const data = (await axios.get(`${params.url}`, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token') || ''}`
        },
        ...params.config
    })).data as IAxiosResp
    if (data.code === 0) {
        return dataFormat(data)
    }
    if (!mustErrHandle(data)) {
        if (errHandle) {
            errHandle(data.code, data)

        } else {
            defaultErrHandle(data)
        }
    }
    return data
}

export const post = async (params: IAxiosParams, errHandle?: IErrHandle) => {
    if (!params.noNeedCheckLogin) {
        if (!checkLogin()) {
            return noLoginError
        }
    }
    const data = (await axios.post(`${params.url}`, params.data, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token') || ''}`
        },
        ...params.config
    })).data as IAxiosResp
    if (data.code === 0) {
        return dataFormat(data)
    }
    if (!mustErrHandle(data)) {
        if (errHandle) {
            errHandle(data.code, data)

        } else {
            defaultErrHandle(data)
        }
    }
    return data
}

export const put = async (params: IAxiosParams, errHandle?: IErrHandle) => {
    if (!params.noNeedCheckLogin) {
        if (!checkLogin()) {
            return noLoginError
        }
    }
    const data = (await axios.put(`${params.url}`, params.data, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token') || ''}`
        },
        ...params.config
    })).data as IAxiosResp
    if (data.code === 0) {
        return dataFormat(data)
    }
    if (!mustErrHandle(data)) {
        if (errHandle) {
            errHandle(data.code, data)

        } else {
            defaultErrHandle(data)
        }
    }
    return data
}

export const deleteRequest = async (params: IAxiosParams, errHandle?: IErrHandle) => {
    if (!params.noNeedCheckLogin) {
        if (!checkLogin()) {
            return noLoginError
        }
    }
    const data = (await axios.delete(`${params.url}`, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token') || ''}`
        },
        ...params.config
    })).data as IAxiosResp
    if (data.code === 0) {
        return dataFormat(data)
    }
    if (!mustErrHandle(data)) {
        if (errHandle) {
            errHandle(data.code, data)

        } else {
            defaultErrHandle(data)
        }
    }
    return data
}

export const request = async (params: IAxiosParams1, errHandle?: IErrHandle) => {
    if (!params.noNeedCheckLogin) {
        if (!checkLogin()) {
            return noLoginError
        }
    }
    switch (params.request.method) {
        case "get": {
            return get({
                url: `${params.request.url}?${params.data ? objToQueryString(params.data) : ''}`,
                noNeedCheckLogin: true,
                config: params.config
            }, errHandle)
        }
        case "post": {
            return post({
                url: params.request.url,
                noNeedCheckLogin: true,
                data: params.data,
                config: params.config
            }, errHandle)
        }
        case "put": {
            return put({
                url: params.request.url,
                noNeedCheckLogin: true,
                data: params.data,
                config: params.config
            }, errHandle)
        }
        case "delete": {
            return deleteRequest({
                url: `${params.request.url}?${params.data ? objToQueryString(params.data) : ''}`,
                noNeedCheckLogin: true,
                config: params.config
            }, errHandle)
        }
    }
}
