package test

import (
	"fmt"
	"helloGo/dto"
	"helloGo/service"
	"testing"
	"unicode/utf8"
)

func TestCommon(t *testing.T) {
	t.Run("test getTokenByConfigYaml", func(t *testing.T) {
		token, err := dto.GetToken("config.yaml")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(token)
	})
	t.Run("test getDbConnectUrlYaml", func(t *testing.T) {
		url, err := GetDbConnectUrl("config.yaml")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(url)
	})
	t.Run("test getDb", func(t *testing.T) {
		url, err := GetDbConnectUrl("config.yaml")
		if err != nil {
			t.Error(err)
		}
		_, err1 := service.GetDb(url)
		if err1 != nil {
			t.Error(err1)
		}
	})
	t.Run("test IdiomTrie", func(t *testing.T) {
		idioms := []string{"亡羊补牢", "望梅止渴", "望穿秋水", "画龙点睛"}
		for _, idiom := range idioms {
			dto.Insert(idiom)
		}
		// 随机返回一个以"望"开头的成语
		result, ok := dto.StartsWithRandom("望")
		if ok {
			fmt.Printf("Random idiom starting with '望': %s\n", result)
		} else {
			fmt.Println("No idioms found starting with '望'")
		}
		result1, ok := dto.StartsWithRandom("望")
		if ok {
			fmt.Printf("Random idiom starting with '望': %s\n", result1)
		} else {
			fmt.Println("No idioms found starting with '望'")
		}
		result2, ok := dto.StartsWithRandom("望")
		if ok {
			fmt.Printf("Random idiom starting with '望': %s\n", result2)
		} else {
			fmt.Println("No idioms found starting with '望'")
		}
	})
	t.Run("test common", func(t *testing.T) {
		idiom := "百花香无敌好拉绍" // 一个四字成语
		_, size := utf8.DecodeLastRuneInString(idiom)

		// 从字符串末尾获取最后一个字符
		lastChar, _ := utf8.DecodeRuneInString(idiom[len(idiom)-size:])

		fmt.Printf("The last character of the idiom is: %c\n", lastChar)
	})
	t.Run("test 11", func(t *testing.T) {
		//s := "百花齐放"
		//lastRune := dto.GetLastRune(s)
		//// 打印结果
		//fmt.Printf("The last character '%c' in rune type is: %U\n", lastRune, lastRune)
		s := "成语提示！"
		if "成语提示！" == s {
			fmt.Println(1)
		}
	})
}
