package middleware

import (
	"auth-go/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

var loggerKey string = "logger_key"

type LoggerRecord struct {
	URL         string       `json:"url"`
	Time        string       `json:"time"`
	Body        string       `json:"body"`
	Header      http.Header  `json:"header"`
	Method      string       `json:"method"`
}

func logFormatter(params gin.LogFormatterParams) string {
	content, _ := ioutil.ReadAll(params.Request.Body)

	message, _ := json.Marshal(LoggerRecord {
		params.Path,
		params.TimeStamp.Format(time.RFC3339),
		string(content),
		params.Request.Header,
		params.Method,
	})

	return string(message)
}

type LoggerOutput struct {
}

func (lo LoggerOutput) Write(p []byte) (int, error) {
	content := string(p)

	if _, error := cache.GetRedisCacheInstance().UnShiftList(loggerKey, content); error != nil {
		return 0, error
	}
	return len(p), nil
}

func HandleLoggerMiddleware() gin.HandlerFunc {
	config := gin.LoggerConfig {
		Formatter: logFormatter,
		Output: LoggerOutput{},
		SkipPaths: []string{},
	}
	return gin.LoggerWithConfig(config)
}
