import {useEffect} from "react";
import {get} from "../../utils/axios";
import {getComponentInfoUrl, getPreAuthCodeUrl} from "../../utils/apis";

export default function AuthPage() {

    useEffect(() => {
        jumpAuthPage()
    }, [])

    const jumpAuthPage = async () => {
        const resp = await get({
            url: getComponentInfoUrl,
            notNeedCheckLogin: true
        })
        if (resp.code === 0) {
            const resp1 = await get({
                url: getPreAuthCodeUrl,
                notNeedCheckLogin: true
            })
            if (resp1.code === 0) {
                window.location.href = `https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=${resp.data.appid}&pre_auth_code=${resp1.data.preAuthCode}&auth_type=3`
            }
        }
    }

    return (
        <div />
    )
}
