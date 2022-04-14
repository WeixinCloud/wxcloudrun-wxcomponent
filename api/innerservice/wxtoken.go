package innerservice

import (
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx"

	wxbase "github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/base"
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
	token, err := wx.GetComponentAccessToken()
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
	token, err := wx.GetAuthorizerAccessToken(c.Query("appid"))
	if err != nil {
		c.JSON(http.StatusOK, errno.ErrSystemError.WithData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"token": token}))
}
