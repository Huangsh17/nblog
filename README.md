## nblog
nblog是非常方便，好用的日志库，封装zap的很多特性，而且比zap更加容易上手，zap是对于日志比较底层的封装，有很多的特性，是我们需要深入研究源代码，才可以做一些定制化功能，目前这个仓库，还在完善和开发中，后面的话也会把精力放到这个仓库的开发中，继续完善这个日志库。
## 快速入门
安装： go get github.com/Huangsh17/nblog  
示例代码  
```
logger := nblog.DefaultLogger()
logger.Info("info")  
logger.Debug("debug")  
logger.Error("error")
```

## 特性
支持对接消息中间件  
支持邮件通知  
支持对某个完整的业务链进行日志输出（比如说购物流程，我们就可以通过一个uuid，把一条完整购物流程日志串联起来，更容易排查问题）  
更多功能还在进一步开发中  
