package wxcallback

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"strconv"
	"time"
)

var EventTicket = "component_verify_ticket"    //ticket推送
var EventUnauthorized = "unauthorized"         //取消授权
var EventUpdateAuthorized = "updateauthorized" //更新授权
var EventAuthorized = "authorized"             //授权成功

var MsgTypeText = "text"   //文本消息
var MsgTypeImage = "image" //文本消息
var MsgTypeVoice = "voice" //语音消息
var MsgTypeVideo = "Video" //视频消息
var MsgTypeMusic = "music" //音乐消息
var MsgTypeNews = "news"   //图文消息

// EventMessageBody 事件推送
type EventMessageBody struct {
	XMLName                      xml.Name `xml:"xml"`
	AppId                        string   `xml:"AppId" json:"app_id"`
	CreateTime                   int      `xml:"CreateTime" json:"create_time"`
	InfoType                     string   `xml:"InfoType" json:"info_type"`
	ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket" json:"component_verify_ticket"`
	AuthorizerAppid              string   `xml:"AuthorizerAppid" json:"authorizer_appid"`
	AuthorizationCode            string   `xml:"AuthorizationCode" json:"authorization_code"`
	AuthorizationCodeExpiredTime string   `xml:"AuthorizationCodeExpiredTime" json:"authorization_code_expired_time"`
	PreAuthCode                  string   `xml:"PreAuthCode" json:"pre_auth_code"`
}

// MessageBodyDecrypt 消息体
type MessageBodyDecrypt struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Url          string   `xml:"Url"`
	PicUrl       string   `xml:"PicUrl"`
	MediaId      string   `xml:"MediaId"`
	ThumbMediaId string   `xml:"ThumbMediaId"`
	Content      string   `xml:"Content"`
	MsgId        int      `xml:"MsgId"`
	Location_X   string   `xml:"Location_x"`
	Location_Y   string   `xml:"Location_y"`
	Label        string   `xml:"Label"`
}

