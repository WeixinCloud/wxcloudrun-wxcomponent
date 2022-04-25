import {Upload, Button, Alert, Select, Textarea, Input, Form, MessagePlugin, Loading} from 'tdesign-react';
import styles from './index.module.less'
import {useEffect, useMemo, useRef, useState} from "react";
import {request} from "../../../utils/axios";
import {getCategoryRequest, submitAuditRequest, uploadFileRequest} from "../../../utils/apis";
import {useSearchParams} from "react-router-dom";
import {UploadFile} from "tdesign-react/es/upload/type";
import {routes} from "../../../config/route";

const { FormItem } = Form
const { Option } = Select

const tipMessage = <div>
    <p className="text">提交审核前端须知</p>
    <p className={styles.bottom10}>- 提交审核前小程序需完成名称、头像、简介以及类目设置</p>
    <p className={styles.bottom10}>- 如该小程序中使用了涉及用户隐私接口，例如获取用户头像、手机号等，需先完成"用户隐私保护指引"</p>
    <p className={styles.bottom10}>- 如该小程序已经绑定为第三方平台开发小程序，需前往第三方平台-代开发小程序进行解除绑定</p>
    <p className={styles.bottom10}>- 提交的小程序功能完整，可正常打开和运行，而不是测试版或 Demo，多次提交测试内容或 Demo，将受到相应处罚</p>
    <p style={{ margin: 0 }}>- 确保小程序符合<a className="a" href="https://developers.weixin.qq.com/miniprogram/product/" target="_blank">《微信小程序平台运营规范》</a>和确保已经提前了解<a href="https://developers.weixin.qq.com/miniprogram/product/reject.html" target="_blank" className="a">《微信小程序平台审核常见被拒绝情形》</a></p>
</div>

const sceneOptions = [{
    label: '不涉及用户生成内容',
    value: 0
}, {
    label: '用户资料',
    value: 1
}, {
    label: '图片',
    value: 2
}, {
    label: '视频',
    value: 3
}, {
    label: '文本',
    value: 4
}, {
    label: '其他',
    value: 5
}]

const methodOptions = [{
    label: '使用平台建议的内容安全API',
    value: 1
}, {
    label: '使用其他的内容审核产品',
    value: 2
}, {
    label: '通过人工审核把关',
    value: 3
}, {
    label: '未做内容审核把关',
    value: 4
}]

const hasAuditTeamOptions = [{
    label: '无',
    value: 0
}, {
    label: '有',
    value: 1
}]

// 用来做类目map时候的key
let id = 0

type ICategory = {
    id?: number
    firstClass: string
    secondClass: string
    thirdClass?: string
    firstId: number
    secondId: number
    thirdId?: number
}

