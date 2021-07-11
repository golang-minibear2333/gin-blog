package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/pkg/logger"
)

// AccessLogWriter 实现 http.ResponseWriter
//type ResponseWriter interface {
//	Header() http.Header
//	Write([]byte) (int, error)
//	WriteHeader(statusCode int)
//}
type AccessLogWriter struct {
	// 这是继承实现的语法
	gin.ResponseWriter // 实现了 WriteHeader
	body               *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	// 将返回结果写入body
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	// 将返回结果写入输出流
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		// 相当于做了一个切面
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		// method：当前的调用方法。
		// request：当前的请求参数。
		// response：当前的请求结果响应主体。
		// status_code：当前的响应结果状态码。
		// begin_time/end_time：调用方法的开始时间，调用方法结束的结束时间。
		s := "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
