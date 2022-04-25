import styles from "./index.module.less";
import {
    Dropdown,
    Button,
    Tag,
    Dialog,
    Form,
    Select,
    Textarea,
    Input,
    Loading,
    Alert,
    MessagePlugin
} from 'tdesign-react';
import {Icon} from 'tdesign-icons-react';
import {useEffect, useMemo, useRef, useState} from "react";
import {truncate} from 'lodash'
import {request} from "../../../utils/axios";
import {
    commitCodeRequest,
    getDevVersionRequest, getTemplateListRequest,
    releaseCodeRequest,
    revokeAuditRequest, rollbackReleaseRequest,
    speedUpAuditRequest,
} from "../../../utils/apis";
import {useSearchParams} from "react-router-dom";
import moment from "moment";
import {routes} from "../../../config/route";

const {FormItem} = Form;
const {Option} = Select

type IVersionData = {
    auditInfo?: {
        auditId: number
        status: number
        reason: string
        screenShot: string
        userVersion: string
        userDesc: string
        submitAuditTime: number
    }
    releaseInfo?: {
        releaseTime: number
        releaseVersion: string
        releaseDesc: string
        releaseQrCode: string // base64
    }
    expInfo?: {
        expTime: number
        expVersion: string
        expDesc: string
        expQrCode: string // base64
    }
}

type ITemplateList = {
    templateId: number
    userVersion: string
    userDesc: string
}[]

// 线上版选项
const onlineOptions = [{
    content: '版本回退',
    value: 1,
}]

