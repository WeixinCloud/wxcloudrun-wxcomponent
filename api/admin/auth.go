package admin

// 系统鉴权，登录，人员管理

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/gin-gonic/gin"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/utils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/model"
)

const ErrHourlyLimit = 10

func checkAuth(req model.UserRecord) (int32, error) {
	record, err := dao.GetUserRecord(req.Username, req.Password)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	if len(record) > 0 {
		return record[0].ID, nil
	}
	return 0, err
}

func authHandler(c *gin.Context) {
	var req model.UserRecord
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("Auth Ip: ", ip)
	key := fmt.Sprintf("AUTH_%s_%d", ip, time.Now().Hour())
	current, err := dao.GetCurrent(key)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	log.Info("current: ", current)
	if current >= ErrHourlyLimit {
		c.JSON(http.StatusOK, errno.ErrAuthErrExceedLimit)
		return
	}

	ID, err := checkAuth(req)
	if err != nil {
		log.Error(err.Error())
		if err := dao.AddOne(key, ErrHourlyLimit); err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errno.ErrAuthErrExceedLimit)
			return
		}
		c.JSON(http.StatusOK, errno.ErrAuthErr.WithData(err.Error()))
		return
	}

	token, _err := utils.GenerateToken(strconv.Itoa(int(ID)), req.Username)
	if _err != nil {
		log.Error(_err.Error())
		c.JSON(http.StatusOK, errno.ErrAuthErr.WithData(_err.Error()))
		return
	}

	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"jwt": token}))

}

func refreshAuthHandler(c *gin.Context) {
	jwt, _ := c.Get("jwt")
	id, userName := jwt.(*utils.Claims).ID, jwt.(*utils.Claims).UserName
	log.Debugf("id: %s name: %s", id, userName)
	token, err := utils.GenerateToken(id, userName)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errno.ErrAuthErr.WithData(err.Error()))
		return
	}

	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"jwt": token}))
}
