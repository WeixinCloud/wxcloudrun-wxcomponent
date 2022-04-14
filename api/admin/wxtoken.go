package admin

import (
	"net/http"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/wx/cloudbasetoken"
	"github.com/gin-gonic/gin"
)

func getCloudbaseAccessTokenHandler(c *gin.Context) {
	c.JSON(http.StatusOK, errno.OK.WithData(gin.H{"token": cloudbasetoken.GetCloudBaseAccessToken()}))
}
