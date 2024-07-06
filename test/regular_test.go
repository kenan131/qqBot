package test

import (
	"fmt"
	"helloGo/dto"
	"testing"
)

func TestRegular(t *testing.T) {
	t.Run(
		"", func(t *testing.T) {
			var input = "/001你好啊！"
			matches := dto.CommandRegular.FindStringSubmatch(input)
			if len(matches) > 2 {
				fmt.Println("指令编号:", matches[1])
				fmt.Println("指令后面的内容:", matches[2])
			} else {
				fmt.Println("没有匹配到内容")
			}
		})
}
