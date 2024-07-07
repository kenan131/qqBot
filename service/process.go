package service

import (
	"database/sql"
	"fmt"
	"helloGo/dto"
	"helloGo/net"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Processor struct {
	Api net.MyHttpClient
	Db  *sql.DB
}

// 猜数游戏 答案
var targetNum int

// 猜数游戏 状态
var guessNumFlag = dto.GameEnd

// 成语接龙 上一个成语的末尾字符
var preRune rune

// 成语接龙 状态
var idiomFlag = dto.GameEnd

func (p *Processor) ProcessAtMessage(data *dto.Message) error {
	//指令解析
	match := dto.CommandRegular.FindStringSubmatch(data.Content)
	var replyContent = ""
	if len(match) > 2 {
		// 指令模式
		switch match[1] {
		case "001":
			// 猜数游戏
			replyContent = Instructions001(match[2])
		case "002":
			// 成语接龙
			replyContent = p.Instructions002(match[2])
		case "003":
			replyContent = p.Instructions003(match[2])
		case "004":
			replyContent = p.Instructions004(match[2])
		case "005":
			replyContent = cancelGameStatus()
		}
	} else {
		content := ETLInput(data.Content)
		// 普通消息模式
		if content == "玩法介绍" {
			replyContent =
				`指令一，猜数游戏：输入指令/001 可指定猜数范围，如指定100则会随机生成一个范围为0-100的数。如不指定则生成0-10000之间的数哦!
				指令二，成语接龙：输入指令/002 
				指令三，给机器人添加默认回复。格式/003 key : reply，例如怎么获取资料 : 第一步...... 
				指令四，删除指令三添加的默认回复，格式/004 key`
		} else if content == "你好！" {
			replyContent = "你好！"
		} else if isNumeric(content) && guessNumFlag != dto.GameEnd {
			// 猜数游戏
			replyContent = guessNumberGame(content)
		} else if IsLengthFour(content) && idiomFlag != dto.GameEnd {
			replyContent, _ = handlerIdiom(content, false)
		} else if content == "成语提示！" && idiomFlag != dto.GameEnd {
			replyContent, _ = dto.HintIdiom(preRune)
		} else {
			replyContent = "暂时看不懂你的指令，可以发送`玩法介绍`查看本机器人小弟的使用说明哦！<emoji:16>"
		}
	}
	replyMessage := &dto.MessageToCreate{
		Content: replyContent,
		MessageReference: &dto.MessageReference{
			// 引用这条消息
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}
	p.Api.PostMethodParam(dto.MessagesURI, "channel_id", data.ChannelID, replyMessage)
	return nil
}

func (p *Processor) ProcessMessage(data *dto.Message) error {
	if replyContent, exists := dto.DefaultMessageMap[data.Content]; exists {
		replyMessage := &dto.MessageToCreate{
			Content: replyContent,
			MessageReference: &dto.MessageReference{
				// 引用这条消息
				MessageID:             data.ID,
				IgnoreGetMessageError: true,
			},
		}
		p.Api.PostMethodParam(dto.MessagesURI, "channel_id", data.ChannelID, replyMessage)
	}
	// 不存在默认回复的key 则不发送消息
	return nil
}

func Instructions001(content string) string {
	rand.Seed(time.Now().UnixNano())
	if num, err := strconv.Atoi(content); err != nil {
		targetNum = rand.Intn(10000)
	} else if num > 0 {
		targetNum = rand.Intn(num)
	} else {
		targetNum = rand.Intn(10000)
	}
	guessNumFlag = dto.GameStart
	return "猜数游戏开始啦，@我参与答题，猜小猜大都会有提示哦！"
}
func (p *Processor) Instructions002(content string) string {
	if IsLengthFour(content) {
		// 清空上一轮被使用的成语
		dto.UseIdiomMap = make(map[string]struct{})
		reply, flag := handlerIdiom(content, true)
		if flag {
			// 如果第一个成语回复成功了，则开启游戏状态
			idiomFlag = dto.GameStart
		}
		return reply
	}
	return "指令后面好像不是四字成语，得你先起个头哦！"
}

func (p *Processor) Instructions003(content string) string {
	parameters := strings.SplitN(content, ":", 2)
	if len(parameters) != 2 {
		return "添加失败！格式好像有些问题。(key:value)中间是英文冒号哦！ "
	} else {
		key := parameters[0]
		reply := parameters[1]
		err := InsertFixReply(key, reply, p.Db)
		if err != nil {
			return "操作失败，请联系管理员查看，或者等等再试试！"
		}
		dto.DefaultMessageMap[key] = reply
		return "添加成功，赶快来试试吧！"
	}
}

func (p *Processor) Instructions004(key string) string {
	// 删除key
	res, err := DeleteFixReply(key, p.Db)
	if err == nil && res == "删除成功！" {
		// 删除map中的元素
		delete(dto.DefaultMessageMap, key)
	}
	return res
}

func guessNumberGame(content string) string {
	num, _ := strconv.Atoi(content)
	if num == targetNum {
		// 猜对则切换状态
		targetNum = -1
		guessNumFlag = dto.GameEnd
		return "恭喜你猜对了!<emoji:79>"
	} else if num < targetNum {
		return "猜的数有点小哦！"
	} else {
		return "猜的数有点大哦！"
	}
}

// isNumeric 检查字符串是否可以转换为整数
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func ETLInput(input string) string {
	etlData := string(dto.AtRE.ReplaceAll([]byte(input), []byte("")))
	etlData = strings.Trim(etlData, dto.SpaceCharSet)
	return etlData
}
func handlerIdiom(idiom string, first bool) (string, bool) {
	lastRule := dto.GetFirstRune(idiom)
	if lastRule != preRune && !first {
		content := fmt.Sprintf("你输入的成语好像没接上上一个成语哦！该以:`%s`开头哦！", string(preRune))
		return content, false
	}
	res := dto.IsIdiom(idiom)
	if res == false {
		return "不好意思，你输入的这四个字好像不是个成语吧。或者暂时还没有被收入进成语库中！", false
	}
	replyIdiom, flag := dto.StartsWithRandom(idiom)
	if flag == false {
		// 回答不上来！ 结束游戏状态
		idiomFlag = dto.GameEnd
		return replyIdiom, flag
	}
	// 保存回复的最后一个汉字
	saveIdiomLastRune(replyIdiom)
	return replyIdiom, flag
}

func saveIdiomLastRune(idiom string) {
	preRune = dto.GetLastRune(idiom)
}
func IsLengthFour(content string) bool {
	if utf8.RuneCountInString(content) != 4 {
		return false // 确保字符串长度为4
	}
	return true
}
func cancelGameStatus() string {
	guessNumFlag = dto.GameEnd
	idiomFlag = dto.GameEnd
	return "游戏状态清空啦！"
}
