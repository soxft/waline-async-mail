package mail

import (
	"encoding/json"
	"fmt"
	"github.com/soxft/waline-async-mail/app"
	"github.com/soxft/waline-async-mail/config"
	"github.com/soxft/waline-async-mail/library/mq"
	"github.com/soxft/waline-async-mail/process/redisutil"
	"log"
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

func Handler(data app.CommentStruct) {
	var _mq mq.MessageQueue
	var mail Mail
	_mq = mq.New(redisutil.R, 3)

	_reply := data.Data.Reply
	_comment := data.Data.Comment
	_permalink := config.BlogInfo.Addr + _comment.Url
	_siteTitle := config.BlogInfo.Title

	// 评论邮件 > 发送给 Owner
	ownerArgs := OwnerArgs{
		Author:    _comment.Nick,
		Permalink: _permalink,
		SiteTitle: _siteTitle,
		Text:      _comment.Comment,
		Ip:        _comment.Ip,
		Time:      _comment.InsertedAt,
		Status:    _comment.Status,
		Mail:      _comment.Mail,
	}
	content := parse(config.OwnerTemplate, ownerArgs)
	mail = Mail{
		Subject:   fmt.Sprintf("%s 上有新评论了", _siteTitle),
		Content:   content,
		ToAddress: config.BlogInfo.AuthorEmail,
		Typ:       "toOwner",
	}
	if config.Redis.Enable {
		mailMsg, err := json.Marshal(mail)
		if err != nil {
			log.Printf("[ERROR] json.Marshal error: %s", err)
		}
		_ = _mq.Publish("mail", string(mailMsg))
	} else {
		err := Send(mail, PlatformSmtp)
		if err != nil {
			log.Printf("[ERROR] Send To Owener: %s", err)
		}
	}

	// 回复邮件 > 发送给 Owner & 被回复者
	if _reply.Status != "" {
		guestArgs := GuestArgs{
			Author:    _comment.Nick,
			AuthorP:   _reply.Nick,
			Permalink: _permalink,
			Text:      _comment.Comment,
			TextP:     _reply.Comment,
			SiteTitle: _siteTitle,
		}
		content = parse(config.GuestTemplate, guestArgs)
		mail = Mail{
			Subject:   fmt.Sprintf("%s 回复了你的评论 - %s", _reply.Nick, _siteTitle),
			Content:   content,
			ToAddress: _reply.Mail,
			Typ:       "toGuest",
		}
		if config.Redis.Enable {
			mailMsg, err := json.Marshal(mail)
			if err != nil {
				log.Printf("[ERROR] json.Marshal error: %s", err)
			}
			_ = _mq.Publish("mail", string(mailMsg))
		} else {
			err := Send(mail, PlatformSmtp)
			if err != nil {
				log.Printf("[ERROR] Send To Guest: %s: %s", mail.ToAddress, err)
			}
		}
	}
}
