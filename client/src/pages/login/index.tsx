import styles from './index.module.less'
import { Button, Input, MessagePlugin } from 'tdesign-react'
import {useState} from "react";
import {loginUrl} from '../../utils/apis'
import {useNavigate} from "react-router-dom";
import moment from "moment";
import md5 from 'js-md5'
import {put} from "../../utils/axios";
import {routes} from "../../components/Console";

export default function Login() {

    const [username, setUsername] = useState<string | number>('')
    const [password, setPassword] = useState<string | number>('')

    const nav = useNavigate()

    const handleLogin = async () => {
        const resp = await put({
            url: loginUrl,
            data: {
                username,
                password: md5(String(password))
            },
            notNeedCheckLogin: true
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
                <p style={{ margin: 0 }}>登录</p>
                <p className="desc" style={{ marginBottom: '20px' }}>如忘记初始帐号和密码，可返回"微信开放平台-第三方平台-详情-云服务"进行查看</p>
                <Input value={username} onChange={(val) => setUsername(val)} placeholder="请输入帐号" style={{ marginBottom: '15px' }} />
                <Input value={password} onChange={(val) => setPassword(val)} placeholder="请输入密码" type="password" />
                <div style={{ textAlign: 'center' }}>
                    <Button style={{ marginTop: '20px', width: '100px' }} disabled={!Boolean(username && password)} onClick={handleLogin}>登录</Button>
                </div>
            </div>
        </div>
    )
}
