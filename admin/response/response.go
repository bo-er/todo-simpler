package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义上下文中的键
const (
	prefix = "todo"
	//userID的键
	UserIDKey = prefix + "/user-id"
	//请求体的键
	ReqBodyKey = prefix + "/req-body"
	//返回体的键
	ResBodyKey = prefix + "/res-body"
	//日志请求体的键
	LoggerReqBodyKey = prefix + "/logger-req-body"
)

// Result 后端返回给前端的请求结果结构体
type Result struct {
	Code int         //消息代码
	Msg  string      //消息信息
	Data interface{} //消息体
}

// ResponseError 定义响应错误
type ResponseError struct {
	Code       int    // 错误码
	Message    string // 错误消息
	StatusCode int    // 响应状态码
	ERR        error  // 响应错误
}

// Error ResponseError实现Error接口
func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

// UnWrapErrorToResponsError 如果错误是*ResponseError的指针返回
func UnWrapErrorToResponsError(err error) *ResponseError {
	if v, ok := err.(*ResponseError); ok {
		return v
	}
	return nil
}

// ResSuccess 响应成功消息
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {

	var responseErr *ResponseError

	if err != nil && UnWrapErrorToResponsError(err) != nil {

		responseErr = UnWrapErrorToResponsError(err)

	} else {
		responseErr = UnWrapErrorToResponsError(NewResponse(500, 500, "服务器内部错误"))
	}

	if len(status) > 0 {
		responseErr.StatusCode = status[0]
	}

	if err := responseErr.ERR; err != nil {
		if responseErr.Message == "" {
			responseErr.Message = err.Error()
		}

		if status := responseErr.StatusCode; status >= 400 && status < 500 {

		} else if status >= 500 {

		}
	}
	ResJSON(c, responseErr.StatusCode, responseErr.Message)
}

// ResJSON 请求返回JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	message, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, message)
	if status == 200 {
		message = []byte("请求成功")
	}
	c.JSON(status, Result{Code: status, Msg: string(message), Data: v})
	c.Abort()
}

// NewResponse 创建响应错误
func NewResponse(code, statusCode int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		StatusCode: statusCode,
	}
	return res
}
