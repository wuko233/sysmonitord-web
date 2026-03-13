package database

import "log"

func InitSchema() error {

	createTableSQL := `
		CREATE TABLE IF NOT EXISTS alerts (
			id SERIAL PRIMARY KEY,
			agent_id VARCHAR(64) NOT NULL,
			alert_type VARCHAR(50) NOT NULL,
			level VARCHAR(20),
			message TEXT,
			source_ip VARCHAR(64),
			username VARCHAR(64),
			file_path TEXT,
			raw_data JSONB,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		-- 创建索引
		CREATE INDEX IF NOT EXISTS idx_alerts_created_at ON alerts (created_at);
		CREATE INDEX IF NOT EXISTS idx_alerts_agent_id ON alerts (agent_id);
	`

	_, err := DB.Exec(createTableSQL)
	if err != nil {
		return err
	}

	log.Printf("数据库表结构初始化完成")
	return nil
}
