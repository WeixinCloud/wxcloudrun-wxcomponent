import styles from './index.module.less'
import {MessagePlugin, Dialog, Input} from "tdesign-react";
import {request} from "../../utils/axios";
import {
    changePasswordRequest,
    changeUserNameRequest,
    getComponentInfoRequest,
    getSecretRequest,
    setSecretRequest
} from "../../utils/apis";
import {useEffect, useState} from "react";
import md5 from 'js-md5';

export default function PasswordManage() {

    const [componentAppId, setComponentAppId] = useState<string | number>('')
    const [secret, setSecret] = useState<string>('')
    const [username, setUsername] = useState<string | number>(localStorage.getItem('username') || '')

    const [showSecretModal, setShowSecretModal] = useState(false)
    const [showUsernameModal, setShowUsernameModal] = useState(false)
    const [showPasswordModal, setShowPasswordModal] = useState(false)
    const [secretInput, setSecretInput] = useState<string | number>('')
    const [usernameInput, setUsernameInput] = useState<string | number>('')
    const [oldPasswordInput, setOldPasswordInput] = useState<string | number>('')
    const [passwordInput, setPasswordInput] = useState<string | number>('')
    const [repeatPasswordInput, setRepeatPasswordInput] = useState<string | number>('')

    useEffect(() => {
        getSomeInfo()
    }, [])

    const getSomeInfo = async () => {
        const resp = await request({
            request: getComponentInfoRequest
        })
        if (resp.code === 0) {
            setComponentAppId(resp.data.appid)
        }

        const resp2 = await request({
            request: getSecretRequest
        })
        if (resp2.code === 0) {
            setSecret(resp2.data.secret)
        }
    }

    const handleChangeSecret = async () => {
        if (!secretInput) {
            return MessagePlugin.error('有信息未输入', 2000)
        }
        const resp = await request({
            request: setSecretRequest,
            data: {
                secret: secretInput
            }
        })
        if (resp.code === 0) {
            MessagePlugin.success('修改 secret 成功', 2000)
            setShowSecretModal(false)
            setSecret(secretInput as string)
            setSecretInput('')
        }
    }

    const handleChangeUsername = async () => {
        if (!usernameInput) {
            return MessagePlugin.error('有信息未输入', 2000)
        }
        const resp = await request({
            request: changeUserNameRequest,
            data: {
                username: usernameInput
            }
        })
        if (resp.code === 0) {
            MessagePlugin.success('修改帐号成功', 2000)
            localStorage.setItem('username', String(usernameInput))
            setShowUsernameModal(false)
            setUsername(usernameInput)
            setUsernameInput('')
        }
    }

    const handleChangePassword = async () => {
        if (!oldPasswordInput || !passwordInput || !repeatPasswordInput) {
            return MessagePlugin.error('有信息未输入', 2000)
        }
        if (passwordInput !== repeatPasswordInput) {
            return MessagePlugin.error('两次密码输入不一致', 2000)
        }
        if (passwordInput == oldPasswordInput) {
            return MessagePlugin.error('新旧密码相同', 2000)
        }
        if (!(/^[\w!@#$%^&*()+.]{6,10}$/.test(String(passwordInput)))) {
            return MessagePlugin.error('密码不符合要求', 2000)
        }
        const resp = await request({
            request: changePasswordRequest,
            data: {
                password: md5(String(passwordInput)),
                oldPassword: md5(String(oldPasswordInput)),
            }
        })
        if (resp.code === 0) {
            MessagePlugin.success('修改密码成功', 2000)
            setShowPasswordModal(false)
            setPasswordInput('')
            setOldPasswordInput('')
            setRepeatPasswordInput('')
        }
    }

    return (
        <div>
            <p className="text">设置 Secret</p>
            <p className="desc">第三方平台 Secret 补充完整后可使用该平台完整功能；且 Secret 在第三方平台重置后，需要及时修改，否则会影响功能使用</p>
            <div className={styles.line} />
            <div className="normal_flex">
                <p style={{width: '100px'}}>第三方 AppID</p>
                <p style={{marginRight: '20px'}} className="desc">{componentAppId}</p>
            </div>
            <div className="normal_flex">
                <p style={{width: '100px'}}>第三方 Secret</p>
                <p style={{marginRight: '20px'}} className="desc">{secret ? `${secret.slice(0,5)}***********${secret.slice(secret.length - 4, secret.length)}` : ''}</p>
                <a style={{marginRight: '20px'}} className="a" onClick={() => setShowSecretModal(true)}>修改</a>
            </div>

            <p style={{marginTop: '40px'}} className="text">修改帐号和密码</p>
            <p className="desc">修改帐号和密码后，使用新的帐号和密码才可登录该系统，请谨慎操作。此外，如需修改数据库密码可前往<a href="https://cloud.weixin.qq.com/clourun?utm_source=Third-party-Platform" target="_blank" className="a">微信云托管</a>进行操作</p>
            <div className={styles.line} />
            <div className="normal_flex">
                <p style={{width: '100px'}}>登录帐号</p>
                <p style={{marginRight: '20px'}} className="desc">{username}</p>
                <a style={{marginRight: '20px'}} className="a" onClick={() => setShowUsernameModal(true)}>修改帐号</a>
            </div>
            <div className="normal_flex">
                <p style={{width: '100px'}}>登录密码</p>
                <p style={{marginRight: '20px'}} className="desc">***************</p>
                <a style={{marginRight: '20px'}} className="a" onClick={() => setShowPasswordModal(true)}>修改密码</a>
            </div>

            <Dialog visible={showSecretModal} onClose={() => setShowSecretModal(false)} onConfirm={handleChangeSecret} header="修改 Secret">
                <div className="normal_flex">
                    <p className={styles.text}>Secret</p>
                    <Input className={styles.input} value={secretInput} onChange={setSecretInput} placeholder="请输入第三方帐号 Secret" />
                </div>
            </Dialog>

            <Dialog visible={showUsernameModal} onClose={() => setShowUsernameModal(false)} onConfirm={handleChangeUsername} header="修改登录帐号">
                <div className="normal_flex">
                    <p className={styles.text}>当前帐号</p>
                    <p className={styles.input}>{username}</p>
                </div>
                <div className="normal_flex">
                    <p className={styles.text}>新的帐号</p>
                    <Input className={styles.input} value={usernameInput} onChange={setUsernameInput} placeholder="请输入新帐号" />
                </div>
            </Dialog>

            <Dialog visible={showPasswordModal} onClose={() => setShowPasswordModal(false)} onConfirm={handleChangePassword} header="修改登录密码">
                <div className="normal_flex">
                    <p className={styles.text}>原密码</p>
                    <Input className={styles.input} value={oldPasswordInput} onChange={setOldPasswordInput} placeholder="请输入原登录密码" />
                </div>
                <div className="normal_flex">
                    <p className={styles.text}>输入密码</p>
                    <Input className={styles.input} value={passwordInput} onChange={setPasswordInput} placeholder="请输入密码" />
                    密码要求：长度为6-10，由大小写字母、数字、符号（!@#$%^&*()_+.）组成
                </div>
                <div className="normal_flex">
                    <p className={styles.text}>确认密码</p>
                    <Input className={styles.input} value={repeatPasswordInput} onChange={setRepeatPasswordInput} placeholder="请确认密码" />
                </div>
            </Dialog>

        </div>
    )
}
