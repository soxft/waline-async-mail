package handler

import (
	"encoding/json"
	"fmt"
	"github.com/soxft/waline-async-mail/app"
	"github.com/soxft/waline-async-mail/config"
	"github.com/soxft/waline-async-mail/library/mail"
	"github.com/soxft/waline-async-mail/process/mqutil"
	"github.com/soxft/waline-async-mail/templates"
	"log"
	"os"
	"reflect"
	"strings"
)

// parse parse mail template
func parse(template string, args interface{}) string {
	r := reflect.TypeOf(args)
	v := reflect.ValueOf(args)
	for i := 0; i < r.NumField(); i++ {
		tag := fmt.Sprintf("{{%s}}", r.Field(i).Tag.Get("args"))
		data := v.Field(i).String()
		template = strings.ReplaceAll(template, tag, data)
	}

	return template
}

func Send(data app.CommentStruct) {
	var _mail mail.Mail

	_reply := data.Data.Reply
	_comment := data.Data.Comment
	_permalink := config.BlogInfo.Addr + _comment.Url
	_siteTitle := config.BlogInfo.Title

	// 评论邮件 > 发送给 Owner
	ownerArgs := mail.OwnerArgs{
		Author:    _comment.Nick,
		Permalink: _permalink,
		SiteTitle: _siteTitle,
		Text:      _comment.Comment,
		Ip:        _comment.Ip,
		Time:      _comment.InsertedAt,
		Status:    _comment.Status,
		Mail:      _comment.Mail,
	}
	content := parse(templates.Owner, ownerArgs)
	_mail = mail.Mail{
		Subject:   fmt.Sprintf("%s 上有新评论了", _siteTitle),
		Content:   content,
		ToAddress: config.BlogInfo.AuthorEmail,
		Typ:       "toOwner",
	}
	if _comment.Mail != config.BlogInfo.AuthorEmail {
		handlerSendMail(_mail)
	}

	// 回复邮件 > 发送给 Owner & 被回复者
	if _reply.Status != "" {
		guestArgs := mail.GuestArgs{
			Author:    _comment.Nick,
			AuthorP:   _reply.Nick,
			Permalink: _permalink,
			Text:      _comment.Comment,
			TextP:     _reply.Comment,
			SiteTitle: _siteTitle,
		}
		content = parse(templates.Guest, guestArgs)
		_mail = mail.Mail{
			Subject:   fmt.Sprintf("%s 回复了你的评论 - %s", _comment.Nick, _siteTitle),
			Content:   content,
			ToAddress: _reply.Mail,
			Typ:       "toGuest",
		}
		if _reply.Mail != "" && _reply.Mail != config.BlogInfo.AuthorEmail {
			handlerSendMail(_mail)
		}
	}
}

func handlerSendMail(_mail mail.Mail) {
	log.SetOutput(os.Stdout)
	if config.Redis.Enable {
		mailMsg, err := json.Marshal(_mail)
		if err != nil {
			log.Printf("[ERROR] json.Marshal error: %s", err)
		}
		_ = mqutil.Q.Publish("mail", string(mailMsg))
	} else {
		err := mail.Send(_mail, mail.PlatformSmtp)
		log.Printf("[INFO] Mail send [%s]: %s", _mail.Typ, _mail.ToAddress)
		if err != nil {
			log.Printf("[ERROR] Mail send err [%s] %s: %s", _mail.Typ, _mail.ToAddress, err)
		}
	}
}
