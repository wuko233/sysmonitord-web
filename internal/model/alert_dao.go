package model

import (
	"encoding/json"
	"sysmonitor-web/internal/database"
	"time"
)

type AlertRecord struct {
	AgentID  string
	Type     string
	Level    string
	Message  string
	SourceIP string
	Username string
	FilePath string
	RawData  interface{}
}

func (r *AlertRecord) Save() error {
	rawBytes, err := json.Marshal(r.RawData)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO alerts (agent_id, alert_type, level, message, source_ip, username, file_path, raw_data, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = database.DB.Exec(query,
		r.AgentID, r.Type, r.Level, r.Message,
		r.SourceIP, r.Username, r.FilePath, rawBytes, time.Now(),
	)
	return err
}
