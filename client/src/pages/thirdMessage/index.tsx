import styles from './index.module.less'
import { Tabs, Button, DatePicker, Table, Input, MessagePlugin } from 'tdesign-react';
import {PrimaryTableCol} from "tdesign-react/es/table/type";
import {useEffect, useState} from "react";
import moment from 'moment'
import {get} from "../../utils/axios";
import {getAuthMessageUrl, getMessageConfigUrl, getNormalMessageUrl} from "../../utils/apis";

const { TabPanel } = Tabs;

const authMessageColumn: PrimaryTableCol[] = [
    {
        align: 'center',
        width: 200,
        minWidth: 100,
        className: 'row',
        colKey: 'receiveTime',
        title: '推送时间',
        render: ({ row }) => moment(row.receiveTime).format('YYYY-MM-DD HH:mm:ss')
    },
    {
        align: 'center',
        width: 200,
        minWidth: 100,
        className: 'row',
        colKey: 'infoType',
        title: 'InfoType',
    },
    {
        align: 'center',
        width: 400,
        minWidth: 100,
        className: 'row',
        colKey: 'postBody',
        title: '推送内容',
        render({ row }) {
            return (
                <p style={{ maxWidth: '600px', wordWrap: 'break-word' }}>{row.postBody}</p>
            )
        }
    },
]

const normalMessageColumn: PrimaryTableCol[] = [
    {
        align: 'center',
        width: 100,
        minWidth: 100,
        className: 'row',
        colKey: 'receiveTime',
        title: '推送时间',
        render: ({ row }) => moment(row.receiveTime).format('YYYY-MM-DD HH:mm:ss')
    },
    {
        align: 'center',
        width: 100,
        minWidth: 100,
        className: 'row',
        colKey: 'event',
        title: 'Event',
    },
    {
        align: 'center',
        width: 100,
        minWidth: 100,
        className: 'row',
        colKey: 'msgType',
        title: 'MsgType',
    },
    {
        align: 'center',
        width: 100,
        minWidth: 100,
        className: 'row',
        colKey: 'appid',
        title: 'Appid',
    },
    {
        align: 'center',
        width: 400,
        minWidth: 100,
        className: 'row',
        colKey: 'postBody',
        title: '推送内容',
        render({ row }) {
            return (
                <p style={{ maxWidth: '600px', wordWrap: 'break-word' }}>{row.postBody}</p>
            )
        }
    },
]

const tabs = [{
    label: '授权事件',
    value: 'auth'
}, {
    label: '普通消息与事件',
    value: 'normal'
}]

