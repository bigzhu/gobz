package mailbz

import (
	"fmt"
	"net"
	"testing"

	"github.com/bigzhu/gobz/confbz"
)

var (
	from        = Address{Address: "oeohomos@163.com", Name: "jyw"}
	to          = []Address{{Address: "vermiliondun@gmail.com", Name: "Mr jyw"}}
	replyTo     = Address{Address: "oeohomos@163.com", Name: "jyw"}
	attachments = []Attachment{
		{FileName: "/home/jre/Documents/challenge.png"},
		{FileName: "/home/jre/Documents/challenge.png", Inline: true},
	}
)

func TestSendText(t *testing.T) {
	err := SendText(&from, to, "测试text邮件", "text email")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendHTML(t *testing.T) {
	err := SendHTML(&from, to, "测试html邮件", "html email <img src='http://img1.imgtn.bdimg.com/it/u=3873040750,1105020127&fm=26&gp=0.jpg'/>")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendWithAttachments(t *testing.T) {
	err := SendHTML(
		&from,
		to,
		"测试html邮件",
		"html email <img src='http://img1.imgtn.bdimg.com/it/u=3873040750,1105020127&fm=26&gp=0.jpg'/>",
		attachments...,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendWithTemplate(t *testing.T) {
	data := struct {
		Title   string
		Content string
	}{
		"测试模板邮件", "内容由模板生成",
	}
	err := SendTemplate(
		&from,
		to,
		"测试模板邮件(内容不应该是空的)",
		[]Template{{Name: "hello", Layout: "application", Data: data}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDialSMTP(t *testing.T) {
	conf := confbz.GetEmailConf()
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		fmt.Println(err)
	}
}
