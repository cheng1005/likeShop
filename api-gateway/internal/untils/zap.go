package untils

import (
	"go.uber.org/zap"
	"net/http"
)

var SugarLogger *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()
	SugarLogger = logger.Sugar()
}

func SimpleHttpGet(url string) {
	SugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		SugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		SugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