type MessageEncryptBody struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      CDATA    `xml:"Encrypt"`
	MsgSignature CDATA    `xml:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce        CDATA    `xml:"Nonce"`
}

type MessageText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Content      CDATA    `xml:"Content"`
}

type MessageImage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Image        Media    `xml:"Image"`
}

type MessageVoice struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Voice        Media    `xml:"Voice"`
}

type MessageVideo struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Video        Video    `xml:"Video"`
}

type MessageMusic struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Music        Music    `xml:"Music"`
}

type MessageArticle struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   CDATA         `xml:"ToUserName"`
	FromUserName CDATA         `xml:"FromUserName"`
	CreateTime   string        `xml:"CreateTime"`
	MsgType      CDATA         `xml:"MsgType"`
	ArticleCount string        `xml:"ArticleCount"`
	Articles     []ArticleItem `xml:"Articles"`
}

type CDATA struct {
	Text string `xml:",innerxml"`
}

type Media struct {
	MediaId CDATA `xml:"MediaId"`
}

type Video struct {
	MediaId     CDATA `xml:"MediaId"`
	Title       CDATA `xml:"Title"`
	Description CDATA `xml:"Description"`
}

type Music struct {
	Title        CDATA `xml:"Title"`
	Description  CDATA `xml:"Description"`
	MusicUrl     CDATA `xml:"MusicUrl"`
	HQMusicUrl   CDATA `xml:"HQMusicUrl"`
	ThumbMediaId CDATA `xml:"ThumbMediaId"`
}

type ArticleItem struct {
	Title       CDATA `xml:"Title"`
	Description CDATA `xml:"Description"`
	PicUrl      CDATA `xml:"PicUrl"`
	Url         CDATA `xml:"Url"`
}

func FormatMessage(plainText []byte, data interface{}) (*interface{}, error) {
	length := GetMessageLength(plainText)
	err := xml.Unmarshal(plainText[20:20+length], data)
	if err != nil {
		return nil, errors.New("格式化消息失败：format message error")
	}
	return &data, nil
}

func GetMessageLength(plainText []byte) int32 {
	// Read length
	buf := bytes.NewBuffer(plainText[16:20])
	var length int32
	err := binary.Read(buf, binary.BigEndian, &length)
	if err != nil {
		panic("获取消息长度失败：read message length error")
	}
	return length
}

func ValueToCDATA(content string) CDATA {
	return CDATA{"<![CDATA[" + content + "]]>"}
}

// FormatTextMessage 格式化文本消息
func FormatTextMessage(fromUserName string, toUserName string, content string) []byte {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	textMessage := MessageText{FromUserName: ValueToCDATA(fromUserName), ToUserName: ValueToCDATA(toUserName), Content: ValueToCDATA(content), CreateTime: timestamp, MsgType: ValueToCDATA(MsgTypeText)}
	messageBytes, err := xml.MarshalIndent(textMessage, " ", "  ")
	if err != nil {
		panic("格式化文本消息失败：xml marsha1 error；" + err.Error())
	}
	return messageBytes
}

// FormatImageMessage 格式化图片消息
func FormatImageMessage(fromUserName string, toUserName string, mediaId string) []byte {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	imageMessage := MessageImage{FromUserName: ValueToCDATA(fromUserName), ToUserName: ValueToCDATA(toUserName), CreateTime: timestamp, MsgType: ValueToCDATA(MsgTypeImage), Image: Media{ValueToCDATA(mediaId)}}
	messageBytes, err := xml.MarshalIndent(imageMessage, " ", "  ")
	if err != nil {
		panic("格式化图片消息失败：xml marsha1 error；" + err.Error())
	}
	return messageBytes
}

// FormatVoiceMessage 格式语音消息
func FormatVoiceMessage(fromUserName string, toUserName string, mediaId string) []byte {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	voiceMessage := MessageVoice{FromUserName: ValueToCDATA(fromUserName), ToUserName: ValueToCDATA(toUserName), CreateTime: timestamp, MsgType: ValueToCDATA(MsgTypeVoice), Voice: Media{ValueToCDATA(mediaId)}}
	messageBytes, err := xml.MarshalIndent(voiceMessage, " ", "  ")
	if err != nil {
		panic("格式化语音消息失败：xml marsha1 error；" + err.Error())
	}
	return messageBytes
}

// FormatVideoMessage 格式化视频消息
func FormatVideoMessage(fromUserName string, toUserName string, mediaId string, title string, description string) []byte {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	videoMessage := MessageVideo{FromUserName: ValueToCDATA(fromUserName), ToUserName: ValueToCDATA(toUserName), CreateTime: timestamp, MsgType: ValueToCDATA(MsgTypeVideo), Video: Video{
		MediaId:     ValueToCDATA(mediaId),
		Title:       ValueToCDATA(title),
		Description: ValueToCDATA(description),
	}}
	messageBytes, err := xml.MarshalIndent(videoMessage, " ", "  ")
	if err != nil {
		panic("格式化语音消息失败：xml marsha1 error；" + err.Error())
	}
	return messageBytes
}

// FormatMusicMessage 格式化音乐消息
func FormatMusicMessage(fromUserName string, toUserName string, thumbMediaId string, title string, description string, musicUrl string, hQMusicUrl string) []byte {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	musicMessage := MessageMusic{FromUserName: ValueToCDATA(fromUserName), ToUserName: ValueToCDATA(toUserName), CreateTime: timestamp, MsgType: ValueToCDATA(MsgTypeMusic), Music: Music{
		Title:        ValueToCDATA(title),
		Description:  ValueToCDATA(description),
		MusicUrl:     ValueToCDATA(musicUrl),
		HQMusicUrl:   ValueToCDATA(hQMusicUrl),
		ThumbMediaId: ValueToCDATA(thumbMediaId),
	}}
	messageBytes, err := xml.MarshalIndent(musicMessage, " ", "  ")
	if err != nil {
		panic("格式化音乐消息失败：xml marsha1 error；" + err.Error())
	}
	return messageBytes
}

// FormatArticlesMessage 格式化音乐消息
func FormatArticlesMessage(fromUserName string, toUserName string, items []ArticleItem) []byte {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	articleNum := strconv.Itoa(len(items))
	musicMessage := MessageArticle{FromUserName: ValueToCDATA(fromUserName), ToUserName: ValueToCDATA(toUserName), CreateTime: timestamp, MsgType: ValueToCDATA(MsgTypeNews), Articles: items, ArticleCount: articleNum}
	messageBytes, err := xml.MarshalIndent(musicMessage, " ", "  ")
	if err != nil {
		panic("格式化图文消息失败：xml marsha1 error；" + err.Error())
	}
	return messageBytes
}

// FormatEncryptData 格式化微信加密数据
func FormatEncryptData(encrypted string, token string) MessageEncryptBody {
	nonce := makeRandomString(16)
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	data := MessageEncryptBody{Encrypt: ValueToCDATA(encrypted), Nonce: ValueToCDATA(nonce), MsgSignature: ValueToCDATA(GetSignature(timestamp, nonce, encrypted, token)), TimeStamp: timestamp}
	return data
}
