package check

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// 返回带数据成功byte  data=args[0] msg=args[1]
func ReturnMsgCustomSuccessByte(msgId string, args ...any) []byte {
	m := make(map[string]any, 5)
	m["msgId"] = msgId
	m["code"] = 0
	m["data"] = nil
	if len(args) > 0 {
		m["data"] = args[0]
	}
	if len(args) > 1 {
		m["msg"] = args[1]
	}
	byte_data, _ := json.Marshal(m)
	return byte_data
}

// 返回byte
func ReturnMsgSuccessTipByte(args ...string) []byte {
	m := make(map[string]any, 4)
	m["msgId"] = "success"
	m["code"] = 1
	m["msg"] = "无此选项"
	if len(args) > 0 {
		m["msg"] = args[0]
	}
	if len(args) > 1 {
		m["append"] = true
	}
	byte_data, _ := json.Marshal(m)
	return byte_data
}

// 返回byte
func ReturnMsgWarningTipByte(args ...any) []byte {
	m := make(map[string]any, 4)
	m["msgId"] = "warning"
	m["code"] = -2
	m["msg"] = "无此选项"
	if len(args) > 0 {
		m["msg"] = args[0]
	}
	if len(args) > 1 {
		m["append"] = true
	}
	byte_data, _ := json.Marshal(m)
	return byte_data
}

// 返回byte
func ReturnMsgLog(args ...string) []byte {
	m := make(map[string]any, 4)
	m["msgId"] = "append_log"
	m["code"] = 1
	m["msg"] = ""
	if len(args) > 0 {
		m["msg"] = args[0]
	}
	byte_data, _ := json.Marshal(m)
	return byte_data
}

// 返回成功数据
func ErrSuccess(data any) map[string]any {
	return gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	}
}

// 自定义msg 警告
func ErrWarningNull(msg string, args ...any) map[string]any {
	_h := gin.H{
		"code": 0,
		"msg":  msg,
	}
	if len(args) > 0 {
		_h["data"] = args[0]
	}
	return _h
}

// 自定义msg 警告
func ErrParmsWarning(args ...any) map[string]any {
	_h := gin.H{
		"code": 0,
		"msg":  "参数错误!",
	}
	if len(args) > 0 {
		_h["data"] = args[0]
	}
	return _h
}

// 自定义msg 警告
func ErrServerError(args ...string) map[string]any {
	msg := "服务器内部错误"
	if len(args) > 0 {
		msg = args[0]
	}
	return gin.H{
		"code": 0,
		"msg":  msg,
	}
}

// 返回成功且空数据
func ErrSuccessNull(args ...string) map[string]any {
	msg := "success"
	if len(args) > 0 {
		msg = args[0]
	}
	return gin.H{
		"code": 200,
		"msg":  msg,
	}
}
