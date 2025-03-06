package internettester

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunSpeedTest(t *testing.T) {
	result, err := RunSpeedTest()

	assert.NoError(t, err, "RunSpeedTest should not return an error")
	assert.NotNil(t, result, "RunSpeedTest should return a valid result")

	if result != nil {
		assert.Greater(t, result.DownloadSpeedMbps, 0.0, "Download speed should be greater than 0")
		assert.Greater(t, result.UploadSpeedMbps, 0.0, "Upload speed should be greater than 0")
		assert.GreaterOrEqual(t, result.Ping, 0.0, "Ping should be 0 or greater")
		assert.NotEmpty(t, result.ISP, "ISP should not be empty")
		assert.NotEmpty(t, result.IPAddress, "IP Address should not be empty")
	}
}