export default function ThirdMessage() {

    const pageSize = 15

    const [selectedTab, setSelectedTab] = useState<string | number>(tabs[0].value)
    const [isTableLoading, setIsTableLoading] = useState<boolean>(false)

    const [authPage, setAuthPage] = useState(1)
    const [normalPage, setNormalPage] = useState(1)

    const [authData, setAuthData] = useState([])
    const [authDataTotal, setAuthDataTotal] = useState(0)

    const [normalData, setNormalData] = useState([])
    const [normalDataTotal, setNormalDataTotal] = useState(0)

    const [infoTypeInput, setInfoTypeInput] = useState('')
    const [authTimeInput, setAuthTimeInput] = useState<[string, string]>(['', ''])

    const [msgTypeInput, setMsgTypeInput] = useState('')
    const [eventInput, setEventInput] = useState('')
    const [toUserNameInput, setToUserNameInput] = useState('')
    const [normalTimeInput, setNormalTimeInput] = useState<[string, string]>(['', ''])

    const [messageConfig, setMessageConfig] = useState({
        bizPath: "",
        componentPath: "",
        envId: "",
        service: "",
        textMode: ""
    })

    useEffect(() => {
        getMessageConfig()
    }, [])

    const initParams = () => {
        setInfoTypeInput('')
        setAuthTimeInput(['', ''])
        setMsgTypeInput('')
        setEventInput('')
        setToUserNameInput('')
        setNormalTimeInput(['', ''])
    }

    const getMessageConfig = async () => {
        const resp = await get({
            url: getMessageConfigUrl
        })
        if (resp.code === 0) {
            setMessageConfig(resp.data)
        }
    }

    const changeTabs = (val: string | number) => {
        setSelectedTab(val)
        initParams()
    }

    const getTableData = async () => {
        setIsTableLoading(true)
        switch (selectedTab) {
            case tabs[0].value: {
                if (!authTimeInput[0] || !authTimeInput[1]) {
                    MessagePlugin.error('请选择推送时间范围', 2000)
                    break;
                }
                const resp = await get({
                    url: `${getAuthMessageUrl}?infoType=${infoTypeInput}&limit=${pageSize}&offset=${(authPage -1) * pageSize}&startTime=${moment(authTimeInput[0]).valueOf() / 1000}&endTime=${moment(authTimeInput[1]).valueOf() / 1000}`
                })
                if (resp.code === 0) {
                    setAuthData(resp.data.records)
                    setAuthDataTotal(resp.data.total)
                }
                break
            }
            case tabs[1].value: {
                if (!normalTimeInput[0] || !normalTimeInput[1]) {
                    MessagePlugin.error('请选择推送时间范围', 2000)
                    break;
                }
                const resp = await get({
                    url: `${getNormalMessageUrl}?appid=${toUserNameInput}&event=${eventInput}&msgType=${msgTypeInput}&limit=${pageSize}&offset=${(authPage -1) * pageSize}&startTime=${moment(normalTimeInput[0]).valueOf() / 1000}&endTime=${moment(normalTimeInput[1]).valueOf() / 1000}`
                })
                if (resp.code === 0) {
                    setNormalData(resp.data.records)
                    setNormalDataTotal(resp.data.total)
                }
                break
            }
        }
        setIsTableLoading(false)
    }

    return (
        <div className={styles.message}>
            <p className="text">第三方平台消息推送介绍</p>
            <p className="desc">第三方平台消息与事件 URL 用于第三方服务商接收已授权公众号/小程序的消息和事件，第三方平台授权事件URL用于第三方服务商接收只推送给服务商的消息与事件。可通过下方工具快速查看接收到的消息与事件。当前仅支持查看推送至云服务的消息与事件，如果第三方平台的授权事件URL或者消息事件URL配置与下方不符合，则无法查看对应的消息。</p>
            <div className={styles.line} />

            <div className={styles.setting}>
                <div style={{ width: '45%' }}>
                    <p className="text">授权事件 URL 配置</p>
                    <div className={styles.setting_box}>
                        <p className={styles.setting_box_text}>环境ID：{messageConfig.envId}</p>
                        <p className={styles.setting_box_text}>服务名称：{messageConfig.service}</p>
                        <p className={styles.setting_box_text}>消息格式：{messageConfig.textMode}</p>
                        <p className={styles.setting_box_text}>业务路径：{messageConfig.componentPath}</p>
                    </div>
                </div>
                <div style={{ width: '45%' }}>
                    <p className="text">消息与事件 URL 配置</p>
                    <div className={styles.setting_box}>
                        <p className={styles.setting_box_text}>环境ID：{messageConfig.envId}</p>
                        <p className={styles.setting_box_text}>服务名称：{messageConfig.service}</p>
                        <p className={styles.setting_box_text}>消息格式：{messageConfig.textMode}</p>
                        <p className={styles.setting_box_text}>业务路径：{messageConfig.bizPath}</p>
                    </div>
                </div>
            </div>

            <p className="text" style={{ marginTop: '30px' }}>查看消息</p>
            <Tabs value={selectedTab} placement={'top'} size="medium" theme="normal" onChange={changeTabs}>
                <TabPanel value={tabs[0].value} label={tabs[0].label}>
                    <div className="normal_flex" style={{ margin: '10px 0' }}>
                        <div className="normal_flex">
                            <p style={{ marginRight: '20px' }}>InfoType</p>
                            <Input value={infoTypeInput} onChange={(val) => setInfoTypeInput(val as string)} clearable style={{ width: '210px', marginRight: '20px' }} />
                        </div>
                        <div className="normal_flex">
                            <p style={{ marginRight: '20px' }}>推送时间</p>
                            <DatePicker placeholder="必填，需选择日期，否则无法查询" style={{ width: '340px', marginRight: '20px' }} mode="date" onChange={(val: any) => setAuthTimeInput(val)} enableTimePicker range />
                        </div>
                        <Button onClick={getTableData}>查询</Button>
                    </div>
                    <Table
                        loading={isTableLoading}
                        data={authData}
                        columns={authMessageColumn}
                        rowKey="index"
                        tableLayout="auto"
                        verticalAlign="top"
                        size="small"
                        hover
                        // 与pagination对齐
                        pagination={{
                            pageSize,
                            total: authDataTotal,
                            current: authPage,
                            showJumper: true,
                            onCurrentChange:(current) => setAuthPage(current),
                        }}
                    />
                </TabPanel>
                <TabPanel value={tabs[1].value} label={tabs[1].label}>
                    <div className="normal_flex" style={{ margin: '10px 0' }}>
                        <div className="normal_flex">
                            <p style={{ marginRight: '20px' }}>MsgType</p>
                            <Input value={msgTypeInput} onChange={(val) => setMsgTypeInput(val as string)} clearable style={{ width: '120px', marginRight: '20px' }} />
                        </div>
                        <div className="normal_flex">
                            <p style={{ marginRight: '20px' }}>Event</p>
                            <Input value={eventInput} onChange={(val) => setEventInput(val as string)} clearable style={{ width: '120px', marginRight: '20px' }} />
                        </div>
                        <div className="normal_flex">
                            <p style={{ marginRight: '20px' }}>ToUserName</p>
                            <Input value={toUserNameInput} onChange={(val) => setToUserNameInput(val as string)} clearable style={{ width: '140px', marginRight: '20px' }} />
                        </div>
                        <div className="normal_flex">
                            <p style={{ marginRight: '20px' }}>推送时间</p>
                            <DatePicker placeholder="必填，需选择日期，否则无法查询" style={{ width: '340px', marginRight: '20px' }} mode="date" onChange={(val: any) => setNormalTimeInput(val)} enableTimePicker range />
                        </div>
                        <Button onClick={getTableData}>查询</Button>
                    </div>
                    <Table
                        loading={isTableLoading}
                        data={normalData}
                        columns={normalMessageColumn}
                        rowKey="index"
                        tableLayout="auto"
                        verticalAlign="top"
                        size="small"
                        hover
                        // 与pagination对齐
                        pagination={{
                            pageSize,
                            total: normalDataTotal,
                            current: normalPage,
                            showJumper: true,
                            onCurrentChange:(current) => setNormalPage(current),
                        }}
                    />
                </TabPanel>
            </Tabs>
        </div>
    )
}
