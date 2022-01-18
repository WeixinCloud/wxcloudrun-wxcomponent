
export const HOST = import.meta.env.DEV ? '/api' : ''

export const getTicketUrl = `${HOST}/admin/ticket`
export const getComponentTokenUrl = `${HOST}/admin/component-access-token`
export const getAuthAccessTokenUrl = `${HOST}/admin/authorizer-access-token`
export const getSecretUrl = `${HOST}/admin/secret`
export const setSecretUrl = `${HOST}/admin/secret`
export const getAuthorizedAccountUrl = `${HOST}/admin/authorizer-list`
export const getNormalMessageUrl = `${HOST}/admin/wx-biz-records`
export const getAuthMessageUrl = `${HOST}/admin/wx-component-records`
export const getMessageConfigUrl = `${HOST}/admin/callback-config`
export const changePasswordUrl = `${HOST}/admin/userpwd`
export const changeUserNameUrl = `${HOST}/admin/username`

export const getPreAuthCodeUrl = `${HOST}/authpage/preauthcode`
export const getComponentInfoUrl = `${HOST}/authpage/componentinfo`

export const loginUrl = `${HOST}/auth`
export const refreshTokenUrl = `${HOST}/admin/refresh-auth`
