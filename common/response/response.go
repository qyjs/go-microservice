package response

//Msg 消息结构
type Msg struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

//Response 返回
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message *Msg        `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
}

const (
	//MsgStatusSuccess success
	MsgStatusSuccess = "SUCCESS"
	//MsgStatusWarn Warn
	MsgStatusWarn = "WARN"
	//MsgStatusError Error
	MsgStatusError = "ERROR"
)

//NewSuccess 正常返回
func NewSuccess(data interface{}) *Response {
	return &Response{
		Data: data,
		Message: &Msg{
			Code: 0,
		},
		Status: MsgStatusSuccess,
	}
}

//NewFail 错误返回
func NewFail(msg string, code int) *Response {
	if code == 0 {
		code = 99
	}
	if msg == "" {
		msg = "error"
	}
	return &Response{
		Message: &Msg{
			Code:    code,
			Message: msg,
		},
		Status: MsgStatusError,
	}
}
