package errno

// JsonResult 返回结果，json格式
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg"`
	Data     interface{} `json:"data,omitempty"`
}

// Result 返回结果
type Result interface {
	WithData(data interface{}) Result
}

// WithData 填充data内容
func (e *JsonResult) WithData(data interface{}) Result {
	a := *e
	a.Data = data
	return &a
}
