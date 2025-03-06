package internettester

import (
	"net/http"
	"time"

	"github.com/subipranuvem/internet_tester/pkg/model"
	"go.uber.org/zap"
)

func MakeRequest(targetURL string, timeoutInSecs time.Duration) (*model.RequestLog, error) {
	logger := zap.S()

	client := &http.Client{
		Timeout: timeoutInSecs,
	}

	startTime := time.Now()

	method := http.MethodGet
	statusCode := 0
	responseTime := 0

	resp, err := client.Get(targetURL)
	if err != nil {
		return nil, err
	}

	responseTime = int(time.Since(startTime).Milliseconds())
	logger.Infow("request finished", "duration", time.Since(startTime).Round(time.Millisecond))

	defer resp.Body.Close()
	statusCode = resp.StatusCode

	requestLog := &model.RequestLog{
		RequestAddress: targetURL,
		StatusCode:     statusCode,
		Method:         method,
		ResponseTimeMs: responseTime,
		TimeoutTimeMs:  int(timeoutInSecs.Milliseconds()),
	}

	return requestLog, nil
}
