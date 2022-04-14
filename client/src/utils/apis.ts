export const HOST = import.meta.env.DEV ? '/api/wxcomponent' : '/wxcomponent'

type IRequestMsg = {
    url: string
    method: "get" | "post" | "delete" | "put"
}

export const getTicketRequest : IRequestMsg = {
    url: `${HOST}/admin/ticket`,
    method: "get"
}
export const getComponentTokenRequest: IRequestMsg = {
    url: `${HOST}/admin/component-access-token`,
    method: "get"
}
export const getAuthAccessTokenRequest: IRequestMsg = {
    url: `${HOST}/admin/authorizer-access-token`,
    method: "get"
}
export const getSecretRequest: IRequestMsg = {
    url: `${HOST}/admin/secret`,
    method: "get"
}
export const setSecretRequest: IRequestMsg = {
    url: `${HOST}/admin/secret`,
    method: "post"
}
export const getAuthorizedAccountRequest: IRequestMsg = {
    url: `${HOST}/admin/authorizer-list`,
    method: "get"
}
export const getNormalMessageRequest: IRequestMsg = {
    url: `${HOST}/admin/wx-biz-records`,
    method: "get"
}
export const getAuthMessageRequest: IRequestMsg = {
    url: `${HOST}/admin/wx-component-records`,
    method: "get"
}
export const getMessageConfigRequest: IRequestMsg = {
    url: `${HOST}/admin/callback-config`,
    method: "get"
}
export const changePasswordRequest: IRequestMsg = {
    url: `${HOST}/admin/userpwd`,
    method: "post"
}
export const changeUserNameRequest: IRequestMsg = {
    url: `${HOST}/admin/username`,
    method: "post"
}
export const updateComponentInfoRequest: IRequestMsg = {
    url: `${HOST}/admin/componentinfo`,
    method: "post"
}

export const getPreAuthCodeRequest: IRequestMsg = {
    url: `${HOST}/authpage/preauthcode`,
    method: "get"
}
export const getComponentInfoRequest: IRequestMsg = {
    url: `${HOST}/authpage/componentinfo`,
    method: "get"
}

export const loginRequest: IRequestMsg = {
    url: `${HOST}/auth`,
    method: "put"
}
export const refreshTokenRequest: IRequestMsg = {
    url: `${HOST}/admin/refresh-auth`,
    method: "get"
}

export const getProxyConfigRequest: IRequestMsg = {
    url: `${HOST}/admin/proxy`, // 获取代理配置
    method: "get"
}
export const updateProxyConfigRequest: IRequestMsg = {
    url: `${HOST}/admin/proxy`, // 更新代理配置
    method: "post"
}

export const getCallbackRuleRequest: IRequestMsg = {
    url: `${HOST}/admin/callback-proxy-rule-list`, // 获取消息推送转发配置列表
    method: "get"
}
export const updateCallbackRuleRequest: IRequestMsg = {
    url: `${HOST}/admin/callback-proxy-rule`, // 更新消息推送转发配置
    method: "post"
}
export const addCallbackRuleRequest: IRequestMsg = {
    url: `${HOST}/admin/callback-proxy-rule`, // 更新消息推送转发配置
    method: "put"
}
export const deleteCallbackRuleRequest: IRequestMsg = {
    url: `${HOST}/admin/callback-proxy-rule`, // 删除消息推送转发配置
    method: "delete"
}
export const testCallbackRuleRequest: IRequestMsg = {
    url: `${HOST}/admin/callback-test`, // 测试消息推送转发配置
    method: "post"
}

/* 代开发小程序相关 */
export const getDevMiniProgramListRequest: IRequestMsg = {
    url: `${HOST}/admin/dev-weapp-list`, // 获取代开发小程序列表
    method: "get"
}
export const getQrcodeRequest: IRequestMsg = {
    url: `${HOST}/admin/qrcode`, // 获取小程序码
    method: "get"
}
export const getDevVersionRequest: IRequestMsg = {
    url: `${HOST}/admin/dev-versions`, // 获取代开发小程序版本详情
    method: "get"
}
export const submitAuditRequest: IRequestMsg = {
    url: `${HOST}/admin/submit-audit`, // 提交审核
    method: "post"
}
export const revokeAuditRequest: IRequestMsg = {
    url: `${HOST}/admin/revoke-audit`, // 撤回审核
    method: "post"
}
export const speedUpAuditRequest: IRequestMsg = {
    url: `${HOST}/admin/speed-up-audit`, // 加急审核
    method: "post"
}
export const commitCodeRequest: IRequestMsg = {
    url: `${HOST}/admin/commit-code`, // 提交代码
    method: "post"
}
export const releaseCodeRequest: IRequestMsg = {
    url: `${HOST}/admin/release-code`, // 发布代码
    method: "post"
}
export const getTemplateListRequest: IRequestMsg = {
    url: `${HOST}/admin/template-list`, // 拉取模板列表
    method: "get"
}
export const rollbackReleaseRequest: IRequestMsg = {
    url: `${HOST}/admin/rollback-release-version`, // 拉取模板列表
    method: "post"
}
export const changeServiceStatusRequest: IRequestMsg = {
    url: `${HOST}/admin/change-visit-status`, // 修改服务状态
    method: "post"
}
export const getCategoryRequest: IRequestMsg = {
    url: `${HOST}/admin/category`, // 查询类目
    method: "get"
}
export const uploadFileRequest: IRequestMsg = {
    url: `${HOST}/admin/upload-media`, // 修改服务状态
    method: "post"
}
