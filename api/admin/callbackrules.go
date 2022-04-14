package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"github.com/gin-gonic/gin"
)

type getCallBackProxyRuleListReq struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
	Type   int `form:"type"`
}

type callBackProxyRule struct {
	ID         int32                 `json:"id"`
	Name       string                `json:"name"`
	InfoType   string                `json:"infoType"`
	MsgType    string                `json:"msgType"`
	Event      string                `json:"event"`
	Open       int                   `json:"open"`
	Data       model.HttpProxyConfig `json:"data"`
	CreateTime int64                 `json:"createTime"`
	UpdateTime int64                 `json:"updateTime"`
}

func getCallBackProxyRuleListHandler(c *gin.Context) {
	var req getCallBackProxyRuleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	dbValue, total, err := dao.GetWxCallBackRuleList(req.Offset, req.Limit, req.Type)
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	res := make([]callBackProxyRule, 0, 10)
	for _, v := range dbValue {
		var proxyConfig model.HttpProxyConfig
		if err = json.Unmarshal([]byte(v.Info), &proxyConfig); err != nil {
			log.Errorf("Unmarshal err, %v", err)
		} else {
			res = append(res, callBackProxyRule{
				ID:         v.ID,
				Name:       v.Name,
				InfoType:   v.InfoType,
				MsgType:    v.MsgType,
				Event:      v.Event,
				Open:       v.Open,
				Data:       proxyConfig,
				CreateTime: v.CreateTime.Unix(),
				UpdateTime: v.UpdateTime.Unix(),
			})
		}
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{
		"total": total,
		"rules": res,
	}))
}

func updateCallBackProxyRuleHandler(c *gin.Context) {
	var req callBackProxyRule
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	if req.InfoType == "" && req.MsgType == "" && req.Event == "" {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData("消息推送类型为空"))
		return
	}
	value, _ := json.Marshal(req.Data)
	if err := dao.UpdateWxCallBackRule(&model.WxCallbackRule{
		ID:       req.ID,
		Name:     req.Name,
		InfoType: req.InfoType,
		MsgType:  req.MsgType,
		Event:    req.Event,
		Open:     req.Open,
		Type:     model.PROXYTYPE_HTTP,
		Info:     string(value),
	}); err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}

func addCallBackProxyRuleHandler(c *gin.Context) {
	var req callBackProxyRule
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	if req.InfoType == "" && req.MsgType == "" && req.Event == "" {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData("消息推送类型为空"))
		return
	}
	value, _ := json.Marshal(req.Data)
	if err := dao.AddWxCallBackRule(&model.WxCallbackRule{
		Name:     req.Name,
		InfoType: req.InfoType,
		MsgType:  req.MsgType,
		Event:    req.Event,
		Open:     req.Open,
		Type:     model.PROXYTYPE_HTTP,
		Info:     string(value),
	}); err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData("该事件已存在转发规则"))
			return
		}
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}

type callBackProxyRuleId struct {
	ID int32 `form:"id"`
}

func delCallBackProxyRuleHandler(c *gin.Context) {
	var req callBackProxyRuleId
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	if err := dao.DelWxCallBackRule(req.ID); err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}

func testCallbackRuleHandler(c *gin.Context) {
	var req callBackProxyRuleId
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	if record, err := dao.GetWxCallBackRuleById(req.ID); err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	} else {
		if record.Open != 0 && record.Type == model.PROXYTYPE_HTTP {
			var proxyConfig model.HttpProxyConfig
			if err = json.Unmarshal([]byte(record.Info), &proxyConfig); err != nil {
				log.Errorf("Unmarshal err, %v", err)
				c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
				return
			}
			resp, err := httputils.PostJson(fmt.Sprintf("http://127.0.0.1:%d%s", proxyConfig.Port,
				strings.Replace(proxyConfig.Path, "$APPID$", "wxtestappid", -1)),
				genWxCallBackReq(record))
			if err != nil {
				log.Error(err)
				c.JSON(http.StatusOK, errno.ErrRequestErr.WithData(err.Error()))
				return
			}
			c.JSON(http.StatusOK, errno.OK.WithData(string(resp)))
			return
		} else {
			c.JSON(http.StatusOK, errno.ErrInvalidStatus.WithData("该规则未启用或类型异常"))
			return
		}
	}
}

type wxCallBackReq struct {
	CreateTime   int64  `json:"CreateTime"`
	ToUserName   string `json:"ToUserName,omitempty"`
	FromUserName string `json:"FromUserName,omitempty"`
	InfoType     string `json:"InfoType,omitempty"`
	MsgType      string `json:"MsgType,omitempty"`
	Event        string `json:"Event,omitempty"`
	Data         string `json:"Data,omitempty"`
}

func genWxCallBackReq(rule *model.WxCallbackRule) *wxCallBackReq {
	if rule.InfoType != "" {
		return &wxCallBackReq{
			CreateTime: time.Now().Unix(),
			InfoType:   rule.InfoType,
			Data:       "TestData",
		}
	} else {
		return &wxCallBackReq{
			CreateTime:   time.Now().Unix(),
			MsgType:      rule.MsgType,
			Event:        rule.Event,
			ToUserName:   "TestUserName1",
			FromUserName: "TestUserName2",
			Data:         "TestData",
		}
	}
}
