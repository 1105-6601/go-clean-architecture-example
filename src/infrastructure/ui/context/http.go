package context

import (
	"github.com/gin-gonic/gin"
)

type HTTP struct {
	*gin.Context
}

func NewHTTP(gin *gin.Context) *HTTP {
	return &HTTP{
		gin,
	}
}

func (context *HTTP) JSONRaw(code int, body string) {
	context.Context.Writer.Header().Set("Content-Type", "application/json")
	context.Context.String(code, body)
}
