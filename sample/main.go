package main

import (
	"fmt"
	"github.com/huhongda/GoToolBox/smscaptcha"
)

func main() {
	cap := smscaptcha.New("redis://localhost:6379", "LTAIKgSfnWzmXz1c", "DcCbEjzNjv9GYABtPbhs31CictXZXD", "SMS_44290055", "58分享", true)
	fmt.Println(cap.SmsSend("15828566956"))
}
