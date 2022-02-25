import {Dialog, Input, PopConfirm, MessagePlugin} from 'tdesign-react'
import {useEffect, useState} from "react";
import {request} from "../../utils/axios";
import {getProxyConfigRequest, updateProxyConfigRequest} from "../../utils/apis";
import styles from "./index.module.less";

export default function ProxyConfig() {

    const [open, setOpen] = useState(false)
    const [port, setPort] = useState(0)
    const [showPortModal, setShowPortModal] = useState(false)
    const [portInput, setPortInput] = useState(0)

    useEffect(() => {
        getProxyConfig()
    }, [])

    const getProxyConfig = async () => {
        const resp = await request({
            request: getProxyConfigRequest
        })
        if (resp.code === 0) {
            setPort(resp.data.port)
            setOpen(resp.data.open)
            setPortInput(resp.data.port)
        }
    }

    const handleClosePortModal = () => {
        setShowPortModal(false)
        setPortInput(port)
    }

    const changeProxyConfig = async (changeOpen = false) => {
        const resp = await request({
            request: updateProxyConfigRequest,
            data: {
                open: changeOpen ? !open : open,
                port: portInput
            }
        })
        if (resp.code === 0) {
            MessagePlugin.success('proxy 配置修改成功')
            setPort(portInput)
            setOpen(changeOpen ? !open : open)
            handleClosePortModal()
        }
    }

    return (
        <div>
            <p className="text">proxy 介绍</p>
            <div className="normal_flex">
                <div className="blue_circle" />
                <p className="desc"
                   style={{margin: 0}}>proxy 开启后可将外部请求透传转发至内部业务服务，从而实现将微管家与其他业务系统对接。</p>
            </div>
            <div className="normal_flex">
                <div className="blue_circle" />
                <p className="desc">默认转发至 8082 端口，用户可进行修改。</p>
            </div>
            <div className={styles.line} />
            <div className="normal_flex">
                <p style={{width: '100px'}}>操作</p>
                <div className="normal_flex">
                    <p style={{marginRight: '20px'}} className="desc">目前状态：{open ? '开启' : '关闭'}</p>
                    <PopConfirm onConfirm={() => changeProxyConfig(true)} content={open ? '关闭后，微管家将无法再将外部请求转发至后端服务' : '开启后，微管家可以将外部请求转发至后端服务'}>
                        <a className="a">{open ? '关闭' : '开启'}</a>
                    </PopConfirm>
                </div>
            </div>
            <div className="normal_flex">
                <p style={{width: '100px'}}>转发端口</p>
                <p style={{marginRight: '20px'}} className="desc">{port}</p>
                <a className="a" onClick={() => setShowPortModal(true)}>编辑</a>
            </div>

            <Dialog visible={showPortModal} onClose={handleClosePortModal} header="配置端口" onConfirm={() => changeProxyConfig()}>
                <div className="normal_flex">
                    <p style={{ width: '100px' }}>端口</p>
                    <Input type="number" placeholder="请输入需要配置的端口号" value={portInput} onChange={val => setPortInput(+val)} />
                </div>
            </Dialog>

        </div>
    )
}
