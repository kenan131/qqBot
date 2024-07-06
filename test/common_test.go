package test

import (
	"fmt"
	"helloGo/dto"
	"helloGo/log"
	"helloGo/service"
	"testing"
)

func TestCommon(t *testing.T) {
	t.Run("test getTokenByConfigYaml", func(t *testing.T) {
		token, err := dto.GetToken("config.yaml")
		if err != nil {
			log.Error(err)
		}
		fmt.Println(token)
	})
	t.Run("test getDbConnectUrlYaml", func(t *testing.T) {
		url, err := service.GetDbConnectUrl("config.yaml")
		if err != nil {
			log.Error(err)
		}
		fmt.Println(url)
	})
	t.Run("test getDb", func(t *testing.T) {
		_, err := service.GetDb("config.yaml")
		if err != nil {
			log.Error(err)
		}

	})
}
