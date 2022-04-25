import styles from './index.module.less'
import {routes} from "../../config/route";
import {copyMessage} from "../../utils/common";
import {request} from "../../utils/axios";
import {getComponentInfoRequest, updateComponentInfoRequest} from "../../utils/apis";
import {useEffect, useState} from "react";
import {Dialog, MessagePlugin, PopConfirm, Alert, Textarea} from 'tdesign-react'

let componentInfo = {}

const tipMessage = <div>
    <p style={{ margin: 0 }}>服务商需要获得商家授权后方可代商家开发、运营、管理商家公众号和小程序，因此需要生成授权链接，引导商家完成授权。</p>
    <p style={{ margin: '10px 0' }}>复制链接后，可将链接分享给商家，也可以复制授权链接到企业官网，引导用户授权。</p>
    <p style={{ margin: 0 }}>注意事项：如该第三方平台帐号尚未审核通过，则需将待授权的公众号或小程序加入“第三方平台-开发资料-授权测试公众号/小程序列表”后方可完成授权。</p>
</div>

export default function AuthPageManage() {

    const [redirectUrl, setRedirectUrl] = useState('')
    const [redirectUrlInput, setRedirectUrlInput] = useState('')
    const [showRedirectModal, setShowRedirectModal] = useState(false)

    useEffect(() => {
        getRedirectUrl()
    }, [])

    const getRedirectUrl = async () => {
        const resp = await request({
            request: getComponentInfoRequest,
        })
        if (resp.code === 0) {
            componentInfo = resp.data
            setRedirectUrl(resp.data.redirectUrl)
        }
    }

    const openRedirectModal = () => {
        setRedirectUrlInput(redirectUrl)
        setShowRedirectModal(true)
    }

    const updateRedirectUrl = async (isDel = false) => {
        if (!isDel && (!redirectUrlInput || !redirectUrlInput.startsWith('http'))) return MessagePlugin.info('请填入合法 url')
        const resp = await request({
            request: updateComponentInfoRequest,
            data: {
                ...componentInfo,
                redirectUrl: isDel ? '' : redirectUrlInput
            }
        })
        if (resp.code === 0) {
            setShowRedirectModal(false)
            MessagePlugin.success('修改成功')
            getRedirectUrl()
        }
    }

    return (
        <div>
            <Alert theme="info" icon={<span />} maxLine={0} message={tipMessage} style={{ marginBottom: '10px' }} />
            <p className="text">授权回调页配置</p>
            <div style={{ margin: '20px 0' }}>
                <div className="normal_flex" style={{ marginTop: '10px' }}>
                    <div className="blue_circle" />
                    <p className="desc" style={{ margin: '0' }}>商家授权成功后默认停留再授权成功页，如需在商家授权完成后自动跳转至回调页面，可通过下方配置添加回调 uri</p>
                </div>
                <div className="normal_flex" style={{ marginTop: '10px' }}>
                    <div className="blue_circle" />
                    <p className="desc" style={{ margin: '0' }}>支持配置由开发者自定义开发的uri，无域名前缀限制；授权后会自动重定向至该uri，但授权链接redirect_uri仍为微管家域名的uri，可在下方一键复制授权链接。</p>
                </div>
                <div className="normal_flex" style={{ marginTop: '10px' }}>
                    <div className="blue_circle" />
                    <p className="desc" style={{ margin: '0' }}>支持配置基于微管家进行二次开发的url，需与该微管家域名前缀一致。</p>
                </div>
                <div className="normal_flex" style={{ marginTop: '10px' }}>
                    <div className="blue_circle" />
                    <p className="desc" style={{ margin: '0' }}>当前微管家域名前缀为：{window.location.origin}。</p>
                </div>
            </div>
            <div className={styles.line} />
            <div className="normal_flex">
                <p style={{ width: '100px' }}>授权回调uri：</p>
                <p style={{ minWidth: '480px', textAlign: 'center', marginRight: '20px' }}>{redirectUrl}</p>
                <a style={{marginRight: '20px'}} className="a" onClick={openRedirectModal}>{redirectUrl ? '编辑' : '开启'}</a>
                {
                    redirectUrl &&
                    <PopConfirm onConfirm={() => updateRedirectUrl(true)} content="删除后，商家授权完成后将不在自动跳转至回调页面">
                        <a className="a">删除</a>
                    </PopConfirm>
                }
            </div>

            <div style={{ marginTop: '20px' }} />

            <p className="text">授权链接生成器介绍</p>
            <div className="normal_flex">
                <p className={styles.column}>授权链接</p>
                <p className={styles.column1}>使用方式</p>
                <p>操作</p>
            </div>
            <div className={styles.line} />
            <div className="normal_flex">
                <p className={styles.column}>PC 版授权链接</p>
                <p className={styles.column1}>在电脑浏览器里打开后，使用微信扫码</p>
                <a style={{marginRight: '20px'}} className="a"
                   onClick={() => copyMessage(`${window.location.origin}/#${routes.authorize.path}`)}>复制链接</a>
            </div>
            <div className="normal_flex">
                <p className={styles.column}>H5 版授权链接</p>
                <p className={styles.column1}>在手机微信里直接访问授权链接</p>
                <a style={{marginRight: '20px'}} className="a"
                     onClick={() => copyMessage(`${window.location.origin}/#${routes.authorizeH5.path}`)}>复制链接</a>
            </div>

            <Dialog visible={showRedirectModal} onClose={() => setShowRedirectModal(false)} onConfirm={() => updateRedirectUrl()}>
                <div className="normal_flex">
                    <p style={{ color: 'black' }}>授权回调url</p>
                    <Textarea onChange={val => setRedirectUrlInput(val as string)} value={redirectUrlInput} placeholder="请输入完整的url，包含协议前缀" />
                </div>
            </Dialog>
        </div>
    )
}
