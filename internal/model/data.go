package model

import (
	"time"
)

type AgentPacket struct {
	Type      string      `json:"type"`
	Timestamp int64       `json:"timestamp"`
	Payload   interface{} `json:"payload"`
}

type SSHAlertPayload struct {
	Type      string    `json:"type"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Data      struct {
		Hostname string `json:"hostname"`
		Username string `json:"username"`
		Method   string `json:"method"`
		SourceIP string `json:"source_ip"`
		Port     int    `json:"port"`
		Service  string `json:"service"`
		PID      string `json:"pid"`
	} `json:"data"`
}

type FileAlertEvent struct {
	FilePath  string    `json:"file_path"`
	Operation string    `json:"operation"`
	Time      time.Time `json:"time"`
}

// InfoMonitorConfig 信息监控配置
type InfoMonitorConfig struct {
	Enabled        bool          `yaml:"enabled"`
	Interval       time.Duration `yaml:"interval"`        // 采集间隔
	LogFilePath    string        `yaml:"log_file_path"`   // 日志文件路径
	MaxLogSize     int64         `yaml:"max_log_size"`    // 最大日志大小（字节）
	LogRetention   int           `yaml:"log_retention"`   // 日志保留天数
	ProcessLimit   int           `yaml:"process_limit"`   // 显示进程数限制
	CollectNetwork bool          `yaml:"collect_network"` // 是否收集网络信息
	CollectProcess bool          `yaml:"collect_process"` // 是否收集进程信息
}

// ServerMetrics 服务器指标
type ServerMetrics struct {
	Timestamp    time.Time     `json:"timestamp"`
	CPU          CPUInfo       `json:"cpu"`
	Memory       MemoryInfo    `json:"memory"`
	Disk         []DiskInfo    `json:"disk"`
	Network      NetworkInfo   `json:"network"`
	Load         LoadInfo      `json:"load"`
	Processes    []ProcessInfo `json:"processes"`
	Host         HostInfo      `json:"host"`
	Runtime      RuntimeInfo   `json:"runtime"`
	QuickMetrics QuickMetrics  `json:"quick_metrics"`
}

// CPUInfo CPU信息
type CPUInfo struct {
	Model          string    `json:"model"`
	Cores          int       `json:"cores"`
	LogicalCores   int       `json:"logical_cores"`
	UsagePercent   float64   `json:"usage_percent"`
	PerCorePercent []float64 `json:"per_core_percent"`
	Mhz            float64   `json:"mhz"`
	CacheSize      int       `json:"cache_size"`
}

// MemoryInfo 内存信息
type MemoryInfo struct {
	TotalGB     float64 `json:"total_gb"`
	UsedGB      float64 `json:"used_gb"`
	AvailableGB float64 `json:"available_gb"`
	UsedPercent float64 `json:"used_percent"`
	SwapTotalGB float64 `json:"swap_total_gb"`
	SwapUsedGB  float64 `json:"swap_used_gb"`
}

// DiskInfo 磁盘信息
type DiskInfo struct {
	Mountpoint    string  `json:"mountpoint"`
	Device        string  `json:"device"`
	Fstype        string  `json:"fstype"`
	TotalGB       float64 `json:"total_gb"`
	UsedGB        float64 `json:"used_gb"`
	FreeGB        float64 `json:"free_gb"`
	UsedPercent   float64 `json:"used_percent"`
	InodesPercent float64 `json:"inodes_percent"`
}

// NetworkInfo 网络信息
type NetworkInfo struct {
	Interfaces      []NetworkInterface `json:"interfaces"`
	TotalRecvMB     float64            `json:"total_recv_mb"`
	TotalSentMB     float64            `json:"total_sent_mb"`
	TCPConnections  int                `json:"tcp_connections"`
	EstablishedConn int                `json:"established_conn"`
}

// NetworkInterface 网络接口
type NetworkInterface struct {
	Name         string   `json:"name"`
	HardwareAddr string   `json:"hardware_addr"`
	IPAddresses  []string `json:"ip_addresses"`
}

// LoadInfo 负载信息
type LoadInfo struct {
	Load1          float64 `json:"load_1"`
	Load5          float64 `json:"load_5"`
	Load15         float64 `json:"load_15"`
	RelativeLoad1  float64 `json:"relative_load_1"`
	RelativeLoad5  float64 `json:"relative_load_5"`
	RelativeLoad15 float64 `json:"relative_load_15"`
	ProcsRunning   int     `json:"procs_running"` // 改为 int 类型
	ProcsTotal     int     `json:"procs_total"`   // 改为 int 类型
}

// ProcessInfo 进程信息
type ProcessInfo struct {
	PID        int32   `json:"pid"`
	Name       string  `json:"name"`
	Cmdline    string  `json:"cmdline"`
	MemoryMB   float64 `json:"memory_mb"`
	CPUPercent float64 `json:"cpu_percent"`
}

// HostInfo 主机信息
type HostInfo struct {
	Hostname        string    `json:"hostname"`
	OS              string    `json:"os"`
	Platform        string    `json:"platform"`
	PlatformVersion string    `json:"platform_version"`
	KernelVersion   string    `json:"kernel_version"`
	BootTime        time.Time `json:"boot_time"`
	Uptime          string    `json:"uptime"`
	CPUCount        uint64    `json:"cpu_count"`
	Architecture    string    `json:"architecture"`
	HostID          string    `json:"host_id"`
}

// RuntimeInfo 运行时信息
type RuntimeInfo struct {
	GoVersion    string `json:"go_version"`
	GOOS         string `json:"goos"`
	GOARCH       string `json:"goarch"`
	GOROOT       string `json:"goroot"`
	GOMAXPROCS   int    `json:"gomaxprocs"`
	NumCPU       int    `json:"num_cpu"`
	NumGoroutine int    `json:"num_goroutine"`
}

// QuickMetrics 快速指标
type QuickMetrics struct {
	CPUPercent        float64 `json:"cpu_percent"`
	MemoryPercent     float64 `json:"memory_percent"`
	RootDiskPercent   float64 `json:"root_disk_percent"`
	AvailableMemoryGB float64 `json:"available_memory_gb"`
}

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
