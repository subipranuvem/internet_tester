package main

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/subipranuvem/internet_tester/pkg/model"
	repo "github.com/subipranuvem/internet_tester/pkg/repository/internet_tester"
	postgresrepo "github.com/subipranuvem/internet_tester/pkg/repository/internet_tester/postgres"
	internettester "github.com/subipranuvem/internet_tester/pkg/tools/internet_tester"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	err := godotenv.Load()
	if err != nil {
		logger.Sugar().Info("no .env file detected")
	}

	var appConfig model.AppConfig
	err = env.Parse(&appConfig)
	if err != nil {
		logger.Sugar().Fatal("loading env variables:", err)
	}

	db, err := gorm.Open(postgres.Open(appConfig.DBConnStr), &gorm.Config{})
	if err != nil {
		logger.Sugar().Fatal("database connection error:", err)
	} else {
		logger.Sugar().Info("database connection successful")
	}

	var internetTesterRepo repo.InternetTester
	internetTesterRepo, err = postgresrepo.NewPostgresInternetTester(db)
	if err != nil {
		logger.Sugar().Fatal("repostiory creation error:", err)
	}

	registerConfig := internettester.InternetTesterConfig{
		Repository: internetTesterRepo,
		AppConfig:  appConfig,
	}

	tester := internettester.NewInternetTester(registerConfig)

	for {
		err := tester.TestInternet()
		if err != nil {
			logger.Sugar().Error("error at testing internet: %v\n", err)
		}
		time.Sleep(appConfig.IntervalBetweenRequestsInSecs * time.Second)
	}
}
