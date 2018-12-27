package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wncbb/be_blog/internal/pkg/env"
	"github.com/wncbb/be_blog/internal/app/be_blog/config"
	"github.com/sirupsen/logrus"
	"os"
)

func Init() {
	env.Init()
	logrus.SetOutput(os.Stdout)
	if env.IsProduct() {
		logrus.SetLevel(logrus.WarnLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	config.Init()
}

func main() {
	Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(config.GetRunAddr()) // listen and serve on 0.0.0.0:8080
}
