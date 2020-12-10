package nblog

// basic setting
type LogBasicsConfig struct {
	// 服务名称
	ServerName string
	// 等级
	Level string
	// 输出文件路径 默认在当前项目根目录生成logs目录下存放日志
	OutPath string
	// 编码格式，支持console，json
	Encoding string
	// 输出方式 0:输出控制台 1:输出文件 2 ：同时输出文件和控制台;文件输出路径不用填写
	OutputType int
	// 0:不是邮箱通知 1：使用邮箱通知
	IsEmail int
}

type Email struct {
	From string
	To string
	Msg string
	// 授权邮箱服务器 ，比如说qq的邮箱地址为：smtp.qq.com
	Host string
	// 登陆方邮箱地址
	UserName string
	 // 邮箱服务授权码，不是邮箱密码
	PassWord string
}

