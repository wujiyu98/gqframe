package email

import (
	"github.com/wujiyu98/gqframe/config"
	"gopkg.in/gomail.v2"
)

var (
	host     string
	port     int
	username string
	password string
	toMail   string
)

func init() {
	host = config.V.GetString("email.host")
	port = config.V.GetInt("email.port")
	username = config.V.GetString("email.username")
	password = config.V.GetString("email.password")
	toMail = config.V.GetString("email.toMail")

}

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", toMail)
	m.SetHeader("Reply-To", "2030716640@qq.com")
	m.SetHeader("Subject", "你好没有一个中文的邮件看下3")
	m.SetBody("text/html", "你好，这是一个测试邮箱，添加多点文字")

	d := gomail.NewDialer(host, port, username, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
