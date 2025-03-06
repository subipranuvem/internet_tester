package internettester

import (
	"github.com/subipranuvem/internet_tester/pkg/model"
	repo "github.com/subipranuvem/internet_tester/pkg/repository/internet_tester"
	"go.uber.org/zap"
)

type internetTester struct {
	config InternetTesterConfig
	logger *zap.SugaredLogger
}

type InternetTesterConfig struct {
	Repository repo.InternetTester
	AppConfig  model.AppConfig
}

func NewInternetTester(config InternetTesterConfig) *internetTester {
	return &internetTester{
		config: config,
		logger: zap.S(),
	}
}

// TestInternet runs the internet tests and saves the results to the database
func (i *internetTester) TestInternet() error {
	result, err := i.runInternetTests()
	if err != nil {
		i.logger.Error("running internet test:", err)
		return err
	}

	err = i.config.Repository.InsertRequestLog(result)
	if err != nil {
		i.logger.Error("inserting request log:", err)
		return err
	}

	return nil
}

// runInternetTests runs the internet tests and returns the results
func (i *internetTester) runInternetTests() (*model.RequestLog, error) {
	request, err := MakeRequest(i.config.AppConfig.TargetURL, i.config.AppConfig.TimeoutInSecs)
	if err != nil {
		i.logger.Errorf("request failed:", err)
		return nil, err
	}

	speedTestResult, speedTestErr := RunSpeedTest()
	if speedTestErr != nil {
		i.logger.Errorf("speed test failed:", speedTestErr)
		return nil, err
	}

	request.SpeedTestResult = *speedTestResult
	i.logger.Info("request completed")
	request.LogInfo()

	return request, nil
}
