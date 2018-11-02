// mail包提供常用的发邮件功能, 目前只支持SMTP邮件服务器.
// 基于github.com/qor/mailer, 这里只是做了初始化等细节的工作.
package mailbz

import (
	"net/mail"

	"github.com/bigzhu/gobz/confbz"
	"github.com/go-gomail/gomail"
	"github.com/qor/assetfs"
	qorMailer "github.com/qor/mailer"
	"github.com/qor/mailer/gomailer"
)

type (
	// Attachment 附件
	Attachment = qorMailer.Attachment
	// Address email地址
	Address = mail.Address
	// Template 邮件模板
	Template = qorMailer.Template
	// Email email结构
	Email = qorMailer.Email
)

const assetFSNameSpace = "mailer"

// Send 功能最全的邮件发送方式, 但是易错, 比如: 发送HTML时Email.Text字段必须为空,使用模板时Email.Text和Email.HTML都必须为空.
func Send(email Email, templates ...Template) (err error) {
	conf := confbz.GetEmailConf()
	mailer, err := newMailer(conf.Host, conf.Port, conf.User, conf.Password, conf.AssetPaths)
	if err != nil {
		return
	}
	if email.From == nil {
		email.From = &Address{Address: conf.User}
	}
	err = mailer.Send(email, templates...)
	return
}

// SendText 发送纯文本
func SendText(replyTo *Address, to []Address, subject, content string, attachments ...Attachment) (err error) {
	return Send(Email{
		ReplyTo:     replyTo,
		TO:          to,
		Subject:     subject,
		Text:        content,
		Attachments: attachments,
	})
}

// SendHTML 发送HTML内容
func SendHTML(replyTo *Address, to []Address, subject, content string, attachments ...Attachment) (err error) {
	return Send(Email{
		ReplyTo:     replyTo,
		TO:          to,
		Subject:     subject,
		HTML:        content,
		Attachments: attachments,
	})
}

// SendTemplate 使用邮件模板发送邮件. 目前只支持模板文件.
// 模板文件:
// 		"layout模板文件", 用于装饰"模板文件", 可选.
//      "内容模板文件", 生成的内容被插入到"layout模板文件"生成的内容中.
// 		layout文件要放在"${assetPath}/layouts/"目录中, 内容模板文件放在"${assetPath}/"目录.
// 内容格式:
//		与html/template包要求的格式一致.
//		"layout模板文件"的内容需要包含`{{ yield }}`, 才能把"内容模板文件"包含进去.
func SendTemplate(replyTo *Address, to []Address, subject string, templates []Template, attachments ...Attachment) (err error) {
	return Send(Email{
		ReplyTo:     replyTo,
		TO:          to,
		Subject:     subject,
		Attachments: attachments,
	}, templates...)
}

func newMailerByConf() (mailer *qorMailer.Mailer, err error) {
	conf := confbz.GetEmailConf()
	mailer, err = newMailer(conf.Host, conf.Port, conf.User, conf.Password, conf.AssetPaths)
	return
}

func newMailer(host string, port int, username, password string, assetBaseDirs []string) (mailer *qorMailer.Mailer, err error) {
	//TODO 支持imap等服务器
	dailer := gomail.NewDialer(host, port, username, password)
	sender, err := dailer.Dial()
	if err != nil {
		return
	}
	assetFS := assetfs.AssetFS().NameSpace(assetFSNameSpace)
	for _, dir := range assetBaseDirs {
		assetFS.RegisterPath(dir)
	}
	mailer = qorMailer.New(&qorMailer.Config{
		Sender:  gomailer.New(&gomailer.Config{Sender: sender}),
		AssetFS: assetFS,
	})
	return
}
