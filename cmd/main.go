package main

import (
	"flag"
	"time"

	g "github.com/sundy-li/wechat_autojump_game"
)

var (
	mills int64
)

func init() {
	flag.Int64Var(&mills, "m", 0, "millseconds sleep after each jump")
	flag.Parse()
}
func main() {
	if mills != 0 {
		g.SleepMills = time.Duration(mills)
	}
	g.Run()
}
