package nbutil

import (
	"gopkg.in/gomail.v2"
)
func SendEmail(msg string)  {
	m := gomail.NewMessage()
	m.SetHeader("From", "1132771621@qq.com")                     //发件人
	m.SetHeader("To", "huangsonghui@dm-ai.cn")           //收件人
	m.SetHeader("Subject", "Hello!")                     //邮件标题
	m.SetBody("text/html", msg)     //邮件内容
	m.Attach("C:\\Users\\DM\\Desktop\\微信图片_20201023140433.png")       //邮件附件

	d := gomail.NewDialer("smtp.qq.com", 587, "1132771621@qq.com", "foychtxjtxijffcd")
	//邮件发送服务器信息,使用授权码而非密码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}