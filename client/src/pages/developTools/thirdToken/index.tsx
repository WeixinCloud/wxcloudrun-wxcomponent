import {useState} from 'react'
import styles from './index.module.less'
import {Table, Button, PopConfirm} from 'tdesign-react';
import {PrimaryTableCol} from "tdesign-react/es/table/type";
import {request} from "../../../utils/axios";
import {getComponentTokenRequest, getTicketRequest} from "../../../utils/apis";
import moment from "moment";
import {copyMessage} from "../../../utils/common";
import {routes} from "../../../config/route";

const componentTokenColumn: PrimaryTableCol[] = [{
    align: 'left',
    minWidth: 100,
    className: 'row',
    colKey: 'token',
    title: 'Token',
    render({ row }) {
        return (
            <p style={{ maxWidth: '600px', wordWrap: 'break-word' }}>{row.token}</p>
        )
    }
}, {
    align: 'center',
    width: 300,
    minWidth: 100,
    className: 'row',
    colKey: 'expiresTime',
    title: '过期时间',
    render: ({ row }) => moment(row.expiresTime).format('YYYY-MM-DD HH:mm:ss')
}, {
    align: 'center',
    width: 100,
    minWidth: 100,
    className: 'row',
    colKey: 'a',
    title: '操作',
    render({ row }) {
        return (
            <a className="a" onClick={() => copyMessage(row.token)}>复制</a>
        );
    },
}]

const ticketColumn: PrimaryTableCol[] = [{
    align: 'left',
    minWidth: 100,
    width: 800,
    className: 'row',
    colKey: 'ticket',
    title: 'Ticket',
}, {
    align: 'center',
    width: 100,
    minWidth: 100,
    className: 'row',
    colKey: 'a',
    title: '操作',
    render({ row }) {
        return (
            <a className="a" onClick={() => copyMessage(row.ticket)}>复制</a>
        );
    },
}]

export default function ThirdToken() {

    const [isTicketLoading, setIsTicketLoading] = useState<boolean>(false)
    const [isComponentTokenLoading, setIsComponentTokenLoading] = useState<boolean>(false)
    const [ticket, setTicket] = useState<{
        ticket: string
    }[]>([])
    const [componentToken, setComponentToken] = useState<{
        token: string
        expiresTime: string
    }[]>([])

    const getComponentVerifyTicket = async () => {
        setIsTicketLoading(true)
        const resp = await request({
            request: getTicketRequest
        })
        if (resp.code === 0) {
            setTicket([{
                ticket: resp.data.ticket,
            }])
        }
        setIsTicketLoading(false)
    }

    const getComponentToken = async () => {
        setIsComponentTokenLoading(true)
        const resp = await request({
            request: getComponentTokenRequest
        })
        if (resp.code === 0) {
            setComponentToken([{
                expiresTime: moment().add(7200, 'seconds').format('YYYY-MM-DD HH:mm:ss'),
                token: resp.data.token,
            }])
        }
        setIsComponentTokenLoading(false)
    }

    return (
        <div className={styles.token}>

            <p className="text">component_verify_ticket 介绍</p>
            <div className="normal_flex">
                <div className="blue_circle" />
                <p className="desc" style={{ margin: 0 }}>获取 component_verify_ticket 后可通过接口或者通过下方的功能生成 component_access_token。更多介绍可前往查看<a href="https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket.html" target="_blank" className="a">官方文档</a></p>
            </div>
            <div className="normal_flex">
                <div className="blue_circle" />
                <p className="desc">为了节省成本，ticket 推送默认是关闭的，如需获取 ticket 则需要返回"第三方平台-开发信息"模块进行开启，<a href="https://open.weixin.qq.com/" target="_blank" className="a">立即前往</a></p>
            </div>
            <div className={styles.line} />
            <Button style={{ marginTop: '20px' }} onClick={getComponentVerifyTicket}>立即获取</Button>
            <Table
                loading={isTicketLoading}
                data={ticket}
                columns={ticketColumn}
                rowKey="ticket"
                tableLayout="auto"
                verticalAlign="middle"
                size="small"
            />

            <p style={{marginTop: '40px'}} className="text">component_access_token 介绍</p>
            <div className="normal_flex" style={{ flexWrap: 'nowrap', alignItems: 'center' }}>
                <div className="blue_circle" />
                <p className="desc" style={{ margin: '0' }}>component_access_token 是第三方平台接口的调用凭据。令牌的获取是有限制的，每个令牌的有效期为 2 小时，请自行做好令牌的管理，在令牌快过期时（比如1小时50分），重新调用接口获取。接口详情可查看<a href="https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html" target="_blank" className="a">官方文档</a></p>
            </div>
            <div className="normal_flex">
                <div className="blue_circle" />
                <p className="desc">生成 component_access_token 依赖第三方平台账号 appid 和 secret 信息，需前往"系统管理-Secret与密码管理"完成配置后再使用，<a href={`#${routes.passwordManage.path}`} target="_blank" className="a">立即前往</a></p>
            </div>
            <div className={styles.line} />
            <PopConfirm content={'点击确认生成 token 后会导致上一个 token 被刷新而失效，请谨慎操作'} onConfirm={getComponentToken}>
                <Button style={{ marginTop: '20px' }}>立即获取</Button>
            </PopConfirm>
            <Table
                loading={isComponentTokenLoading}
                data={componentToken}
                columns={componentTokenColumn}
                rowKey="componentToken"
                tableLayout="auto"
                verticalAlign="middle"
                size="small"
            />

        </div>
    )
}
