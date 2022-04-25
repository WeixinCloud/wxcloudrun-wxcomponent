import styles from './index.module.less'
import { Alert } from 'tdesign-react'
import home01 from '../../assets/icons/home01.png'
import home02 from '../../assets/icons/home02.png'
import home03 from '../../assets/icons/home03.png'
import home04 from '../../assets/icons/home04.png'
import home05 from '../../assets/icons/home05.png'
import home06 from '../../assets/icons/home06.png'
import {routes} from "../../config/route";

export default function Home() {

    return (
        <div className={styles.home}>
            <Alert theme="info" icon={<span />} message="欢迎体验基于微信第三方平台和微信云托管平台为基础搭建的“服务商微管家”SaaS应用。如有更多的需求或者使用问题可加入官方群进行反馈。" />
            <div className={styles.home_header}>
                <div className={styles.home_header_box}>
                    <p className={styles.home_title}>产品体验指引</p>

                    <div className={styles.home_header_step}>

                        <div className={styles.home_header_step_item}>
                            <p className={styles.home_header_step_item_number}>1</p>
                            <p className={styles.home_header_step_item_title}>获取商家授权</p>
                            <p className={`${styles.home_header_step_item_desc} desc`}>服务商获得商家小程序授权后，即可代商家开发和运营小程序。<a href={`#${routes.authPageManage.path}`} className="a">前往体验</a></p>
                        </div>

                        <div className={styles.home_header_step_line}>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</div>

                        <div className={styles.home_header_step_item}>
                            <p className={styles.home_header_step_item_number}>2</p>
                            <p className={styles.home_header_step_item_title}>管理授权帐号</p>
                            <p className={`${styles.home_header_step_item_desc} desc`}>可批量管理已授权商家公众号和小程序，可对授权帐号进行批量操作。<a href={`#${routes.authorizedAccountManage.path}`} className="a">前往体验</a></p>
                        </div>

                        <div className={styles.home_header_step_line}>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</div>

                        <div className={styles.home_header_step_item}>
                            <p className={styles.home_header_step_item_number}>3</p>
                            <p className={styles.home_header_step_item_title}>二次开发与系统对接</p>
                            <p className={`${styles.home_header_step_item_desc} desc`}>该应用支持通过修改源码进行二次开发，支持通过openApi与业务系统对接。</p>
                        </div>

                    </div>
                </div>
                <div className={styles.home_header_box}>
                    <p className={styles.home_title}>联系我们</p>

                    <div className={styles.home_header_contact_detail}>
                        <img className={styles.home_header_contact_detail_img} src="https://static-index-4gtuqm3bfa95c963-1304825656.tcloudbaseapp.com/cd6125c-c249-4d19-891f-1016ed218a6e.png" alt="" />
                        <p className={`${styles.home_header_contact_detail_text} desc`} style={{ marginTop: '10px' }}>加入官方交流群，获得技术支持</p>
                    </div>
                </div>
            </div>

            <div className={styles.home_characteristic}>
                <p className={styles.home_title}>产品特性</p>
                <div className={styles.home_characteristic_list}>
                    <div className={styles.home_characteristic_list_item}>
                        <img className={styles.home_characteristic_list_item_icon} src={home01} alt="" />
                        <div className={styles.home_characteristic_list_item_box}>
                            <p className={styles.home_characteristic_list_item_box_title}>开箱即用</p>
                            <p className={`${styles.home_characteristic_list_item_box_desc} desc`}>内置开箱即用的SaaS应用，便于服务商快速获得商家小程序授权、0成本启动服务商业务。</p>
                        </div>
                    </div>

                    <div className={styles.home_characteristic_list_item}>
                        <img className={styles.home_characteristic_list_item_icon} src={home02} alt="" />
                        <div className={styles.home_characteristic_list_item_box}>
                            <p className={styles.home_characteristic_list_item_box_title}>最佳实践</p>
                            <p className={`${styles.home_characteristic_list_item_box_desc} desc`}>内置应用是官方基于最佳实践进行设计，帮助新手服务商快速掌握基于第三方平台开展业务。</p>
                        </div>
                    </div>

                    <div className={styles.home_characteristic_list_item}>
                        <img className={styles.home_characteristic_list_item_icon} src={home03} alt="" />
                        <div className={styles.home_characteristic_list_item_box}>
                            <p className={styles.home_characteristic_list_item_box_title}>可轻松扩展</p>
                            <p className={`${styles.home_characteristic_list_item_box_desc} desc`}>基于云原生设计理念，且项目开源，开发者可快速低成本进行二次开发以及与业务系统集成。</p>
                        </div>
                    </div>

                    <div className={styles.home_characteristic_list_item}>
                        <img className={styles.home_characteristic_list_item_icon} src={home04} alt="" />
                        <div className={styles.home_characteristic_list_item_box}>
                            <p className={styles.home_characteristic_list_item_box_title}>节省成本</p>
                            <p className={`${styles.home_characteristic_list_item_box_desc} desc`}>根据负载自动扩缩容，超细粒度资源控制，以及按实际使用量计费，节省成本。</p>
                        </div>
                    </div>

                    <div className={styles.home_characteristic_list_item}>
                        <img className={styles.home_characteristic_list_item_icon} src={home05} alt="" />
                        <div className={styles.home_characteristic_list_item_box}>
                            <p className={styles.home_characteristic_list_item_box_title}>安全稳定</p>
                            <p className={`${styles.home_characteristic_list_item_box_desc} desc`}>基于VPC 网络隔离，提供应用运行时安全保障；支持微信私有协议，可免鉴权安全调用微信接口。</p>
                        </div>
                    </div>

                    <div className={styles.home_characteristic_list_item}>
                        <img className={styles.home_characteristic_list_item_icon} src={home06} alt="" />
                        <div className={styles.home_characteristic_list_item_box}>
                            <p className={styles.home_characteristic_list_item_box_title}>简易运维</p>
                            <p className={`${styles.home_characteristic_list_item_box_desc} desc`}>服务部署于微信云托管，微信云托管提供丰富的运维管理工具，极大降低开发者维护成本。</p>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    )
}
