package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/errno"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/db/dao"

	"github.com/gin-gonic/gin"
)

// ProxyConfig 代理配置
type ProxyConfig struct {
	Open bool   `json:"open"`
	Port int    `json:"port"`
	Path string `json:"path,omitempty"`
	Url  string `json:"url"`
}

var proxyConfig ProxyConfig
var target *url.URL

// ProxyHandler 代理处理器
func ProxyHandler(c *gin.Context) {
	if proxyConfig.Open && target != nil {
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			log.Errorf("http: proxy error: %v", err)
			result, _ := json.Marshal(errno.ErrSystemError.WithData(err.Error()))
			rw.Header().Set("Content-Type", "application/json")
			rw.Write([]byte(result))
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	} else {
		c.Status(http.StatusNotFound)
	}
}

// GetProxyConfig 获取代理配置
func GetProxyConfig() ProxyConfig {
	return proxyConfig
}

// SetProxyConfig 设置代理配置
func SetProxyConfig(open bool, port int, path string) error {
	proxyConfig = ProxyConfig{
		Open: open,
		Port: port,
		Path: path,
		Url:  fmt.Sprintf("http://127.0.0.1:%d%s", port, path),
	}
	if open {
		var err error
		if target, err = url.Parse(proxyConfig.Url); err != nil {
			log.Errorf("url Parse error: %v", err)
			return err
		}
	} else {
		target = nil
	}
	if err := setProxyConfigToKv(&proxyConfig); err != nil {
		log.Errorf("url setProxyConfigToKv error: %v", err)
		return err
	}
	return nil
}

func getProxyConfigFromKv(proxyConfig *ProxyConfig) {
	value := dao.GetCommKv("proxy", "")
	if err := json.Unmarshal([]byte(value), proxyConfig); err != nil {
		log.Errorf("getProxyConfigFromKv fail err %s value %s", err.Error(), value)
		proxyConfig.Open = false
		proxyConfig.Port = 0
		proxyConfig.Path = ""
		proxyConfig.Url = ""
	}
	log.Infof("getProxyConfigFromKv %v", *proxyConfig)
}

func setProxyConfigToKv(proxyConfig *ProxyConfig) error {
	value, _ := json.Marshal(*proxyConfig)
	log.Infof("setProxyConfigToKv %v", *proxyConfig)
	return dao.SetCommKv("proxy", string(value))
}

// Init 初始化
func Init() error {
	getProxyConfigFromKv(&proxyConfig)
	if proxyConfig.Open {
		var err error
		if target, err = url.Parse(proxyConfig.Url); err != nil {
			log.Errorf("url Parse error: %v", err)
			proxyConfig.Open = false
			target = nil
		}
	} else {
		target = nil
	}
	return nil
}
