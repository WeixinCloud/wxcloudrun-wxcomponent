import styles from './index.module.less'
import {routes} from "../../config/route";

export default function DevelopTools() {

    return (
        <div className={styles.box}>
            <div className={styles.box_item}>
                <img src="https://res.wx.qq.com/op_res/y3myMR2C6QlqM7iwob-WEGUrD3BeFjD7LWIlFGpEVrFeJhd-w0pBKofG--aEdOx02yvazKTQut3NevGhFa8CNg" alt="" className={styles.box_item_logo} />
                <p className={styles.box_item_title}>消息推送查询工具</p>
                <p className={styles.box_item_desc}>可查询以及搜索微信官方推送至服务商的全部消息与事件。<a href={`#${routes.thirdMessage.path}`} className="a">立即前往</a></p>
            </div>

            <div className={styles.box_item}>
                <img src="https://res.wx.qq.com/op_res/y3myMR2C6QlqM7iwob-WEFK5SgSXIVlV_xA-nhtar6FbWjZ9iphx0zja8gPTLzDUmX9OwFv-iMYDzQKnTsiyDg" alt="" className={styles.box_item_logo} />
                <p className={styles.box_item_title}>在线生成 Token 工具</p>
                <p className={styles.box_item_desc}>生成第三方平台 Token 后快速调试官方接口。<a href={`#${routes.thirdToken.path}`} className="a">立即前往</a></p>
            </div>

            <div className={styles.box_item}>
                <img src="https://res.wx.qq.com/op_res/IcLD7eWiljRyBp6gDCD92xx-kBz-AiWzqhE0ubiluxMd730E6ta_rZcnM9pjdu7eu3nNUbEVjspoQ5rwnWuwLQ" alt="" className={styles.box_item_logo} />
                <p className={styles.box_item_title}>在线接口调试工具</p>
                <p className={styles.box_item_desc}>敬请期待，下个版本即将支持。如有相关疑问可加官方交流群进行反馈。</p>
            </div>
        </div>
    )
}
