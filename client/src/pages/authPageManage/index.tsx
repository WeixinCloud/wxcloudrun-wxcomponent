import styles from './index.module.less'
import {routes} from "../../components/Console";
import {copyMessage} from "../../utils/common";

export default function AuthPageManage() {

    return (
        <div>
            <p style={{marginTop: '40px'}} className="text">授权链接生成器介绍</p>
            <div style={{ margin: '20px 0' }}>
                <div className="normal_flex" style={{ marginTop: '10px' }}>
                    <div className="blue_circle" />
                    <p className="desc" style={{ margin: '0' }}>服务商需要获得商家授权后方可代商家开发、运营、管理商家公众号和小程序，因此需要生成授权链接，引导商家完成授权。</p>
                </div>
                <div className="normal_flex" style={{ marginTop: '10px' }}>
                    <div className="blue_circle" />
                    <p className="desc" style={{ margin: '0' }}>你可在该平台直接跳转授权，也可以将链接分享给商家，也可以复制授权链接到企业官网，引导用户授权。</p>
                </div>
            </div>
            <div className="normal_flex">
                <p className={styles.column}>授权链接</p>
                <p>操作</p>
            </div>
            <div className={styles.line} />
            <div className="normal_flex">
                <p className={styles.column} style={{ marginTop: '28px' }}>PC 版授权链接</p>
                <a style={{marginRight: '20px'}} className="a"
                   onClick={() => copyMessage(`${window.location.origin}/#${routes.authorize.path}`)}>复制链接</a>
            </div>
            <div className="normal_flex">
                <p className={styles.column}>H5 版授权链接</p>
                <a style={{marginRight: '20px'}} className="a"
                   onClick={() => copyMessage(`${window.location.origin}/#${routes.authorizeH5.path}`)}>复制链接</a>
            </div>
        </div>
    )
}
