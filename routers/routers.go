package routers

import (
	jwt "userInfoService/middlewares"
	"userInfoService/pkg/logging"
	"userInfoService/pkg/setting"
	"userInfoService/routers/api"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
)

func InitRouter() *gin.Engine {

	logrus.SetLevel(logrus.TraceLevel)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(cors.Default())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/test", func(c *gin.Context) {
			// logrus.Trace("trace msg")
			// logrus.Debug("debug msg")
			// logrus.Info("info msg")
			// logrus.Warn("warn msg")
			// logrus.Error("error msg")
			// logrus.Fatal("fatal msg")
			// logrus.Panic("panic msg")
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}

	r.GET("/ping", func(c *gin.Context) {
		url := "https://www.mayanan.cn"
		logging.SugarLogger.Debugf("ping debug for %s", url)
		logging.SugarLogger.Errorf("ping error for %s", url)
		// logging.SugarLogger.Info("Success..",
		// 	zap.String("statusCode", "200"),
		// 	zap.String("url", "https://www.mayanan.cn"))

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/v1/userinfo")

	return r
}
