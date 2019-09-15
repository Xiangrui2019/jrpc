package middlewares

import (
	"jrpc/modules"

	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

func SentryReportor() gin.HandlerFunc {
	return sentry.Recovery(modules.SentryClient, true)
}
