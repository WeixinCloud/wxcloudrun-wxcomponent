package admin

// 人员管理

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/utils"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"
	"github.com/gin-gonic/gin"
)

type userReq struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码md5
	OldPassword string `json:"oldPassword"` // 旧密码md5
}

// 更新用户密码
func updateUserNameHandler(c *gin.Context) {
	var req userReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	if req.Username == "" {
		log.Error("param empty ", req)
		c.JSON(http.StatusOK, errno.ErrInvalidParam)
		return
	}
	jwt, _ := c.Get("jwt")
	Id, _ := strconv.Atoi(jwt.(*utils.Claims).ID)
	if err := dao.UpdateUserRecord(int32(Id), req.Username, "", ""); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errno.ErrUserErr.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}

// 更新用户密码
func updateUserPwdHandler(c *gin.Context) {
	var req userReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData(err.Error()))
		return
	}
	if req.OldPassword == "" || req.Password == "" {
		log.Error("param empty ", req)
		c.JSON(http.StatusOK, errno.ErrInvalidParam)
		return
	}
	if req.OldPassword == req.Password {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData("the new password is the same as the old password"))
		return
	}
	ok, err := checkPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	if !ok {
		c.JSON(http.StatusOK, errno.ErrInvalidParam.WithData("invalid password"))
		return
	}
	jwt, _ := c.Get("jwt")
	Id, _ := strconv.Atoi(jwt.(*utils.Claims).ID)
	if err := dao.UpdateUserRecord(int32(Id), req.Username, req.Password, req.OldPassword); err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusOK, errno.ErrUserErr.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK)
}

func checkPassword(pwd string) (bool, error) {
	return regexp.MatchString(`^\w{32}$`, pwd)
}
