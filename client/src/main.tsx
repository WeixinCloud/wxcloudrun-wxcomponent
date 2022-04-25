import React, {useEffect} from 'react'
import ReactDOM from 'react-dom'
import {HashRouter, Routes, Route, useNavigate, Outlet} from "react-router-dom";
import Console from "./components/Console";
import { Popup } from 'tdesign-react';
import * as Icon from 'tdesign-icons-react'
import './main.less'
import 'tdesign-react/es/style/index.css'; // 少量公共样式
import { routes } from './config/route'
import {initNav} from "./utils/login";

const InitNav = () => {
    const navigate = useNavigate()

    useEffect(() => {
       initNav(navigate)
    }, [])

    return <Outlet />
}

ReactDOM.render(
    <React.StrictMode>
        <HashRouter>
            <Routes>
                <Route element={<InitNav />}>
                    {
                        Object.values(routes).filter(i => i.dontNeedMenu).map(i => {
                            return (
                                // @ts-ignore
                                <Route key={`route_${i.path}`} path={i.path} element={i.element} />
                            )
                        })
                    }
                    <Route path={"/"} element={<Console />}>
                        {
                            Object.values(routes).filter(i => !i.dontNeedMenu).map(i => {
                                return (
                                    // @ts-ignore
                                    <Route key={`route_${i.path}`} path={i.path} element={i.element} />
                                )
                            })
                        }
                    </Route>
                </Route>
            </Routes>
        </HashRouter>
        <div style={{ backgroundColor: '#f6f7f8' }}>
            <Popup content={<img style={{ width: '150px', height: '150px', marginTop: '5px' }} src="https://static-index-4gtuqm3bfa95c963-1304825656.tcloudbaseapp.com/cd6125c-c249-4d19-891f-1016ed218a6e.png" alt="" />} placement="left" showArrow destroyOnClose>
                <div style={{ display: 'flex', alignItems: 'center', position: 'fixed', right: '0', top: '40vh', backgroundColor: 'white', flexDirection: 'column', padding: '15px', boxShadow: '-3px 4px 5px 1px rgba(0,0,0,0.2)', borderRadius: '5px 0 0 5px' }}>
                    <Icon.ChatIcon />
                    <p style={{ margin: '10px 0 0 0' }}>技</p>
                    <p style={{ margin: '10px 0 0 0' }}>术</p>
                    <p style={{ margin: '10px 0 0 0' }}>支</p>
                    <p style={{ margin: '10px 0 0 0' }}>持</p>
                </div>
            </Popup>
        </div>
    </React.StrictMode>,
    document.getElementById('root')
)
