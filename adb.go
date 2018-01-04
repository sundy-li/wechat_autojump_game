package wechat_autojump_game

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strings"
)

var (
	SCREEN_WIDTH  = 1080
	SCREEN_HEIGHT = 1920
	//这个系数好难调,好像不同机型不一样...
	Speed float64 = 1.392
)

func jump(distance float64) {
	pressTime := distance * Speed
	x, y := randomPosition()
	runAdb("shell", fmt.Sprintf("input swipe %d %d %d %d %d", x, y, x, y, int(pressTime)))
}

func saveScreenShot(filename string) {
	filePath := "/sdcard/" + filename
	runAdb("shell", "screencap -p "+filePath)
	runAdb("pull", filePath, ".")
}

func runAdb(args ...string) {
	var b bytes.Buffer
	cmd := exec.Command("adb", args...)
	cmd.Stdout = &b
	cmd.Stderr = &b
	log.Printf("adb %s", strings.Join(args, " "))
	err := cmd.Run()
	if cmd.Process != nil {
		cmd.Process.Kill()
	}
	if err != nil {
		log.Fatalf("adb %s: %v", strings.Join(args, " "), err.Error())
	}
}

//x : 350 - 450
//y : 850 - 950
func randomPosition() (x int, y int) {
	x = 350 + rand.Intn(100)
	y = 850 + rand.Intn(100)
	return
}
