package mq

import (
	"github.com/gomodule/redigo/redis"
)

type MessageQueue interface {
	Publish(topic string, msg string, delay int64) error
	Subscribe(topic string, processes int, handler func(msg string))
}

type QueueArgs struct {
	redis      *redis.Pool
	maxRetries int
}

type MsgArgs struct {
	Msg   string `json:"msg"`
	Retry int    `json:"retry"`
}
