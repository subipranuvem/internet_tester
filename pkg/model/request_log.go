package model

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RequestLog struct {
	gorm.Model
	RequestAddress    string `gorm:"column:request_address" json:"request_address"`
	StatusCode        int    `gorm:"column:status_code" json:"status_code"`
	Method            string `gorm:"column:method" json:"method"`
	ResponseTimeMs    int    `gorm:"column:response_time_ms" json:"response_time_ms"`
	TimeoutTimeMs     int    `gorm:"column:timeout_time_ms" json:"timeout_time_ms"`
	SpeedTestResultID int    `gorm:"column:speed_test_result_id" json:"speed_test_result_id"`
	SpeedTestResult   SpeedTestResult
}

func (r *RequestLog) TableName() string {
	return "requests"
}

func (r *RequestLog) LogInfo() {
	logger := zap.S()
	logger.Infow("request log", "content", r)
}
