package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	fmt.Println("请以管理员权限运行")
	fmt.Println("有动画时按任意键并回车,即可通过掉线跳过当次动画")
	for {
		_, _ = fmt.Scanln()
		time1 := time.Now().Unix()
		cmd := exec.Command("cmd", "/C", "date 2099/09/24")
		err1 := cmd.Run()
		if err1 != nil {
			fmt.Println("未管理员权限运行,程序将在3秒后退出")
			time.Sleep(3 * time.Second)
			os.Exit(1)
		}
		time.Sleep(1 * time.Second)
		cmd2 := exec.Command("cmd", "/C", "date "+time.Unix(time1, 0).Format("2006/01/02"))
		_ = cmd2.Run()
	}

}
