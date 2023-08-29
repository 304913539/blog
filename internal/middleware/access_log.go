package middleware

import (
	"blog-service/global"
	"blog-service/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type AccessLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWrite) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWrite{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		body, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  string(body),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
