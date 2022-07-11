package main

import (
	"github.com/soxft/waline-async-mail/process/mqutil"
	"github.com/soxft/waline-async-mail/process/webutil"
)

func main() {
	mqutil.Init()
	webutil.Init()
}