export default function SubmitAudit() {

    const formRef = useRef() as any
    const [videoFile, setVideoFile] = useState<IUploadFile[]>([])
    const [picFile, setPicFile] = useState<IUploadFile[]>([])
    const [stuffFile, setStuffFile] = useState<IUploadFile[]>([])
    const [stuffLoading, setStuffLoading] = useState(false)
    const [picLoading, setPicLoading] = useState(false)
    const [videoLoading, setVideoLoading] = useState(false)
    const [submitLoading, setSubmitLoading] = useState(false)
    const [searchParams] = useSearchParams();
    const [otherMpType, setOtherMpType] = useState<ICategory[]>([])
    const [categoryList, setCategoryList] = useState<ICategory[]>([])

    const appId = useMemo(() => {
        return searchParams.get('appId')
    }, [searchParams])

    useEffect(() => {
        getCategoryList()
    }, [])

    interface IUploadFile extends UploadFile {
        mediaId?: string
    }

    const toggleLoading = (name: 'video' | 'pic' | 'stuff' | 'submit') => {
        switch (name) {
            case "video": {
                setVideoLoading(current => !current)
                break
            }
            case "pic": {
                setPicLoading(current => !current)
                break
            }
            case "stuff": {
                setStuffLoading(current => !current)
                break
            }
            case "submit": {
                setSubmitLoading(current => !current)
                break
            }
        }
    }

    const uploadFile = async (files: IUploadFile[], maxSize: number, type: 'image' | 'video', dataName: 'video' | 'pic' | 'stuff') => {
        toggleLoading(dataName)
        // 先查一遍有没有超限，这个组件自动限制有问题，根本没法用
        for (let i = 0; i < files.length; i++) {
            // 鬼知道这个组件设计为什么size可能是undefined
            if ((files[i].size || 0) > maxSize) {
                MessagePlugin.error('请勿上传超过限制大小的文件')
                toggleLoading(dataName)
                return false
            }
        }

        // 将没上传的文件上传
        for (let i = 0; i < files.length; i++) {
            // 鬼知道这个组件设计为什么size可能是undefined
            if (files[i].status === 'success') {
                continue
            }
            const data = new FormData()
            data.append('media', files[i].raw as Blob)
            const resp = await request({
                request: {
                    url: `${uploadFileRequest.url}?appid=${appId}&type=${type}`,
                    method: uploadFileRequest.method
                },
                data
            })
            if (resp.code === 0) {
                files[i].mediaId = resp.data.mediaId
                files[i].status = 'success'
            } else {
                toggleLoading(dataName)
                return false
            }
        }
        switch (dataName) {
            case "video": {
                setVideoFile(files)
                break
            }
            case "pic": {
                setPicFile(files)
                break
            }
            case "stuff": {
                setStuffFile(files)
                break
            }
        }
        toggleLoading(dataName)
        return true
    }

    const getCategoryList = async () => {
        const resp = await request({
            request: getCategoryRequest,
            data: {
                appid: appId,
            }
        })
        if (resp.code === 0) {
            setCategoryList(resp.data.categoryList)
        }
    }

    const addOtherMpType = () => {
        if (otherMpType.length === 4) return
        setOtherMpType(otherMpType.concat({
            id,
            firstClass: '',
            secondClass: '',
            thirdClass: '',
            firstId: -1,
            secondId: -1,
            thirdId: -1,
        }))
        id++
    }

    const removeOtherMpType = (index: number) => {
        if (otherMpType.length === 0) return
        otherMpType.splice(index, 1)
        // 直接set自己不生效
        setOtherMpType(JSON.parse(JSON.stringify(otherMpType)))
    }

    const onChangeSelectCategory = (index: number, categoryIndex: number) => {
        otherMpType[index] = {
            ...otherMpType[index],
            ...categoryList[categoryIndex]
        }
    }

    const submitCode = async (e: {validateResult: any}) => {
        if (e.validateResult !== true) {
            return
        }
        toggleLoading('submit')
        const { templateId, versionDesc, feedbackInfo, auditDesc, otherSceneDesc, hasAuditTeam, scene, method } = formRef.current.getAllFieldsValue()
        let data = {
            appid: appId,
            itemList: otherMpType.filter(i => i.firstId).concat(categoryList[templateId]), // 审核项列表
            previewInfo: {
                videoIdList: (videoFile || []).map((i: IUploadFile) => i.mediaId),
                picIdList: (picFile || []).map((i: IUploadFile) => i.mediaId)
            },// 预览信息
            versionDesc,
            feedbackInfo,
            feedbackStuff: (stuffFile || []).map((i: IUploadFile) => i.mediaId).join('|'),
            ugcDeclare: {
                scene,
                otherSceneDesc,
                method,
                hasAuditTeam,
                auditDesc,
            }
        } as any
        if ((scene || []).includes(0)) {
            data.ugcDeclare.scene = [0]

        }
        const resp = await request({
            request: {
                url: `${submitAuditRequest.url}?appid=${appId}`,
                method: submitAuditRequest.method
            },
            data
        })
        if (resp.code === 0) {
            MessagePlugin.success('提交审核成功')
            window.location.href = `#${routes.miniProgramVersion.path}?appId=${appId}`
        }
        toggleLoading('submit')
    }

    return (
        <div>
            <Alert theme="info" icon={<span />} maxLine={0} message={tipMessage} style={{ marginBottom: '10px' }} />
            <Form ref={formRef} onSubmit={submitCode} labelWidth={250}>
                <p className="text">配置审核列表</p>
                <FormItem name="templateId" label="小程序类目" rules={[{
                    required: true,
                    message: '小程序类目必选',
                    type: 'error'
                }]}>
                    <Select style={{ width: '400px' }} popupProps={{ overlayStyle: { width: '450px' } }}>
                        {
                            categoryList.map((i, index) => {
                                return (
                                    <Option key={`${i.firstClass} - ${i.secondClass}${i.thirdClass ? ` - ${i.thirdClass}` : ''}`}  label={`${i.firstClass} - ${i.secondClass}${i.thirdClass ? ` - ${i.thirdClass}` : ''}`} value={index} />
                                )
                            })
                        }
                    </Select>
                    {
                        otherMpType.length < 4
                        &&
                        <a className="a" style={{ marginLeft: '6px' }} onClick={addOtherMpType}>添加</a>
                    }
                </FormItem>
                {
                    otherMpType.map((i, index) => {
                        return (
                            <div className="t-form__item" key={`other_mp_type_${i.id}`}>
                                <div className="t-form__label" style={{ width: '250px' }} />
                                <div className="t-form__controls" style={{ marginLeft: '250px' }}>
                                    <div className="normal_flex">
                                        <Select style={{ width: '400px' }} popupProps={{ overlayStyle: { width: '450px' } }} onChange={(val) => onChangeSelectCategory(index, val as number)}>
                                            {
                                                categoryList.map((i, categoryIndex) => {
                                                    return (
                                                        <Option key={`${i.firstClass} - ${i.secondClass}${i.thirdClass ? ` - ${i.thirdClass}` : ''}`} label={`${i.firstClass} - ${i.secondClass}${i.thirdClass ? ` -${i.thirdClass}` : ''}`} value={categoryIndex} />
                                                    )
                                                })
                                            }
                                        </Select>
                                        <a className="a" style={{ marginLeft: '6px' }} onClick={() => removeOtherMpType(index)}>删除</a>
                                    </div>
                                </div>
                            </div>
                        )
                    })
                }
                <p className="text">配置预览信息</p>
                <FormItem name="videoIdList" label="视频预览(video_id_list)" help="可上传小程序使用录屏，最多上传1个视频。视频支持mp4格式，视频大小不超过10MB">
                    <Loading loading={videoLoading}>
                        <Upload
                            files={videoFile}
                            onChange={(files) => uploadFile(files, 10 * 1024 * 1024, 'video', 'video')}
                            theme="file-flow"
                            accept="video/mp4"
                            multiple
                            max={1}
                            sizeLimit={{ size: 20, unit: 'MB', message: '视频大小不超过 10 MB' }}
                        />
                    </Loading>
                </FormItem>
                <FormItem name="picIdList" label="图片预览(pic_id_list)" help="可上传小程序截图，最多上传10张图片。图片支持jpg、jpeg、bmp、gif或png格式，图片大小不超过5MB">
                    <Loading loading={picLoading}>
                    <Upload
                        files={picFile}
                        onChange={(files) => uploadFile(files, 5 * 1024 * 1024, 'image', 'pic')}
                        theme="file-flow"
                        accept="image/*"
                        multiple
                        max={10}
                        sizeLimit={{ size: 5, unit: "MB", message: "图片大小不超过 5 MB" }}
                    />
                    </Loading>
                </FormItem>
                <p className="text">信息安全声明</p>
                <FormItem name="scene" label="UGC场景(scene)">
                    <Select style={{ width: '400px' }} options={sceneOptions} multiple />
                </FormItem>
                <FormItem name="otherSceneDesc" label="场景说明(other_scene_desc)" help="当scene选'其他'时需填写场景说明，不超过256个字">
                    <Input style={{ width: '400px' }} />
                </FormItem>
                <FormItem name="method" label="内容安全机制(method)">
                    <Select style={{ width: '400px' }} options={methodOptions} multiple />
                </FormItem>
                <FormItem name="hasAuditTeam" label="是否有审核团队(has_audit_team)">
                    <Select style={{ width: '400px' }} options={hasAuditTeamOptions} />
                </FormItem>
                <FormItem name="auditDesc" label="审核机制描述(audit_desc)" help="说明当前对UGC内容的审核机制，不超过256个字">
                    <Input style={{ width: '400px' }} />
                </FormItem>
                <p className="text">其他配置信息</p>
                <FormItem name="versionDesc" label="版本描述(version_desc)" help="小程序版本说明和功能解释；如登录小程序需要帐号密码">
                    <Input style={{ width: '400px' }} />
                </FormItem>
                <FormItem name="feedbackInfo" label="反馈内容(feedback_desc)" help="如上一次提审被驳回，可在该出补充就驳回项进行解释说明，以便审核人员再次进行审核">
                    <Textarea style={{ width: '400px' }} />
                </FormItem>
                <FormItem name="feedbackStuff" label="反馈截图(feedback_stuff)" help="可上传截图辅助解释，最多上传5张图片。图片支持jpg、jpeg、bmp、gif或png格式，图片大小不超过5MB">
                    <Loading loading={stuffLoading}>
                    <Upload
                        files={stuffFile}
                        onChange={(files) => uploadFile(files, 5 * 1024 * 1024, 'image', 'stuff')}
                        theme="file-flow"
                        accept="image/*"
                        multiple
                        max={5}
                        sizeLimit={{ size: 5, unit: 'MB', message: '图片大小不超过 5 MB' }}
                    />
                    </Loading>
                </FormItem>
                <FormItem statusIcon={false} style={{ marginTop: '20px' }}>
                    <Button variant="outline" style={{ marginRight: 10 }} onClick={() => window.history.back()}>
                        上一步
                    </Button>
                    <Button theme="primary" type="submit" loading={submitLoading}>
                        提交
                    </Button>
                </FormItem>
            </Form>
        </div>
    )
}
