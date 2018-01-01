package wechat_autojump_game

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var (
	SCREEN_WIDTH  = 1080
	SCREEN_HEIGHT = 1920
	//这个系数好难调,好像不同机型不一样...
	Speed float64 = 1.392
)

const (
	ADB_TAP_COMMAND = "input swipe 510 953 510 953"
)

func jump(distance float64) {
	pressTime := distance * Speed
	runAdb("shell", ADB_TAP_COMMAND+" "+strconv.Itoa(int(pressTime)))
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
