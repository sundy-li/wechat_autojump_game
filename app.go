package wechat_autojump_game

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"time"

	"image/png"
)

var (
	filename = "wechat_jump.png"

	SleepMills time.Duration = 1000
)

func Run() {
	// filename = "tmp_jump3.png"
	// currentX, currentY, nextX, nextY := getLocation(filename)
	// println(currentX, currentY, nextX, nextY)
	// return
	for true {
		time.Sleep(SleepMills * time.Millisecond)
		saveScreenShot(filename)
		currentX, currentY, nextX, nextY := getLocation(filename)
		if currentX == 0 || currentY == 0 {
			continue
		}
		distance := getDistance(currentX, currentY, nextX, nextY)
		jump(distance)
	}
}

func getLocation(name string) (x, y, nextX, nextY int) {
	f1, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	data, _ := ioutil.ReadAll(f1)
	img, _ := png.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	location := NewLocation(img)
	x, y, nextX, nextY = location.judgeByRGBRange()
	return
}