export default function MiniProgramVersion() {

    const [versionData, setVersionData] = useState<IVersionData>({
        auditInfo: undefined,
        releaseInfo: undefined,
        expInfo: undefined
    })
    const [visibleSubmitModal, setVisibleSubmitModal] = useState(false)
    const [loading, setLoading] = useState(true)
    const [templateList, setTemplateList] = useState<ITemplateList>([])
    const [searchParams] = useSearchParams();

    const formRef = useRef() as any

    const appId = useMemo(() => {
        return searchParams.get('appId')
    }, [searchParams])

    // 审核版选项
    const auditOptions = useMemo(() => {
        if (versionData.auditInfo?.status === 0) {
            return [{
                content: '提交发布',
                value: 4,
            }]
        }
        if (versionData.auditInfo?.status === 1 || versionData.auditInfo?.status === 3) {
            return [{
                content: '提交审核',
                value: 1,
            }]
        }
        if (versionData.auditInfo?.status === 2 || versionData.auditInfo?.status === 4) {
            return [{
                content: '加急审核',
                value: 2,
            }, {
                content: '撤回审核',
                value: 3,
            }]
        }
        return []
    }, [versionData])

    // 体验版选项
    const experienceOptions = useMemo(() => {
        const arr = [{
            content: '重新提交代码',
            value: 1,
        }]
        if (versionData.auditInfo?.status !== 2) {
            arr.push( {
                content: '提交审核',
                value: 2,
            })
        }
        return arr
    }, [versionData])

    useEffect(() => {
        if (!searchParams.get('appId')) {
            return
        }
        getVersion()
        getTemplateList()
    }, [])

    const onClickOnlineOptions = async () => {
        const resp = await request({
            request: {
                url: `${rollbackReleaseRequest.url}?appid=${appId}`,
                method: rollbackReleaseRequest.method
            }
        })
        if (resp.code === 0) {
            MessagePlugin.success('版本回退成功')
            getVersion()
        }
    }

    // 这个组件这里的interface定义的不好
    const onClickAuditOptions = async (option: any) => {
        const { value } = option
        switch (value) {
            case 1: {
                // 提交审核
                window.location.href = `#${routes.submitAudit.path}?appId=${appId}`
                break
            }
            // 加急审核
            case 2: {
                const resp = await request({
                    request: {
                        url: `${speedUpAuditRequest.url}?appid=${appId}`,
                        method: speedUpAuditRequest.method
                    },
                })
                if (resp.code === 0) {
                    MessagePlugin.success('加急审核成功')
                    getVersion()
                }
                break
            }
            // 撤回审核
            case 3: {
                const resp = await request({
                    request: {
                        url: `${revokeAuditRequest.url}?appid=${appId}`,
                        method: revokeAuditRequest.method
                    },
                })
                if (resp.code === 0) {
                    MessagePlugin.success('撤回审核成功')
                    getVersion()
                }
                break
            }
            // 提交发布
            case 4: {
                const resp = await request({
                    request: {
                        url: `${releaseCodeRequest.url}?appid=${appId}`,
                        method: releaseCodeRequest.method
                    },
                })
                if (resp.code === 0) {
                    MessagePlugin.success('提交发布成功')
                    getVersion()
                }
                break
            }
        }
    }

    // 这个组件这里的interface定义的不好
    const onClickExpOptions = (option: any) => {
        const { value } = option
        switch (value) {
            case 1: {
                // 重新提交代码
                setVisibleSubmitModal(true)
                break
            }
            case 2: {
                // 提交审核
                window.location.href = `#${routes.submitAudit.path}?appId=${appId}`
                break
            }
        }
    }

    const getVersion = async () => {
        setLoading(true)
        const resp = await request({
            request: getDevVersionRequest,
            data: {
                appid: appId,
            }
        })
        if (resp.code === 0) {
            setVersionData(resp.data)
        }
        setLoading(false)
    }

    const getTemplateList = async () => {
        const resp = await request({
            request: getTemplateListRequest,
            data: {
                templateType: 0
            }
        })
        if (resp.code === 0) {
            setTemplateList(resp.data.templateList)
        }
    }

    const closeSubmitModal = () => {
        formRef.current.reset()
        setVisibleSubmitModal(false)
    }

    const submitCode = async (e: { validateResult: any }) => {
        if (e.validateResult !== true) {
            return
        }
        const { templateId, extJson, userVersion, userDesc } = formRef.current.getAllFieldsValue()
        const resp = await request({
            request: {
                url: `${commitCodeRequest.url}?appid=${appId}`,
                method: commitCodeRequest.method
            },
            data: {
                templateId: String(templateId),
                extJson,
                userVersion,
                userDesc
            }
        })
        if (resp.code === 0) {
            MessagePlugin.success('提交发布成功')
            getVersion()
            closeSubmitModal()
        }
    }

    return (
        <div>
            <p className="text">线上版本</p>
            <div className={styles.line} />
            {
                versionData.releaseInfo
                    ?
                    <div style={{display: 'flex', alignItems: 'flex-start', justifyContent: 'space-between'}}>
                        <div>
                            <div style={{display: 'flex'}}>
                                <div style={{
                                    display: 'flex',
                                    flexDirection: 'column',
                                    width: '60px',
                                    marginRight: '16px'
                                }}>
                                    <p className="desc">版本号</p>
                                    <p style={{fontSize: '16px'}}>{versionData.releaseInfo.releaseVersion}</p>
                                </div>
                                <div style={{display: 'flex', flexDirection: 'column'}}>
                                    <div className="normal_flex">
                                        <p className="desc" style={{width: '74px'}}>发布时间</p>
                                        <p>{moment(versionData.releaseInfo.releaseTime).format('YYYY-MM-DD HH:mm:ss')}</p>
                                    </div>
                                    <div className="normal_flex">
                                        <p className="desc" style={{width: '74px'}}>版本描述</p>
                                        <p>{versionData.releaseInfo.releaseDesc}</p>
                                    </div>
                                </div>
                            </div>
                            {
                                versionData.releaseInfo.releaseQrCode
                                &&
                                <div style={{width: '370px', textAlign: 'center'}}>
                                    <img style={{width: '200px', height: '200px'}}
                                         src={`data:image/png;base64,${versionData.releaseInfo.releaseQrCode}`}
                                         alt="" />
                                </div>
                            }
                        </div>
                        <Dropdown options={onlineOptions} onClick={onClickOnlineOptions}>
                            <Button style={{marginTop: '20px', paddingRight: '10px', marginRight: '40px'}}
                                    theme="primary">操作<Icon
                                style={{marginLeft: '3px'}} name="chevron-down" size="16" /></Button>
                        </Dropdown>
                    </div>
                    :
                    <Loading size="small" loading={loading} showOverlay>
                        <div className="desc" style={{textAlign: 'center', margin: '100px 0'}}>
                            尚未提交线上版本
                        </div>
                    </Loading>
            }

            <div className="normal_flex">
                <p className="text" style={{marginTop: '30px'}}>审核版本</p>
                {
                    versionData.auditInfo
                    &&
                    <div style={{margin: '30px 10px 18px 10px'}}>
                        {versionData.auditInfo.status === 0 && <Tag theme="success">审核通过</Tag>}
                        {versionData.auditInfo.status === 1 && <Tag theme="danger">审核不通过</Tag>}
                        {versionData.auditInfo.status === 2 && <Tag theme="warning">审核中</Tag>}
                        {versionData.auditInfo.status === 3 && <Tag theme="primary">已撤回</Tag>}
                        {versionData.auditInfo.status === 4 && <Tag theme="warning">审核延后</Tag>}
                    </div>
                }
            </div>
            <div className={styles.line} />
            {
                versionData.auditInfo
                    ?
                    <div style={{display: 'flex', alignItems: 'flex-start', justifyContent: 'space-between'}}>
                        <div>
                            <div style={{display: 'flex'}}>
                                <div style={{
                                    display: 'flex',
                                    flexDirection: 'column',
                                    width: '60px',
                                    marginRight: '16px'
                                }}>
                                    <p className="desc">版本号</p>
                                    <p style={{fontSize: '16px'}}>{versionData.auditInfo.userVersion}</p>
                                </div>
                                <div style={{display: 'flex', flexDirection: 'column'}}>
                                    <div className="normal_flex">
                                        <p className="desc" style={{width: '74px'}}>审核ID</p>
                                        <p>{versionData.auditInfo.auditId}</p>
                                    </div>
                                    <div className="normal_flex">
                                        <p className="desc" style={{width: '74px'}}>提交时间</p>
                                        <p>{moment(versionData.auditInfo.submitAuditTime).format('YYYY-MM-DD HH:mm:ss')}</p>
                                    </div>
                                    <div className="normal_flex">
                                        <p className="desc" style={{width: '74px'}}>版本描述</p>
                                        <p>{versionData.auditInfo.userDesc}</p>
                                    </div>
                                    {
                                        versionData.auditInfo.status === 1
                                        &&
                                        <Alert theme="error" icon={<span />} maxLine={0} message={`【驳回原因】：${versionData.auditInfo.reason}`} style={{ marginBottom: '10px' }} />
                                    }
                                </div>
                            </div>
                        </div>
                        <Dropdown options={auditOptions} onClick={onClickAuditOptions}>
                            <Button style={{marginTop: '20px', paddingRight: '10px', marginRight: '40px'}}
                                    theme="primary">操作<Icon
                                style={{marginLeft: '3px'}} name="chevron-down" size="16" /></Button>
                        </Dropdown>
                    </div>
                    :
                    <Loading size="small" loading={loading} showOverlay>
                        <div className="desc" style={{textAlign: 'center', margin: '100px 0'}}>
                            暂无提交审核的版本或者版本已发布上线
                        </div>
                    </Loading>
            }

            <p className="text" style={{marginTop: '30px'}}>体验版本</p>
            <div className={styles.line} />
            {
                versionData.expInfo
                    ?
                    <div style={{display: 'flex', alignItems: 'flex-start', justifyContent: 'space-between'}}>
                        <div>
                            <div style={{display: 'flex'}}>
                                <div style={{
                                    display: 'flex',
                                    flexDirection: 'column',
                                    width: '60px',
                                    marginRight: '16px'
                                }}>
                                    <p className="desc">版本号</p>
                                    <p style={{fontSize: '16px'}}>{versionData.expInfo.expVersion}</p>
                                </div>
                                <div style={{display: 'flex', flexDirection: 'column'}}>
                                    <div className="normal_flex">
                                        <p className="desc" style={{width: '74px'}}>提交时间</p>
                                        <p>{moment(versionData.expInfo.expTime).format('YYYY-MM-DD HH:mm:ss')}</p>
                                    </div>
                                    <div className="normal_flex">
                                        <p className="desc" style={{width: '74px'}}>版本描述</p>
                                        <p>{versionData.expInfo.expDesc}</p>
                                    </div>
                                </div>
                            </div>
                            <div style={{width: '370px', textAlign: 'center'}}>
                                <img style={{width: '200px', height: '200px'}}
                                     src={`data:image/png;base64,${versionData.expInfo.expQrCode}`}
                                     alt="" />
                            </div>
                        </div>
                        <Dropdown options={experienceOptions} onClick={onClickExpOptions}>
                            <Button style={{marginTop: '20px', paddingRight: '10px', marginRight: '40px'}}
                                    theme="primary">操作<Icon
                                style={{marginLeft: '3px'}} name="chevron-down" size="16" /></Button>
                        </Dropdown>
                    </div>
                    :
                    <Loading size="small" loading={loading} showOverlay>
                        <div style={{textAlign: 'center', margin: '100px 0'}}>
                            <p className="desc">尚未提交体验版</p>
                            <Button onClick={() => setVisibleSubmitModal(true)}>提交代码</Button>
                        </div>
                    </Loading>
            }

            <Dialog visible={visibleSubmitModal} onClose={closeSubmitModal} confirmBtn={null} width={700}
                    cancelBtn={null} header="提交代码">
                <Form ref={formRef} onSubmit={submitCode} labelWidth={200}>
                    <FormItem name="templateId" label="模板ID(template_id)"
                              help="第三方平台小程序模板库的模板id。需从开发者工具上传代码到第三方平台草稿箱，然后从草稿箱添加到模板库。当前仅支持普通模板，尚未支持标准模板" rules={[{
                        required: true,
                        message: '模板id必选',
                        type: 'error'
                    }]}>
                        <Select style={{width: '300px'}} popupProps={{overlayStyle: {width: '350px'}}}>
                            {/*这个组件极其奇怪，不能静态和动态的混用，会直接白屏，必须一起返回*/}
                            {
                                [<Option label={'123'} value={0} disabled key={`template_-1`}>
                                    <div className="normal_flex">
                                        <p style={{width: '80px', margin: 0}}>模板ID</p>
                                        <p style={{width: '80px', margin: 0}}>版本号</p>
                                        <p style={{flex: 1, margin: 0}}>模板描述</p>
                                    </div>
                                </Option>].concat(templateList.map(i => {
                                    return (
                                        <Option key={`template_${i.templateId}`} value={i.templateId} label={`ID：${i.templateId}`}>
                                            <div className="normal_flex">
                                                <p style={{width: '80px', margin: 0}}>{i.templateId}</p>
                                                <p style={{width: '80px', margin: 0}}>{i.userVersion}</p>
                                                <p style={{flex: 1, margin: 0}}>{truncate(i.userDesc, {length: 10})}</p>
                                            </div>
                                        </Option>
                                    )
                                }))
                            }
                        </Select>
                    </FormItem>
                    <FormItem name="extJson" style={{marginBottom: 0}} label="ext.json配置(ext_json)" rules={[{
                            required: true,
                            message: 'extJson必填',
                            type: 'error'
                        }]}>
                        <Textarea style={{width: '300px'}} />
                    </FormItem>
                    <div className="t-form__controls" style={{marginLeft: '200px'}}>
                        <div className="t-form__help">用于控制ext.json配置文件的内容的参数 <a className="a" href="https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/commit.html" target="_blank">提交代码api说明</a></div>
                    </div>
                    <FormItem name="userVersion" label="代码版本号(user_version)" help="代码版本号，开发者可自定义(长度不超过64个字符)" rules={[{
                        required: true,
                        message: 'userVersion必填',
                        type: 'error'
                    }]}>
                        <Input style={{width: '300px'}} />
                    </FormItem>
                    <FormItem name="userDesc" label="版本描述(user_desc)" help="代码版本描述，开发者可自定义" rules={[{
                        required: true,
                        message: 'userDesc必填',
                        type: 'error'
                    }]}>
                        <Input style={{width: '300px'}} />
                    </FormItem>
                    <FormItem statusIcon={false}>
                        <Button theme="primary" type="submit" style={{marginRight: 10}}>
                            提交
                        </Button>
                        <Button onClick={closeSubmitModal}>取消</Button>
                    </FormItem>
                </Form>
            </Dialog>

            {/*<Drawer visible={visibleDrawer} onClose={() => setVisibleDrawer(false)} confirmBtn={<span />}*/}
            {/*        cancelBtn={<span />} destroyOnClose={true} size="medium" header={"查看驳回原因"}>*/}
            {/*    <div className="normal_flex">*/}
            {/*        <p style={{width: '160px'}}>驳回时间</p>*/}
            {/*        <p>驳回原因</p>*/}
            {/*    </div>*/}
            {/*    <div className={styles.line} />*/}
            {/*    <div className="normal_flex" style={{flexWrap: 'nowrap'}}>*/}
            {/*        <p style={{width: '160px'}}>2022-03-10 15:12:56</p>*/}
            {/*        <p style={{flex: 1}}>驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因驳回原因</p>*/}
            {/*    </div>*/}
            {/*</Drawer>*/}
        </div>
    )
}
