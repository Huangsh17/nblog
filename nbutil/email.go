package nbutil

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(msg string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "1132771621@qq.com") //发件人
	m.SetHeader("To", ".cn")                 //收件人
	m.SetHeader("Subject", "Hello!")         //邮件标题
	m.SetBody("text/html", msg)              //邮件内容
	d := gomail.NewDialer("smtp.qq.com", 587, "1132771621@qq.com", "")
	//邮件发送服务器信息,使用授权码而非密码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
