package mail

import (
	"fmt"
	"github.com/Bottlist/bottlist/pkg/config"
	"net/smtp"
	"strings"
)

type Client struct {
	HostName string
	Port     string
	UserName string
	Auth     smtp.Auth
}

func NewMailClient() *Client {
	conf := config.LoadConfig()
	auth := smtp.PlainAuth("", conf.MailInfo.UserName, conf.MailInfo.PassWord, conf.MailInfo.HostName)
	return &Client{
		HostName: conf.MailInfo.HostName,
		Port:     conf.MailInfo.Port,
		UserName: conf.MailInfo.UserName,
		Auth:     auth,
	}
}

func (c *Client) SendMail(
	// 宛先アドレス
	to []string,
	// 件名
	subject string,
	// 本文
	body string,
) error {
	msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(to, ","), subject, body), "\n", "\r\n"))
	return smtp.SendMail(fmt.Sprintf("%s:%s", c.HostName, c.Port), c.Auth, c.UserName, to, msg)
}
