import React from 'react'
import ReactDOM from 'react-dom'
import Login from './pages/login'
import {HashRouter, Routes, Route} from "react-router-dom";
import Console, { routes } from "./components/Console";
import { Popup } from 'tdesign-react';
import * as Icon from 'tdesign-icons-react'
import './main.less'
import 'tdesign-react/es/style/index.css'; // 少量公共样式
import ThirdToken from "./pages/thirdToken";
import ThirdMessage from "./pages/thirdMessage";
import AuthorizedAccountManage from "./pages/authorizedAccountManage";
import SystemVersion from "./pages/systemVersion";
import PasswordManage from "./pages/passwordManage";
import AuthPageManage from "./pages/authPageManage";
import AuthPage from "./pages/authPage";
import AuthPageH5 from "./pages/authPageH5";
import Home from "./pages/home";
import DevelopTools from "./pages/developTools";

ReactDOM.render(
    <React.StrictMode>
        <HashRouter>
            <Routes>
                <Route path={routes.login.path} element={<Login />} />
                <Route path={routes.authorize.path} element={<AuthPage />} />
                <Route path={routes.authorizeH5.path} element={<AuthPageH5 />} />
                <Route path={"/"} element={<Console />}>
                    <Route path={routes.home.path} element={<Home />} />
                    <Route path={routes.thirdToken.path} element={<ThirdToken />} />
                    <Route path={routes.thirdMessage.path} element={<ThirdMessage />} />
                    <Route path={routes.authorizedAccountManage.path} element={<AuthorizedAccountManage />} />
                    <Route path={routes.systemVersion.path} element={<SystemVersion />} />
                    <Route path={routes.passwordManage.path} element={<PasswordManage />} />
                    <Route path={routes.authPageManage.path} element={<AuthPageManage />} />
                    <Route path={routes.developTools.path} element={<DevelopTools />} />
                </Route>
            </Routes>
        </HashRouter>
        <Popup content={<img style={{ width: '150px', height: '150px', marginTop: '5px' }} src="https://static-index-4gtuqm3bfa95c963-1304825656.tcloudbaseapp.com/cd6125c-c249-4d19-891f-1016ed218a6e.png" alt="" />} placement="left" showArrow destroyOnClose>
            <div style={{ display: 'flex', alignItems: 'center', position: 'fixed', right: '0', top: '40vh', backgroundColor: 'white', flexDirection: 'column', padding: '15px', boxShadow: '-3px 4px 5px 1px rgba(0,0,0,0.2)', borderRadius: '5px 0 0 5px' }}>
                <Icon.ChatIcon />
                <p style={{ margin: '10px 0 0 0' }}>技</p>
                <p style={{ margin: '10px 0 0 0' }}>术</p>
                <p style={{ margin: '10px 0 0 0' }}>支</p>
                <p style={{ margin: '10px 0 0 0' }}>持</p>
            </div>
        </Popup>
    </React.StrictMode>,
    document.getElementById('root')
)
