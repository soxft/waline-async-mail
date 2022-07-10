package mq

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
)

func New(redis *redis.Pool, maxRetries int) MessageQueue {
	return &QueueArgs{
		redis:      redis,
		maxRetries: maxRetries,
	}
}

func (q *QueueArgs) Publish(topic string, msg string, delay int64) error {
	_redis := q.redis.Get()
	defer func(_redis redis.Conn) {
		_ = _redis.Close()
	}(_redis)

	data := &MsgArgs{
		Msg:   msg,
		Retry: 0,
	}
	if _data, err := json.Marshal(data); err != nil {
		return err
	} else {
		_, err := _redis.Do("LPUSH", "rmq:"+topic, string(_data))
		return err
	}
}

func (q *QueueArgs) Subscribe(topic string, processes int, handler func(data string)) {
	for i := 0; i < processes; i++ {
		go func() {
			_redis := q.redis.Get()
			defer func() {
				// handle error
				if err := recover(); err != nil {
					q.Subscribe(topic, 1, handler)
				}
			}()

			// 阻塞
			wg := sync.WaitGroup{}
			for {
				_data, err := redis.Strings(_redis.Do("BRPOP", "rmq:"+topic, 1))
				if err != nil || _data == nil {
					continue
				}

				var _dataString = _data[1]
				var _msg MsgArgs
				if err := json.Unmarshal([]byte(_dataString), &_msg); err != nil {
					continue
				}

				wg.Add(1)
				// execute handler
				go func(_msg MsgArgs) {
					defer func() {
						// retry if error
						wg.Done()
						if err := recover(); err != nil {
							log.Printf("[ERROR] mq handler: %s", err)
							if _data, err := json.Marshal(_msg); err != nil {
								return
							} else {
								// max retry
								if _msg.Retry > q.maxRetries {
									_, _ = _redis.Do("LPUSH", "rmq:"+topic+"failed", _data)
									return
								}
								_, _ = _redis.Do("LPUSH", "rmq:"+topic, _data)
							}
							_msg.Retry++
						}
					}()
					handler(_msg.Msg)
				}(_msg)
				// wait for handler
				wg.Wait()
			}
		}()
	}
}
