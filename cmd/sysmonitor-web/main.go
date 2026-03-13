package main

import (
	"fmt"
	"log"
	"sysmonitor-web/config"
	"sysmonitor-web/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadConfig(""); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	gin.SetMode(config.GlobalConfig.Server.Mode)
	r := gin.Default()
	// 注册路由
	api := r.Group("/api/v1")
	{
		api.GET("/ws", handler.HandleAgentConnect)
		api.GET("/agents", handler.GetAgents)
		api.GET("/alerts", handler.GetAlerts)
	}

	port := config.GlobalConfig.Server.ServerPort
	log.Printf("启动服务器，监听端口: %d", port)

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}

}
