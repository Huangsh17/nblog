## nblog
nblog是非常方便，好用的日志库，封装zap的很多特性，而且比zap更加容易上手，zap是对于日志比较底层的封装，有很多的特性，是我们需要深入研究源代码，才可以做一些定制化功能。
## 快速入门
安装： go get github.com/Huangsh17/nblog  
示例代码  
logger := nblog.DefaultLogger()  
logger.Info("info)  
logger.Debug("info)  
logger.Error("info)

## 使用细节

