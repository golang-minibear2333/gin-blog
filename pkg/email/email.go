package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	// 发件人
	m.SetHeader("From", e.From)
	// 收件人
	m.SetHeader("To", to...)
	// 邮件主题
	m.SetHeader("Subject", subject)
	// 正文
	m.SetBody("text/html", body)
	// 创建一个 SMTP 拨号实例
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	// 设置对应的拨号信息
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	// 打开与 SMTP 服务器的连接并发送电子邮件
	return dialer.DialAndSend(m)
}
