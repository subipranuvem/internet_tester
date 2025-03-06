package internettester

import (
	"time"

	"github.com/showwin/speedtest-go/speedtest"
	"github.com/showwin/speedtest-go/speedtest/transport"
	"github.com/subipranuvem/internet_tester/pkg/model"
	"go.uber.org/zap"
)

func RunSpeedTest() (*model.SpeedTestResult, error) {
	logger := zap.S()
	var speedtestClient = speedtest.New()

	user, err := speedtest.FetchUserInfo()
	if err != nil {
		logger.Error("fetch user info:", err)
		return nil, err
	}

	serverList, err := speedtestClient.FetchServers()
	if err != nil {
		logger.Error("fetch speed test servers:", err)
		return nil, err
	}

	targets, err := serverList.FindServer([]int{})
	if err != nil {
		logger.Error("find speed test server:", err)
		return nil, err
	}

	server := targets[0]
	logger.Info("selected server:", server)

	now := time.Now()
	server.PingTest(nil)
	logger.Infow("ping test finished", "duration", time.Since(now).Round(time.Millisecond))

	now = time.Now()
	server.DownloadTest()
	logger.Infow("download test finished", "duration", time.Since(now).Round(time.Millisecond))

	now = time.Now()
	server.UploadTest()
	logger.Infow("upload test finished", "duration", time.Since(now).Round(time.Millisecond))
	server.Context.Reset()

	packetLossDup := 0
	packetLossMax := 0
	packetLossSent := 0
	packetLossPercentage := 0.0

	now = time.Now()
	analyzer := speedtest.NewPacketLossAnalyzer(nil)
	err = analyzer.Run(server.Host, func(packetLoss *transport.PLoss) {
		packetLossDup = packetLoss.Dup
		packetLossMax = packetLoss.Max
		packetLossSent = packetLoss.Sent
		packetLossPercentage = packetLoss.LossPercent()
	})
	logger.Infow("packet loss test finished", "duration", time.Since(now).Round(time.Millisecond))

	if err != nil {
		logger.Error("run packet loss analyzer:", err)
		return nil, err
	}

	result := &model.SpeedTestResult{
		Ping:                 float64(server.Latency.Milliseconds()),
		DownloadSpeedMbps:    server.DLSpeed.Mbps(),
		UploadSpeedMbps:      server.ULSpeed.Mbps(),
		ISP:                  user.Isp,
		IPAddress:            user.IP,
		PacketLossDup:        packetLossDup,
		PacketLossMax:        packetLossMax,
		PacketLossSent:       packetLossSent,
		PacketLossPercentage: packetLossPercentage,
	}

	return result, nil
}
