package innerservice

import (
	"net/http"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"

	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/token/wxtoken"
	"github.com/gin-gonic/gin"
)

// GetTicketHandler 获取Ticket
func GetTicketHandler(c *gin.Context) {
	ticket := wxbase.GetTicket()
	if len(ticket) == 0 {
		c.JSON(http.StatusOK, errno.ErrEmptyTicket)
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"ticket": ticket}))
}

// GetComponentAccessTokenHandler 获取ComponentAccessToken
func GetComponentAccessTokenHandler(c *gin.Context) {
	var token string
	var err error
	for i := 0; i < 3; i++ {
		if token, err = wxtoken.GetComponentAccessToken(); err != nil {
			log.Error(err)
			if err.Error() == "lock fail" {
				time.Sleep(200 * time.Millisecond)
				continue
			}
		}
		break
	}
	if err != nil {
		if err.Error() == "empty ticket" {
			c.JSON(http.StatusOK, errno.ErrEmptyTicket)
			return
		}
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"token": token}))
}

// GetAuthorizerAccessTokenHandler 获取AuthorizerAccessToken
func GetAuthorizerAccessTokenHandler(c *gin.Context) {
	var token string
	var err error
	for i := 0; i < 3; i++ {
		if token, err = wxtoken.GetAuthorizerAccessToken(c.Query("appid")); err != nil {
			log.Error(err)
			if err.Error() == "lock fail" {
				time.Sleep(200 * time.Millisecond)
				continue
			}
		}
		break
	}
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"token": token}))
}
