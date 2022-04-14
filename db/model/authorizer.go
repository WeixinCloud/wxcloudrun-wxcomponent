package model

import "time"

// Authorizer 授权账号
type Authorizer struct {
	ID            int       `gorm:"column:id;primaryKey" json:"id"`
	Appid         string    `gorm:"column:appid" json:"appid"`
	AppType       int       `gorm:"column:apptype" json:"appType"`
	ServiceType   int       `gorm:"column:servicetype" json:"serviceType"`
	NickName      string    `gorm:"column:nickname" json:"nickName"`
	UserName      string    `gorm:"column:username" json:"userName"`
	HeadImg       string    `gorm:"column:headimg" json:"headImg"`
	QrcodeUrl     string    `gorm:"column:qrcodeurl" json:"qrcodeUrl"`
	PrincipalName string    `gorm:"column:principalname" json:"principalName"`
	RefreshToken  string    `gorm:"column:refreshtoken" json:"refreshToken"`
	FuncInfo      string    `gorm:"column:funcinfo" json:"funcInfo"`
	VerifyInfo    int       `gorm:"column:verifyinfo" json:"verifyInfo"`
	AuthTime      time.Time `gorm:"column:authtime" json:"authTime"`
}
