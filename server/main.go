package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	middleware "go-gin-boilerplate/server/middleware"
	"go-gin-boilerplate/server/utils/confighelper"
	"go-gin-boilerplate/server/utils/logginghelper"

	"github.com/gin-gonic/gin"
)

func main() {
	setupEnvironment()
	startServer()
}

func startServer() {
	logginghelper.LogInfo("Starting server...")
	router := gin.Default()
	middleware.InitMiddleware(router)
	router.GET("/", checkStatus())

	//Configure server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", confighelper.GetConfig("serverPort")),
		Handler:        router,
		ReadTimeout:    confighelper.GetConfigAsDuration("serverReadTimeout") * time.Second,
		WriteTimeout:   confighelper.GetConfigAsDuration("serverWriteTimeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//Start server
	logginghelper.LogInfo("Server running on Port ", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		logginghelper.LogError("Unable to start server : ", err)
	}

}

func setupEnvironment() {
	env := flag.String("env", "dev", "To set environment debug/production")
	flag.Parse()

	//set Gin mode
	if *env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		confighelper.InitViper("config")
	} else if *env == "stg" {
		confighelper.InitViper("config_stg")
	} else {
		confighelper.InitViper("config_dev")
	}
	//role specific logger
	logginghelper.Init("logs/server.log", confighelper.GetConfigAsBool("writeLogs"), confighelper.GetConfigAsInt("maxBackupCount"), confighelper.GetConfigAsInt("maxBackupFileSize"), confighelper.GetConfigAsInt("maxAgeForBackupFiles"), true)
}

func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully!!!")
	}
}
