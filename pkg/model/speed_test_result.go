package model

import "gorm.io/gorm"

type SpeedTestResult struct {
	gorm.Model
	Ping                 float64 `gorm:"column:ping_ms" json:"ping_ms"`
	DownloadSpeedMbps    float64 `gorm:"column:download_speed_mbps" json:"download_speed_mbps"`
	UploadSpeedMbps      float64 `gorm:"column:upload_speed_mbps" json:"upload_speed_mbps"`
	ISP                  string  `gorm:"column:isp" json:"isp"`
	IPAddress            string  `gorm:"column:ip_address" json:"ip_address"`
	PacketLossDup        int     `gorm:"column:packet_loss_dup" json:"packet_loss_dup"`
	PacketLossMax        int     `gorm:"column:packet_loss_max" json:"packet_loss_max"`
	PacketLossSent       int     `gorm:"column:packet_loss_sent" json:"packet_loss_sent"`
	PacketLossPercentage float64 `gorm:"column:packet_loss_percentage" json:"packet_loss_percentage"`
}

func (s *SpeedTestResult) TableName() string {
	return "speed_test_results"
}
