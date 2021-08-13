package weapp

import (
	"testing"
)

func TestMain(t *testing.T) {

	cli := NewClient("wx83683c749c99edfa", "3940aa2d43058f71bd5de479d36ab442")

	msg := cli.NewSubscribeMessage()
	// GetTemplateList 获取帐号下已存在的模板列表
	//
	_, err := msg.GetTemplateList()
	if err != nil {
		t.Error(err)
		// 处理一般错误信息
		return
	}

	// fmt.Printf("返回结果: %+v", res)

	t.Error("\n==========")
}
