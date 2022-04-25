import {Table, Input, PopConfirm, Dialog, Tabs, MessagePlugin} from 'tdesign-react';
import { SearchIcon } from 'tdesign-icons-react'
import {useEffect, useState} from "react";
import {request} from "../../utils/axios";
import {
    changeServiceStatusRequest,
    getAuthAccessTokenRequest,
    getAuthorizedAccountRequest,
    getDevMiniProgramListRequest, getQrcodeRequest
} from "../../utils/apis";
import {PrimaryTableCol} from "tdesign-react/es/table/type";
import moment from "moment";
import {copyMessage} from "../../utils/common";
import {
    officialAccountAuthType,
    miniProgramAuthType,
    tokenColumn,
    tabs,
    serviceStatus,
    accountStatus, registerType, normalAccountStatus
} from './enum'
import {routes} from "../../config/route";

const { TabPanel } = Tabs

export default function AuthorizedAccountManage() {

    const accountColumn: PrimaryTableCol[] = [
        {
            align: 'center',
            minWidth: 100,
            colKey: 'appid',
            title: 'AppID',
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'userName',
            title: '原始ID',
        },
        {
            align: 'center',
            minWidth: 120,
            colKey: 'nickName',
            title: '名称',
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'appType',
            title: '帐号类型',
            cell: ({ row }) => {
                return row.appType === 0 ? '小程序' : '公众号'
            }
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'authTime',
            title: '授权时间',
            render: ({ row }) => moment(row.authTime).format('YYYY-MM-DD HH:mm:ss')
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'principalName',
            title: '主体信息',
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'registerType',
            title: '注册类型',
            render: ({ row }) => registerType[row.registerType]
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'accountStatus',
            title: '帐号状态',
            render: ({ row }) => normalAccountStatus[row.accountStatus]
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'isPhoneConfigured',
            title: '已绑手机号',
            render: ({ row }) => row.basicConfig ? row.basicConfig.isPhoneConfigured ? '已绑定' : '未绑定' : '-'
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'isEmailConfigured',
            title: '已绑邮箱',
            render: ({ row }) => row.basicConfig ? row.basicConfig.isEmailConfigured ? '已绑定' : '未绑定' : '-'
        },
        {
            align: 'center',
            minWidth: 100,
            className: 'row',
            colKey: 'verifyInfo',
            title: '认证类型',
            cell: ({ row }) => {
                return row.appType === 0 ? miniProgramAuthType[String(row.verifyInfo)] : officialAccountAuthType[String(row.verifyInfo)]
            }
        },
        {
            align: 'center',
            minWidth: 100,
            className: 'row',
            colKey: 'funcInfo',
            title: '授权权限集ID',
        },
        {
            align: 'center',
            fixed: 'right',
            width: 210,
            minWidth: 210,
            className: 'row',
            colKey: 'id',
            title: '操作',
            render({ row }) {
                if (row.accountStatus === 16) {
                    return (
                        <div style={{ width: '210px' }}>
                            <p className="desc">该帐号已封禁</p>
                        </div>
                    );
                } else {
                    return (
                        <div style={{ width: '210px' }}>
                            <PopConfirm content="从数据库获取 token，非重新生成token，不会导致上一个 token 被刷新而失效" onConfirm={() => createToken(row.appid)}>
                                <a className="a" style={{ margin: '0 10px' }}>获取token</a>
                            </PopConfirm>
                            <a className="a" onClick={() => copyMessage(row.refreshToken)}>复制refresh_token</a>
                        </div>
                    );
                }
            },
        },
    ]

    const miniProgramColumn: PrimaryTableCol[] = [
        {
            align: 'center',
            minWidth: 200,
            colKey: 'appid',
            title: 'AppID',
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'nickName',
            title: '名称',
        },
        {
            align: 'center',
            minWidth: 120,
            colKey: 'releaseInfo',
            title: '发布状态',
            render: ({ row }) => row.releaseInfo ? '已有上线版本' : '尚未发布'
        },
        {
            align: 'center',
            minWidth: 100,
            colKey: 'serviceStatus',
            title: '服务状态',
            render: ({ row }) => serviceStatus[row.serviceStatus]
        },
        {
            align: 'center',
            fixed: 'right',
            width: 210,
            minWidth: 210,
            className: 'row',
            colKey: 'id',
            title: '操作',
            render({ row }) {
                return (
                    <div style={{ width: '210px', display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
                        {
                            (row.releaseInfo && row.funcInfo.includes(17))
                            &&
                            <a className="a" style={{ marginRight: '5px' }} onClick={() => getMiniProgramCode(row.appid)}>获取小程序码</a>
                        }
                        {
                            row.serviceStatus === 0
                            &&
                            <a className="a" style={{ marginRight: '5px' }} onClick={() => openServiceStatus(row.appid)}>恢复服务</a>
                        }
                        <a className="a" href={`#${routes.miniProgramVersion.path}?appId=${row.appid}`}>版本管理</a>
                    </div>
                );
            },
        },
    ]

    const pageSize = 15

    const [accountList, setAccountList] = useState([])
    const [miniProgramList, setMiniProgramList] = useState([])
    const [accountTotal, setAccountTotal] = useState(0)
    const [mpAccountTotal, setMpAccountTotal] = useState(0)
    const [currentPage, setCurrentPage] = useState(1)
    const [mpCurrentPage, setMpCurrentPage] = useState(1)
    const [miniProgramAppIdInput, setMiniProgramAppIdInput] = useState<string | number>('')
    const [appIdInput, setAppIdInput] = useState<string | number>('')
    const [visibleTokenModal, setVisibleTokenModal] = useState(false)
    const [tokenData, setTokenData] = useState([{token: ''}])
    const [selectedTab, setSelectedTab] = useState<string | number>(tabs[0].value)
    const [visibleQrcode, setVisibleQrcode] = useState(false)
    const [qrcode, setQrcode] = useState('')

    useEffect(() => {
        if (selectedTab === tabs[0].value) {
            getAccountList()
        }
    }, [currentPage, selectedTab])

    useEffect(() => {
        if (selectedTab === tabs[1].value) {
            getMiniProgramList()
        }
    }, [mpCurrentPage, selectedTab])

    const createToken = async (appId: string) => {
        const resp = await request({
            request: getAuthAccessTokenRequest,
            data: {
                appid: appId
            }
        })
        if (resp.code === 0) {
            setTokenData([{
                token: resp.data.token
            }])
            setVisibleTokenModal(true)
        }
    }

    const getAccountList = async () => {
        const resp = await request({
            request: getAuthorizedAccountRequest,
            data: {
                offset: (currentPage - 1) * pageSize,
                limit: pageSize,
                appid: appIdInput
            }
        })
        if (resp.code === 0) {
            setAccountList(resp.data.records)
            setAccountTotal(resp.data.total)
        }
    }

    const getMiniProgramList = async () => {
        const resp = await request({
            request: getDevMiniProgramListRequest,
            data: {
                offset: (mpCurrentPage - 1) * pageSize,
                count: pageSize,
                appid: miniProgramAppIdInput,
            }
        })
        if (resp.code === 0) {
            setMiniProgramList(resp.data.records)
            setMpAccountTotal(resp.data.total)
        }
    }

    const getMiniProgramCode = async (appId: string) => {
        const resp = await request({
            request: getQrcodeRequest,
            data: {
                appid: appId,
            }
        })
        if (resp.code === 0) {
            setQrcode('data:image/png;base64,' + resp.data.releaseQrCode)
            setVisibleQrcode(true)
        }

    }

    const openServiceStatus = async (appId: string) => {
        const resp = await request({
            request: {
                url: `${changeServiceStatusRequest.url}?appid=${appId}`,
                method: changeServiceStatusRequest.method
            },
            data: {
                action: "open"
            }
        })
        if (resp.code === 0) {
            MessagePlugin.success('恢复服务状态成功')
            getMiniProgramList()
        }
    }

    return (
        <div>
            <p className="text">授权帐号介绍</p>
            <div className="normal_flex">
                <div className="blue_circle" />
                <p className="desc"
                   style={{margin: 0}}>授权帐号指的是获得公众号或者小程序管理员授权的帐号，服务商可为授权帐号提供代开发、代运营等服务。</p>
            </div>
            <div className="normal_flex">
                <div className="blue_circle" />
                <p className="desc">代开发小程序指的是小程序管理员将权限集id为18的"小程序开发与数据分析"权限授权给该第三方，服务商可代小程序提交代码、发布上线等</p>
            </div>
            <Tabs value={selectedTab} placement={'top'} size="medium" theme="normal"
                  onChange={val => setSelectedTab(val)}>
                <TabPanel value={tabs[0].value} label={tabs[0].label}>
                    <Input value={appIdInput} onChange={setAppIdInput} style={{ width: '400px', margin: '10px 0' }} placeholder="请输入 AppID，不支持模糊搜索" suffixIcon={<a className="a" onClick={getAccountList}><SearchIcon /></a>} />
                    <Table
                        data={accountList}
                        columns={accountColumn}
                        rowKey="id"
                        tableLayout="auto"
                        verticalAlign="middle"
                        size="small"
                        hover
                        // 与pagination对齐
                        pagination={{
                            pageSize,
                            total: accountTotal,
                            current: currentPage,
                            pageSizeOptions: [15],
                            onCurrentChange: setCurrentPage,
                        }}
                    />
                </TabPanel>
                <TabPanel value={tabs[1].value} label={tabs[1].label}>
                    <Input value={miniProgramAppIdInput} onChange={setMiniProgramAppIdInput} style={{ width: '400px', margin: '10px 0' }} placeholder="请输入 AppID，不支持模糊搜索" suffixIcon={<a className="a" onClick={getMiniProgramList}><SearchIcon /></a>} />
                    <Table
                        data={miniProgramList}
                        columns={miniProgramColumn}
                        rowKey="appid"
                        tableLayout="auto"
                        verticalAlign="middle"
                        size="small"
                        hover
                        pagination={{
                            pageSize,
                            total: mpAccountTotal,
                            current: mpCurrentPage,
                            pageSizeOptions: [15],
                            onCurrentChange: setMpCurrentPage,
                        }}
                    />
                </TabPanel>
            </Tabs>

            <Dialog header="AuthorizerAccessToken" visible={visibleTokenModal} footer={null} onClose={() => setVisibleTokenModal(false)}>
                <Table
                    data={tokenData}
                    columns={tokenColumn}
                    rowKey="id"
                    tableLayout="auto"
                    verticalAlign="middle"
                    size="small"
                />
            </Dialog>

            <Dialog header="获取小程序码" visible={visibleQrcode} footer={null} onClose={() => setVisibleQrcode(false)}>
                <div style={{ textAlign: 'center' }}>
                    <img src={qrcode} style={{width: '200px', height: '200px'}} alt="" />
                </div>
            </Dialog>

        </div>
    )
}
