package handler

import (
	"net/http"
	"sysmonitor-web/internal/model"

	"github.com/gin-gonic/gin"
)

// TODO: 数据库接入
var OnlineAgents = make(map[string]bool)

func GetAgents(c *gin.Context) {

	agents := []string{"演示1"}
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    200,
		Message: "success",
		Data:    agents,
	})
}

func GetAlerts(c *gin.Context) {

	alerts := []map[string]interface{}{ // 演示数据
		{"time": "10:00", "type": "SSH", "level": "HIGH", "msg": "Root login detected"},
		{"time": "10:05", "type": "FILE", "level": "MEDIUM", "msg": "/etc/passwd modified"},
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    200,
		Message: "success",
		Data:    alerts,
	})
}
