package admin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/httputils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx"
	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
	"github.com/gin-gonic/gin"
)

type getAuthorizerListReq struct {
	ComponentAppid string `json:"component_appid"`
	Offset         int    `json:"offset"`
	Count          int    `json:"count"`
}

type authorizerInfo struct {
	AuthorizerAppid string `json:"authorizer_appid"`
	RefreshToken    string `json:"refresh_token"`
	AuthTime        int64  `json:"auth_time"`
}
type getAuthorizerListResp struct {
	TotalCount int              `json:"total_count"`
	List       []authorizerInfo `json:list`
}

func pullAuthorizerListHandler(c *gin.Context) {
	go func() {
		count := 100
		offset := 0
		total := 0
		now := time.Now()
		for {
			var resp getAuthorizerListResp
			if err := getAuthorizerList(offset, count, &resp); err != nil {
				log.Error(err)
				return
			}
			if total == 0 {
				total = resp.TotalCount
			}
			// 插入数据库
			length := len(resp.List)
			records := make([]model.Authorizer, length)
			var wg sync.WaitGroup
			wg.Add(length)
			for i, info := range resp.List {
				go constructAuthorizerRecord(info, &records[i], &wg)
			}
			wg.Wait()
			dao.BatchCreateOrUpdateAuthorizerRecord(&records)

			if length < count {
				break
			}
			offset += count
		}

		// 删除记录
		if err := dao.ClearAuthorizerRecordsBefore(now); err != nil {
			log.Error(err)
			return
		}
	}()
	c.JSON(http.StatusOK, errno.OK)
}

func constructAuthorizerRecord(info authorizerInfo, record *model.Authorizer, wg *sync.WaitGroup) error {
	defer wg.Done()
	record.Appid = info.AuthorizerAppid
	record.AuthTime = time.Unix(info.AuthTime, 0)
	record.RefreshToken = info.RefreshToken
	var appinfo wx.AuthorizerInfoResp

	if err := wx.GetAuthorizerInfo(record.Appid, &appinfo); err != nil {
		log.Errorf("GetAuthorizerInfo fail %v", err)
		return err
	}
	record.AppType = appinfo.AuthorizerInfo.AppType
	record.ServiceType = appinfo.AuthorizerInfo.ServiceType.Id
	record.NickName = appinfo.AuthorizerInfo.NickName
	record.UserName = appinfo.AuthorizerInfo.UserName
	record.HeadImg = appinfo.AuthorizerInfo.HeadImg
	record.QrcodeUrl = appinfo.AuthorizerInfo.QrcodeUrl
	record.PrincipalName = appinfo.AuthorizerInfo.PrincipalName
	record.FuncInfo = appinfo.AuthorizationInfo.StrFuncInfo
	record.VerifyInfo = appinfo.AuthorizerInfo.VerifyInfo.Id
	return nil
}

func getAuthorizerList(offset, count int, resp *getAuthorizerListResp) error {
	req := getAuthorizerListReq{
		ComponentAppid: wxbase.GetAppid(),
		Offset:         offset,
		Count:          count,
	}
	_, respbody, err := httputils.PostWxJson("/cgi-bin/component/api_get_authorizer_list", req, true)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respbody, &resp); err != nil {
		log.Errorf("Unmarshal err, %v", err)
		return err
	}
	return nil
}

func getAuthorizerListHandler(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	appid := c.DefaultQuery("appid", "")
	records, total, err := dao.GetAuthorizerRecords(appid, offset, limit)
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"total": total, "records": records}))
}
