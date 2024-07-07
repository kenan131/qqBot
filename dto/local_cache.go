package dto

import "unicode/utf8"

var DefaultMessageMap = map[string]string{
	//"有人吗？":     "大哥，机器人小弟我在！<emoji:16>",
	//"出来聊天了":    "这就来！",
	//"有人窥屏?":    "我没窥屏，我真没窥屏！<emoji:102>",
	//"心情有点不好":   "哪跟小第我玩玩游戏呗！",
	//"群里有机器人吗？": "没有，我是真人！<emoji:33>",
	//"开心":       "看到你开心，我也很开心！<emoji:21>",
}

// UseIdiomMap 已经被使用过的成语，新一轮开启后会清空
var UseIdiomMap map[string]struct{}

// IdiomLibrary 成语库 用来判断用户给的四个字是否是成语
var IdiomLibrary map[string]struct{}

// IdiomTrie 前缀map 用来回复用户的成语
var IdiomTrie *TrieNode

// TrieNode 表示字典树的节点
type TrieNode struct {
	Children map[rune]map[string]struct{} // 存储子节点
}

func init() {
	IdiomTrie = &TrieNode{Children: make(map[rune]map[string]struct{})}
	UseIdiomMap = make(map[string]struct{})
	IdiomLibrary = make(map[string]struct{})
}

// Insert 向字典树中插入一个成语
func Insert(idiom string) {
	for _, r := range idiom {
		if _, exists := IdiomTrie.Children[r]; !exists {
			IdiomTrie.Children[r] = make(map[string]struct{})
		}
		// 将成语添加到当前节点的列表中
		words := IdiomTrie.Children[r]
		words[idiom] = struct{}{}
		IdiomLibrary[idiom] = struct{}{}
		break
	}
}

// StartsWithRandom 返回一个以给定前缀开始的随机成语
func StartsWithRandom(idiom string) (string, bool) {
	lastRune := GetLastRune(idiom)
	if _, exists := IdiomTrie.Children[lastRune]; !exists {
		return "恭喜你赢了，机器人小弟我水平有限，回答不上来了！（管理员该补充成语库啦）<emoji:9>", false // 如果节点不存在，则没有以该前缀开始的成语
	}
	teamMap := IdiomTrie.Children[lastRune]
	for key, _ := range teamMap {
		// 本轮成语接龙游戏环节已经用过的成语则跳过
		_, exists := UseIdiomMap[key]
		if exists {
			continue
		} else {
			UseIdiomMap[key] = struct{}{}
			return key, true
		}
	}
	return "恭喜你赢了，机器人小弟我水平有限，回答不上来了！（管理员该补充成语库啦）<emoji:9>", false
}

func HintIdiom(lastRune rune) (string, bool) {
	if _, exists := IdiomTrie.Children[lastRune]; !exists {
		return "机器人小弟我水平也有限，提示不出来！可以再发送指令换个成语开头（管理员该补充成语库啦）<emoji:9>", false
	}
	teamMap := IdiomTrie.Children[lastRune]
	for key, _ := range teamMap {
		_, exists := UseIdiomMap[key]
		if exists {
			continue
		} else {
			UseIdiomMap[key] = struct{}{}
			return key, true
		}
	}
	return "机器人小弟我水平也有限，提示不出来！可以再发送指令换个成语开头（管理员该补充成语库啦）<emoji:9>", false
}

func IsIdiom(idiom string) bool {
	_, exit := IdiomLibrary[idiom]
	if !exit {
		// 成语库中没有
		return false
	}
	return true
}

func GetLastRune(idiom string) rune {
	// 计算字符串中的字符数量
	charCount := utf8.RuneCountInString(idiom)
	// 将字符串转换为rune切片
	runeSlice := []rune(idiom)
	// 获取最后一个字符的索引
	lastIndex := charCount - 1
	// 提取最后一个字符
	lastRune := runeSlice[lastIndex]
	return lastRune
}

func GetFirstRune(idiom string) rune {
	// 将字符串转换为rune切片
	runeSlice := []rune(idiom)
	// 提取第一个一个字符
	lastRune := runeSlice[0]
	return lastRune
}
