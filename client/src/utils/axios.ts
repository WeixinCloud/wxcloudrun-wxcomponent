import axios, {AxiosRequestConfig} from 'axios'
import { checkLogin } from "./login";
import { MessagePlugin } from "tdesign-react";

type IAxiosParams = {
    url: string
    data?: Record<string, any>
    config?: AxiosRequestConfig
    notNeedCheckLogin?: boolean
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

const defaultErrHandle = async (data: {
    code: number
    errorMsg: string
    data: string
}) => {
    switch (data.code) {
        case 1002: {
            return await MessagePlugin.error(`系统错误，请稍后重试 reason: ${data.errorMsg} - ${data.data}`, 2000)
        }
        default: {
            return await MessagePlugin.error(errorMsg[data.code] || `系统错误，请稍后重试 code：${data.code}`, 2000)
        }
    }
}

export const get = async (params: IAxiosParams, errHandle?: IErrHandle) => {
    if (!params.notNeedCheckLogin) {
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
        return data
    }
    if (errHandle) {
        errHandle(data.code, data)
    } else {
        defaultErrHandle(data)
    }
    return data
}

export const post = async (params: IAxiosParams, errHandle?: IErrHandle) => {
    if (!params.notNeedCheckLogin) {
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
        return data
    }
    if (errHandle) {
        errHandle(data.code, data)
    } else {
        defaultErrHandle(data)
    }
    return data
}

export const put = async (params: IAxiosParams, errHandle?: IErrHandle) => {
    if (!params.notNeedCheckLogin) {
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
        return data
    }
    if (errHandle) {
        errHandle(data.code, data)
    } else {
        defaultErrHandle(data)
    }
    return data
}
