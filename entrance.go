// 所有外部程序使用nblog都使用下面这些方法

package nblog

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// 自定义logger对象
func NewLogger(serverName string) (*zap.SugaredLogger, error) {
	for _, logConfig := range LogConfigs {
		if logConfig.ServerName == serverName {
			return InitLogger(logConfig), nil
		}
	}
	return nil, errors.New("serverName和你配置的不一样，请用config.go里面配置的serverName")
}

// 默认初始化的logger对象
func DefaultLogger() *zap.SugaredLogger {
	logConfig := LogBasicsConfig{
		"default",
		"info",
		"default.log",
		"json",
		2,
		0,
	}
	return InitLogger(logConfig)
}