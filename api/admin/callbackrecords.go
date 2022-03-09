package admin

import (
	"net/http"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/gin-gonic/gin"
)

type getWxComponentRecordsReq struct {
	StartTime int64  `form:"startTime"`
	EndTime   int64  `form:"endTime"`
	InfoType  string `form:"infoType"`
	Offset    int    `form:"offset"`
	Limit     int    `form:"limit"`
}

type getWxBizRecordsReq struct {
	StartTime int64  `form:"startTime"`
	EndTime   int64  `form:"endTime"`
	Appid     string `form:"appid"`
	MsgType   string `form:"msgType"`
	Event     string `form:"event"`
	Offset    int    `form:"offset"`
	Limit     int    `form:"limit"`
}

func getWxComponentRecordsHandler(c *gin.Context) {
	var req getWxComponentRecordsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	var endTime time.Time
	if req.EndTime == 0 {
		endTime = time.Now()
	} else {
		endTime = time.Unix(req.EndTime, 0)
	}
	records, total, err := dao.GetComponentCallBackRecordList(time.Unix(req.StartTime, 0),
		endTime, req.InfoType, req.Offset, req.Limit)
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"total": total, "records": records}))
}

func getWxBizRecordsHandler(c *gin.Context) {
	var req getWxBizRecordsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	var endTime time.Time
	if req.EndTime == 0 {
		endTime = time.Now()
	} else {
		endTime = time.Unix(req.EndTime, 0)
	}
	records, total, err := dao.GetBizCallBackRecordList(time.Unix(req.StartTime, 0),
		endTime, req.Appid, req.MsgType, req.Event, req.Offset, req.Limit)
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"total": total, "records": records}))
}

func getWxCallBackConfigHandler(c *gin.Context) {
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{
		"envId":         wxbase.GetEnvId(),
		"service":       wxbase.GetService(),
		"componentPath": "/wxcallback/component",
		"bizPath":       "/wxcallback/biz/$APPID$",
		"textMode":      "json",
	}))
}
