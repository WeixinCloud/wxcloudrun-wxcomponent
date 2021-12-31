package admin

import (
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"

	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/token/cloudbasetoken"
	token "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/token/commtoken"
	"github.com/gin-gonic/gin"
)

func ticketHandler(c *gin.Context) {
	ticket := wxbase.GetTicket()
	if len(ticket) == 0 {
		c.JSON(http.StatusOK, errno.ErrEmptyTicket)
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"ticket": ticket}))
}

func cloudbaseAccessTokenHandler(c *gin.Context) {
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"token": cloudbasetoken.GetCloudBaseAccessToken()}))
}

func componentAccessTokenHandler(c *gin.Context) {
	token, err := token.GetNewComponentAccessToken()
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

func authorizerAccessTokenHandler(c *gin.Context) {
	token, err := token.GetNewAuthorizerAccessToken(c.Query("appid"))
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"token": token}))
}
