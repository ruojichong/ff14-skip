package main

import (
	"os/exec"
	"time"
)

func main() {
	time1 := time.Now().Unix()
	cmd := exec.Command("runas", "/trustlevel:0x40000", "cmd", "/C", "date 2099/09/22")
	cmd.Run()
	time.Sleep(1 * time.Second)
	cmd2 := exec.Command("cmd", "/C", "date "+time.Unix(time1, 0).Format("2006/01/02"))
	cmd2.Run()
}