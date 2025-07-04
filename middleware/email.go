package middleware

import (
	"bytes"
	"io"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	utils2 "github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorToEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var username string
		claims, _ := utils2.GetClaims(c)
		if claims.Username != "" {
			username = claims.Username
		} else {
			id, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			var u system.SysUser
			err := global.GVA_DB.Where("id = ?", id).First(&u).Error
			if err != nil {
				username = "Unknown"
			}
			username = u.Username
		}
		body, _ := io.ReadAll(c.Request.Body)
		// body，ioutil.ReadAllc.Request.Body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
		}
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		status := c.Writer.Status()
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		str := "" + record.Body + "\n" + "" + record.Method + "\n" + "" + record.ErrorMessage + "\n" + "" + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + "" + record.Path + ""
			if err := utils.ErrorToEmail(subject, str); err != nil {
				global.GVA_LOG.Error("ErrorToEmail Failed, err:", zap.Error(err))
			}
		}
	}
}
