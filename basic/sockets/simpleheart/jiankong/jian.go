package jiankong

import (
	"fmt"

	dingtalk_robot "github.com/JetBlink/dingtalk-notify-go-sdk"
)

// https://oapi.dingtalk.com/robot/send?access_token=e31d0acffdccb16f1c0b467cecf34f91297293504197354bdddf3a0d8af21517
// https://oapi.dingtalk.com/robot/send?access_token=736fa44c80f009f57e6fd1868f5c932f607527c4bf35ef0edebb75dd820d8bee

// Send is to send some messages
func Send(LocalIP ,messages string) {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": "iP:" + LocalIP + messages,
		},
		"at": map[string]interface{}{
			"atMobiles": []string{},
			"isAtAll":   false,
		},
	}
	robot := dingtalk_robot.NewRobot("736fa44c80f009f57e6fd1868f5c932f607527c4bf35ef0edebb75dd820d8bee", "SEC706e06066577b22bb8301a6e8f7cfdb4f365e3b81ea85956d5f4e4ea50507375")

	fmt.Println("hello world")
	if err := robot.SendMessage(msg); err != nil {
		fmt.Println(err)
	}
}
