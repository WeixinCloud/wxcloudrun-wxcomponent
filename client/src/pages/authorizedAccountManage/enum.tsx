import {PrimaryTableCol} from "tdesign-react/es/table/type";
import {copyMessage} from "../../utils/common";

export const officialAccountAuthType: Record<string, string> = {
    "-1": "未认证",
    "0": "微信认证",
    "1": "新浪微博认证",
    "2": "腾讯微博认证",
    "3": "已资质认证通过但还未通过名称认证",
    "4": "已资质认证通过、还未通过名称认证，但通过了新浪微博认证",
    "5": "已资质认证通过、还未通过名称认证，但通过了腾讯微博认证"
}

export const miniProgramAuthType: Record<string, string> = {
    "-1": "未认证",
    "0": "微信认证",
}

export const tokenColumn: PrimaryTableCol[] = [
    {
        align: 'center',
        minWidth: 100,
        className: 'row',
        colKey: 'token',
        title: 'Token',
    },
    {
        align: 'center',
        minWidth: 100,
        className: 'row',
        colKey: 'id',
        title: '操作',
        render({ row }) {
            return (
                <a className="a" onClick={() => copyMessage(row.token)}>复制</a>
            );
        },
    },
]

export const tabs = [{
    label: '普通授权帐号',
    value: 'normal'
}, {
    label: '代开发小程序',
    value: 'code'
}]

export const serviceStatus: Record<number, string> = {
    1: '正常',
    0: '主动暂停服务'
}

export  const accountStatus: Record<number, string> = {
    0: '正常',
    1: '已冻结',
    2: '因违规暂停服务',
}

export const registerType: Record<number, string> = {
    '-1': '未知',
    0: '普通方式注册',
    2: '通过复用公众号创建小程序api注册',
    6: '通过法人扫脸创建企业小程序api注册',
    13: '通过创建试用小程序api注册',
    15: '通过联盟控制台注册',
    16: '通过创建个人小程序api注册',
    17: '通过创建个人交易小程序api注册',
    19: '通过试用小程序转正api注册',
    22: '通过复用商户号创建企业小程序api注册',
    23: '通过复用商户号转正api注册'
}

export  const normalAccountStatus: Record<number, string> = {
    1: '正常',
    14: '已注销',
    16: '已封禁',
    18: '已告警',
    19: '已冻结',
}
