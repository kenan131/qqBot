package dto

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"runtime"
)

type Token struct {
	AppID       uint64
	AccessToken string
}

func New() *Token {
	return &Token{}
}

func (t *Token) LoadFromConfig(file string) error {
	var conf struct {
		AppID uint64 `yaml:"appid"`
		Token string `yaml:"token"`
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(content, &conf); err != nil {
		return err
	}
	t.AppID = conf.AppID
	t.AccessToken = conf.Token
	return nil
}
func (t *Token) GetString() string {
	return fmt.Sprintf("%v.%s", t.AppID, t.AccessToken)
}

func GetToken(name string) (*Token, error) {
	token := New()
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		if err := token.LoadFromConfig(fmt.Sprintf("%s/%s", path.Dir(filename), name)); err != nil {
			return nil, err
		}
	}
	return token, nil
}
