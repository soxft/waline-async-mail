package main

import (
	"github.com/soxft/waline-async-mail/process/mqutil"
	"github.com/soxft/waline-async-mail/process/redisutil"
	"github.com/soxft/waline-async-mail/process/webutil"
)

func main() {
	redisutil.Init()
	mqutil.Init()
	webutil.Init()
}
