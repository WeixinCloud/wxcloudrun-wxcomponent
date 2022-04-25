import * as _apis from "../../utils/apis";

type IRequestMsg = {
    url: string
    method: "get" | "post" | "delete" | "put"
}

export const HOST = import.meta.env.DEV ? '/api/wxcomponent' : '/wxcomponent'

// 项目所有接口
export const apis = _apis

export const demoRequest: IRequestMsg =  {
    url: `${HOST}/test/demo`,
    method: 'post'
}
