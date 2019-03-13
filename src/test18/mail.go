package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

var (
	from = "xiayanji@qiniu.com"
	to   = []string{"576407585@qq.com"}
	host = "smtp.exmail.qq.com:25"
	hostWithTls = "smtp.exmail.qq.com:465"
	hostNoPort = "smtp.exmail.qq.com"
	subj = "golang邮件主题"
	body = "golang邮件正文，\r\n http://www.baidu.com"
	passwd = "LAsH6R5wpEX5fjfF"
	toMails = "576407585@qq.com,xiayanji@qiniu.com"
)

func test2() {
	fromMail := mail.Address{"", from}
	toMail := mail.Address{"", toMails}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = fromMail.String()
	headers["To"] = toMail.String()
	headers["Subject"] = subj
	

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := hostWithTls

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", from, passwd, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(fromMail.Address); err != nil {
		log.Panic(err)
	}
	
	toMailList, err := mail.ParseAddressList(toMails)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range toMailList {
		if err = c.Rcpt(v.Address); err != nil {
			log.Panic(err, v)
		}
	}
	

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}

func test1() {
	auth := smtp.PlainAuth("", from, passwd, hostNoPort)

	msg := []byte("From: 来自Golang\r\n" +
		"To: 576407585@qq.com\r\n" +
		"Subject: Golang mail!\r\n" +
		"\r\n" +
		"golang测试body.\r\n")
	err := smtp.SendMail(host, auth, from, to, msg)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//test1()
	test2()
}
