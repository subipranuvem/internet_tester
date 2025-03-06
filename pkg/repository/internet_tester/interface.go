package internettester

import "github.com/subipranuvem/internet_tester/pkg/model"

type InternetTester interface {
	InsertRequestLog(requestLog *model.RequestLog) error
}
