package main

import (
	"flag"
	"fmt"
	"time"

	g "github.com/sundy-li/wechat_autojump_game"
)

var (
	mills int64
	speed float64
)

func init() {
	flag.Int64Var(&mills, "m", 0, "millseconds sleep after each jump")
	flag.Float64Var(&speed, "s", 0, fmt.Sprintf("speed value, default to , %.4f", g.Speed))
	flag.Parse()
}
func main() {
	if mills != 0 {
		g.SleepMills = time.Duration(mills)
	}
	if speed != 0 {
		g.Speed = speed
	}
	g.Run()
}
