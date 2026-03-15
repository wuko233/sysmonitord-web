package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sysmonitor-web/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源，仅限开发使用
	},
}

func HandleAgentConnect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}

	defer conn.Close()

	agentID := c.Query("agent_id")
	if agentID == "" {
		agentID = "unknown"
	}

	log.Printf("Agent %s connected", agentID)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}

		if messageType == websocket.TextMessage {
			var agentPackage model.AgentPacket
			if err := json.Unmarshal(message, &agentPackage); err != nil {
				log.Printf("解析消息失败: %v", err)
				continue
			}

			// TODO: 处理接收到的数据包
			processPacket(agentID, agentPackage)
		}
	}
}

func processPacket(agentID string, agentPackage model.AgentPacket) {
	log.Printf("接受到来自Agent %s 的数据包: 类型=%s, 时间戳=%d", agentID, agentPackage.Type, agentPackage.Timestmap)

	switch agentPackage.Type {

	case "SSH_ROOT_LOGIN", "SSH_ALERT":
		bytes, _ := json.Marshal(agentPackage.Payload)
		var event model.SSHLoginEvent

		if err := json.Unmarshal(bytes, &event); err != nil {
			log.Printf("解析SSH登录事件失败: %v", err)
			return
		}

		handleSSHAlert(agentID, event)

	case "REALTIME_FILE_ALERT", "NON_WHITELISTED_FILE":
		bytes, _ := json.Marshal(agentPackage.Payload)
		var event model.FileAlertEvent

		if err := json.Unmarshal(bytes, &event); err != nil {
			log.Printf("解析文件警报事件失败: %v", err)
			return
		}

		handleFileAlert(agentID, event)

	case "STATUS_UPDATE":

	}
}

func handleSSHAlert(agentID string, event model.SSHLoginEvent) {
	log.Printf("处理SSH登录事件: Agent=%s, 用户=%s, 来源IP=%s",
		agentID, event.Username, event.SourceIP)

	record := model.AlertRecord{
		AgentID:  agentID,
		Type:     "SSH_LOGIN",
		Level:    "HIGH",
		Message:  event.Message,
		SourceIP: event.SourceIP,
		Username: event.Username,
		RawData:  event,
	}

	if err := record.Save(); err != nil {
		log.Printf("保存SSH登录警报记录失败: %v", err)
	}
}

func handleFileAlert(agentID string, event model.FileAlertEvent) {
	log.Printf("处理文件警报事件: Agent=%s, 文件=%s, 操作=%s",
		agentID, event.FilePath, event.Operation)

	record := model.AlertRecord{
		AgentID:  agentID,
		Type:     "FILE_ALERT",
		Level:    "MEDIUM",
		Message: event.Operation,
		FilePath: event.FilePath,
		RawData:  event,
	}

	if err := record.Save(); err != nil {
		log.Printf("保存文件警报记录失败: %v", err)
	}
}
