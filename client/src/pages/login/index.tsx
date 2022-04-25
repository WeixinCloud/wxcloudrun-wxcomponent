import styles from './index.module.less'
import {Button, Dialog, Input, MessagePlugin} from 'tdesign-react'
import {useState} from "react";
import {loginRequest} from '../../utils/apis'
import {useNavigate} from "react-router-dom";
import moment from "moment";
import md5 from 'js-md5'
import {request} from "../../utils/axios";
import {routes} from "../../config/route";

export default function Login() {

    const [username, setUsername] = useState<string | number>('')
    const [password, setPassword] = useState<string | number>('')
    const [showForgetModal, setShowForgetModal] = useState<boolean>(false)

    const nav = useNavigate()

    const handleLogin = async () => {
        const resp = await request({
            request: loginRequest,
            data: {
                username,
                password: md5(String(password))
            },
            noNeedCheckLogin: true
        })
        if (resp.code === 0) {
            localStorage.setItem('token', resp.data.jwt)
            localStorage.setItem('expiresTime', String(moment().add(12, 'hours').valueOf()))
            localStorage.setItem('username', String(username))
            nav(routes.home.path)
            await MessagePlugin.success('登录成功', 2000)
        }
    }

    return (
        <div className={styles.login}>
            <div className={styles.login_modal}>
                <p style={{ marginBottom: '20px' }}>登录</p>
                <Input value={username} onChange={(val) => setUsername(val)} placeholder="请输入帐号" style={{ marginBottom: '15px' }} />
                <Input style={{ marginBottom: '20px' }} value={password} onChange={(val) => setPassword(val)} placeholder="请输入密码" type="password" />
                <a className="a" onClick={() => setShowForgetModal(true)}>忘记帐号或密码</a>
                <div style={{ textAlign: 'center' }}>
                    <Button style={{ marginTop: '20px', width: '100px' }} disabled={!Boolean(username && password)} onClick={handleLogin}>登录</Button>
                </div>
            </div>

            <Dialog header="忘记帐号或密码" visible={showForgetModal} onClose={() => setShowForgetModal(false)} confirmBtn={false}>
                <p className="desc">如忘记初始帐号和密码，可返回"微信开放平台-站内信"进行查看。</p>
                <p className="desc">若已重置过密码，需使用重置后的密码登录；如已忘记重置后的密码则无法再登录，需牢记重置后的密码。</p>
                <p className="desc">如有其他疑问，可扫码加入官方技术支持企业微信群反馈。</p>
            </Dialog>
        </div>
    )
}
