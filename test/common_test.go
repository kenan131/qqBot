package test

import (
	"fmt"
	"helloGo/dto"
	"helloGo/service"
	"testing"
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
}
