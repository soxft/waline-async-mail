package mqutil

import (
	"github.com/soxft/waline-async-mail/config"
	"github.com/soxft/waline-async-mail/library/mq"
	"github.com/soxft/waline-async-mail/process/redisutil"
)

var Q mq.MessageQueue

// Init
// @desc golang消息队列
func Init() {
	if !config.Redis.Enable {
		return
	}
	Q = mq.New(redisutil.R, 3)

	Q.Subscribe("mail", config.Redis.Concurrency, Mail)
}
