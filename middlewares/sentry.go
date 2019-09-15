package middlewares

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

var SentryClient *raven.Client

func SentryReportor(dsn string) gin.HandlerFunc {
	SentryClient, _ := raven.New(dsn)

	return sentry.Recovery(SentryClient, true)
}
