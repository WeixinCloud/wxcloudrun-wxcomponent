package errno

var (
	OK = &JsonResult{Code: 0, ErrorMsg: "OK"}

	ErrNotAuthorized      = &JsonResult{Code: 1000, ErrorMsg: "访问未授权"}
	ErrInvalidParam       = &JsonResult{Code: 1001, ErrorMsg: "参数格式有误"}
	ErrSystemError        = &JsonResult{Code: 1002, ErrorMsg: "系统错误"}
	ErrAuthTimeout        = &JsonResult{Code: 1003, ErrorMsg: "登录超时"}
	ErrAuthTokenErr       = &JsonResult{Code: 1004, ErrorMsg: "Token错误"}
	ErrUserErr            = &JsonResult{Code: 1005, ErrorMsg: "用户更新错误"}
	ErrAuthErr            = &JsonResult{Code: 1006, ErrorMsg: "登录失败"}
	ErrEmptyTicket        = &JsonResult{Code: 1007, ErrorMsg: "Ticket为空"}
	ErrInvalidStatus      = &JsonResult{Code: 1008, ErrorMsg: "状态异常"}
	ErrInvalidType        = &JsonResult{Code: 1009, ErrorMsg: "类型错误"}
	ErrRequestErr         = &JsonResult{Code: 1010, ErrorMsg: "请求错误"}
	ErrAuthErrExceedLimit = &JsonResult{Code: 1011, ErrorMsg: "登录失败次数超过限制"}
)
