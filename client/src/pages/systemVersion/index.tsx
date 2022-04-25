import styles from './index.module.less'

export default function SystemVersion() {

    return (
        <div>
            <p className="text">系统介绍</p>
            <p className="desc">"服务商微管家"旨在帮助服务商更高效基于第三方平台开展业务，后续将持续开放更多功能模块，提供更多开发调试工具以及批量管理工具</p>
            <div className={styles.line} />
            <div className="normal_flex">
                <p style={{width: '100px'}}>当前版本</p>
                <p style={{marginRight: '20px'}} className="desc">V 2.2.0</p>
            </div>
            <div className="normal_flex">
                <p style={{width: '100px'}}>更新时间</p>
                <p style={{marginRight: '20px'}} className="desc">2022-04-25</p>
            </div>

            <p style={{marginTop: '40px'}} className="text">系统更新日志</p>
            <div className="normal_flex">
                <p className="desc">前往第三方平台官方文档中心可查看详细系统更新日志，</p>
                <a style={{marginRight: '20px'}} className="a" href="https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/product/management-tools.html" target="_blank">立即前往</a>
            </div>

        </div>
    )
}
