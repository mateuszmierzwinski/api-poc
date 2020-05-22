package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type HeadersHandler struct {
	version string
}

func (h *HeadersHandler) Middleware(context *gin.Context) {
	headers := map[string]string {
		"Access-Control-Allow-Origin": "*",
		"Access-Control-Allow-Methods": "GET, HEAD",
		"Access-Control-Allow-Headers": "Origin, User-ID, User-Answer, Content-Type",
		"Expire": time.Now().Add(-1*time.Hour*24*365).Format(time.RFC850),
		"Pragma": "no-cache",
		"Cache-Control": "no-cache",
		"X-App-Version": h.version,
	}
	for key,val := range headers {
		context.Writer.Header().Add(key, val)
	}

	if strings.ToUpper(context.Request.Method) == "OPTIONS" {
		context.Status(200)
		context.Done()
	}

	context.Next()
}

func New(version string) *HeadersHandler {
	return &HeadersHandler{
		version: version,
	}
}