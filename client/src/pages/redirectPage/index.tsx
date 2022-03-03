import {useEffect} from "react";
import {request} from "../../utils/axios";
import {getComponentInfoRequest} from "../../utils/apis";

export default function RedirectPage() {

    useEffect(() => {
        jumpRealPage()
    }, [])

    const jumpRealPage = async () => {
        const resp = await request({
            request: getComponentInfoRequest,
            noNeedCheckLogin: true
        })
        if (resp.code === 0) {
            let redirectUrl = resp.data.redirectUrl
            if (redirectUrl.includes('#')) {
                redirectUrl = redirectUrl.replaceAll('changeWxSymbol', '#')
            }
            window.location.href = redirectUrl + window.location.search
        }
    }

    return (
        <div />
    )
}
