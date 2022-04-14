package cloudbasetoken

import (
	"io/ioutil"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/config"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
)

var cloudBaseAccessToken string

const tokenFilePath = "/.tencentcloudbase/wx/cloudbase_access_token"
const defaultAccessToken = "default-tcb-wx-access-token"

func updateCloudBaseAccessToken() {
	content, err := ioutil.ReadFile(tokenFilePath)
	if err != nil {
		log.Errorf(err.Error())
	}
	cloudBaseAccessToken = string(content)
	log.Info("new cloudbase accesstoken: " + cloudBaseAccessToken)
}

func init() {
	if config.WxApiConf.UseCloudBaseAccessToken {
		updateCloudBaseAccessToken()
		go func() {
			if cloudBaseAccessToken == defaultAccessToken {
				timer := time.NewTicker(1 * time.Minute)
				for range timer.C {
					updateCloudBaseAccessToken()
					if cloudBaseAccessToken != defaultAccessToken {
						break
					}
				}
			}
			timer := time.NewTicker(10 * time.Minute)
			for range timer.C {
				updateCloudBaseAccessToken()
			}
		}()
	}
}

// GetCloudBaseAccessToken 获取CloudBaseAccessToken
func GetCloudBaseAccessToken() string {
	return cloudBaseAccessToken
}
