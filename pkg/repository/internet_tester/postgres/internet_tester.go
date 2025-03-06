package postgres

import (
	"github.com/subipranuvem/internet_tester/pkg/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type postgresInternetTester struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewPostgresInternetTester(db *gorm.DB) (*postgresInternetTester, error) {
	logger := zap.S()

	err := db.AutoMigrate(&model.RequestLog{}, &model.SpeedTestResult{})
	if err != nil {
		logger.Error("table migration error:", err)
		return nil, err
	}

	internetTester := &postgresInternetTester{
		db:     db,
		logger: logger,
	}

	return internetTester, nil
}

func (pit *postgresInternetTester) InsertRequestLog(requestLog *model.RequestLog) error {
	result := pit.db.Create(requestLog)
	if result.Error != nil {
		pit.logger.Error("request result save error:", result.Error)
		return result.Error
	}

	return nil
}
