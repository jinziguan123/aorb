package main

import (
	"context"
	"github.com/BigNoseCattyHome/aorb/backend/api-gateway/middleware"
	"github.com/BigNoseCattyHome/aorb/backend/utils/constants/config"
	"github.com/BigNoseCattyHome/aorb/backend/utils/extra/tracing"
	"github.com/BigNoseCattyHome/aorb/backend/utils/logging"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	auth2 "github.com/BigNoseCattyHome/aorb/backend/go-services/auth/web"
	comment2 "github.com/BigNoseCattyHome/aorb/backend/go-services/comment/web"
	user2 "github.com/BigNoseCattyHome/aorb/backend/go-services/user/web"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	tp, err := tracing.SetTraceProvider(config.WebServerName)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Panicf("Error to set the trace")
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logging.Logger.WithFields(logrus.Fields{
				"error": err,
			}).Errorf("Error to set the trace")
		}
	}()

	g := gin.Default()
	// 配置prometheus
	p := ginprometheus.NewPrometheus("AorB-WebGateway")
	p.Use(g)
	// 配置gzip
	g.Use(gzip.Gzip(gzip.DefaultCompression))
	// 配置tracing
	g.Use(otelgin.Middleware(config.WebServerName))
	g.Use(middleware.Authenticate())

	rootPath := g.Group("/aorb")
	{
		rootPath.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		user := rootPath.Group("/user")
		{
			user.GET("/", user2.UserHandler)
			user.POST("/login/", auth2.LoginHandler)
			user.POST("/register/", auth2.RegisterHandle)
		}

		comment := rootPath.Group("/comment")
		{
			comment.GET("/action/", comment2.ActionCommentHandler)
			comment.GET("/list/", comment2.ListCommentHandler)
			comment.POST("/count/", comment2.CountCommentHandler)
		}
	}

	// run
	if err := g.Run(config.WebServerAddr); err != nil {
		panic("Can not run aorb Gateway, binding port: " + config.WebServerAddr)
	}

}
