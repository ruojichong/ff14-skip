package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	time1 := time.Now().Unix()
	cmd := exec.Command("cmd", "/C", "date 2099/09/24")
	err1 := cmd.Run()
	if err1 != nil {
		fmt.Println("使用管理员权限运行,按任意键退出")
		time.Sleep(10 * time.Second)
	}
	time.Sleep(1 * time.Second)
	cmd2 := exec.Command("cmd", "/C", "date "+time.Unix(time1, 0).Format("2006/01/02"))
	cmd2.Run()
}
