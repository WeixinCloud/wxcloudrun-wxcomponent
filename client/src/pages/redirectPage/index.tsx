import {useEffect} from "react";
import {request} from "../../utils/axios";
import {getComponentInfoRequest} from "../../utils/apis";
import {routes} from "../../components/Console";

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
            console.log(window.location.search, location)
            window.location.href = resp.data.redirectUrl + window.location.hash.replaceAll(`#${routes.redirectPage.path}`, '')
        }
    }

    return (
        <div />
    )
}
