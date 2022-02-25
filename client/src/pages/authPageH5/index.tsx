import {useEffect} from "react";
import {request} from "../../utils/axios";
import {getComponentInfoRequest, getPreAuthCodeRequest} from "../../utils/apis";

export default function AuthPageH5() {

    useEffect(() => {
        jumpAuthPage()
    }, [])

    const jumpAuthPage = async () => {
        const resp = await request({
            request: getComponentInfoRequest,
            noNeedCheckLogin: true
        })
        if (resp.code === 0) {
            const resp1 = await request({
                request: getPreAuthCodeRequest,
                noNeedCheckLogin: true
            })
            if (resp1.code === 0) {
                window.location.href = `https://open.weixin.qq.com/wxaopen/safe/bindcomponent?component_appid=${resp.data.appid}&pre_auth_code=${resp1.data.preAuthCode}&auth_type=3#wechat_redirect`
            }
        }
    }

    return (
        <div />
    )
}
