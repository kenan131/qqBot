### 快速运行

项目数据使用数据库进行存储，运行项目前先将本项目根目录下的sql导入到开发数据库。文件名`qqbot.sql`

修改配置文件：config.yaml中的配置参数

```yaml
#写上自己机器人的appid
appid: ???
#写上自己机器人的token
token: ??? 
# 修改mysql连接url
connectUrl: root:123456@tcp(127.0.0.1:3306)/qqBot?charset=utf8mb4&parseTime=True&loc=Local 
```

找到hello.go文件运行main函数。

### 使用说明

- 指令一，猜数游戏：输入指令/001 {整形参数，可选}，如指定100则会随机生成一个范围为0-100的数。如不指定或者整数解析失败则默认生成0-10000之间的数哦!  

- 指令二，成语接龙：输入指令/002 {四字成语}，机器人会自动根据你给的成语开始接龙。每次根据成语的某尾汉字作为开头成语。成语接龙环节，如接不上上一个成语，艾特机器人并发送{成语提示}消息，机器人会给你发送成语提示哦！

- 指令三，给机器人添加默认回复：格式/003 {key:reply}，样例唱跳rap:篮球，添加成功后，发送唱跳rap则机器人会自动回复篮球消息哦！

- 指令四，删除指令三添加的默认回复，格式/004 {key}  

- 指令五，清空游戏状态，处于猜数游戏和成语接龙游戏的状态会被清空掉哦~

#### 指令一：猜数游戏

艾特机器人并选择指令一，即/001 {整形参数，可选}触发猜数游戏，以下为游戏介绍：

参数为**整形**，则根据输入的数据，去随机生成一个目标值

- /001 500 系统解析数据成功，随机生成一个0-500的目标值。

参数为**其他**

- /001 ，不输入参数，则随机生成一个0-10000的目标值。

- /001 中文，输入中文，解析数据失败，则随机生成一个0-10000的目标值。

猜数环节，艾特机器人并发送整形数据即可参与

- 当发送的整数等于目标值，则成功！

- 当发送的整数小于目标值，机器人则提醒数据过小

- 当发送的整数大于目标值，则机器人提醒数据过大

使用截图：

![](https://blogimg-1311682597.cos.ap-guangzhou.myqcloud.com/%E5%85%AC%E7%94%A8/2G%25RJT56WW_%28%28JA%5D%284%294L_D.png)

#### 指令二：成语接龙

艾特机器人并输入指令/002 {四字成语}，则触发游戏。

成语接龙**游戏规则**：以用户发送的第一个成语作为开头，后续需根据四字成语的最后一个字，作为成下一个成语的开头字进行接龙。发送的参数必须是**四字成语**才可触发机器人的成语接龙回复。

<mark>可能用户发送的四字成语会被误判为不是成语，是因为成语库没有维护这个成语导致的。目前数据库只维护了1000多个成语😓</mark>

<mark>如果遇到可以换个四字成语试试</mark>

> 如果用户接不上上一句成语，则可以艾特机器人并发送{成语提示}。获取下一个成语。

使用截图：

<img src="https://blogimg-1311682597.cos.ap-guangzhou.myqcloud.com/%E5%85%AC%E7%94%A8/9VKJ%25Z_L%29Q%5DWQM_O%403%60_Q%5BW.png?q-sign-algorithm=sha1&q-ak=AKIDlDwyBtc-s9A-aOcFDvjX8JwrW0M3RCWA90MnrH6Csl_h3C-5QDT8dbO1p_UnfE6C&q-sign-time=1720344754;1720348354&q-key-time=1720344754;1720348354&q-header-list=host&q-url-param-list=ci-process&q-signature=96b5d241d0057413ec83822ef73857620ccc810e&x-cos-security-token=EPimACmGSBbVCMyTsfmA87ZCL1mEv67aa75aedb70a692d23db64220aa63d0cc27PNBjuXtUzdsKG-S3WrEy4iP1q9_thMGk_Z9jbqN3kJpWHFyZ12mmW1yQyvwKLS0YJ7sTi-5jNgtg8umU4H_pFnxx102EaI7IWz10qxadttPkPyyFXquYPua5YMNB1jkFUif4DZYSD-sDjDxWLzh-NTNBnOkPQpU-FTq4DmgKfkosN9mMeW_-LkyU--D7Zg4&ci-process=originImage" title="" alt="" data-align="center">

![](https://blogimg-1311682597.cos.ap-guangzhou.myqcloud.com/%E5%85%AC%E7%94%A8/MWI%29339P%29C7%29GX2%7B_%240%7DORC.png?q-sign-algorithm=sha1&q-ak=AKIDFXk-jNAerbkKkutz-xB--99u2numjHqn-Ps4yLgbaN60L5JPKRGSDWNufYhfV2Rt&q-sign-time=1720344953;1720348553&q-key-time=1720344953;1720348553&q-header-list=host&q-url-param-list=ci-process&q-signature=9872275155a461c1bca97a7153df1158b99cf257&x-cos-security-token=j1V187XhL7IH3KfLdhfRjAOWyKyKogmaf35db31cd502f5a61a04a9fa26c01a7f8Kb4YlgSj1djU_5UN8LOURYGXyZBioFsIX6J32SURHQDGg8nLb9ITjL0PoztVmB79V2eoxLjpRfGw1N_ea-ftPpH5Y6A-5iEpPs9ab54KxkrFn9cMYgjrzmWqmQ77HucDNUJo-5Gl7SwNs4UHzCIpoejb5dYoZNAues3A86HWGmemrfy9PIejMplz1OG-vLZ&ci-process=originImage)

![](https://blogimg-1311682597.cos.ap-guangzhou.myqcloud.com/%E5%85%AC%E7%94%A8/%7BKV_0V6%25%29U%5BKF%60%5B%28SULR~%29X.png?q-sign-algorithm=sha1&q-ak=AKIDa1NF71WmSLl5S8hgHgc6mCxuZsOTBiaAtbYPaZ2ywf5U0H974Nq6vQZDCa52yhKT&q-sign-time=1720345038;1720348638&q-key-time=1720345038;1720348638&q-header-list=host&q-url-param-list=ci-process&q-signature=2f461f8d6903ad5c724bf15f7f29e9863ebaeff4&x-cos-security-token=EPimACmGSBbVCMyTsfmA87ZCL1mEv67a6445cbe37be63ea160e4dd9044e0774c7PNBjuXtUzdsKG-S3WrEy47liBWjWLk0z6X22bd6d7s0pkcb_rdMRAUdIETFiRBX--z9WbQndRYicsq_Hi158fM5PQ7ClCXpgygSrQ44dlR9IIT6CK0vEqg5-_fzVS9Bmx5QEMKgLlFFZel_rU8xMvzJktB4ceQnTRsws5ChaEgXVP_w4E2bPyRkXQSFU8_g&ci-process=originImage)

#### 指令三：设置默认回复

功能介绍：可以根据自己需求设置一个固定回复话术，如`唱跳rap`:`篮球`，或者`怎么获取资料`:`联系管理员2号哦！`等。

艾特机器人并输入指令/003 {key:value}，进行设置。设置成功和失败都会有相应提示~~**中间的冒号是英文的**

> 如果已经存在key则会默认覆盖掉旧的value哦！

使用截图：

![](https://blogimg-1311682597.cos.ap-guangzhou.myqcloud.com/%E5%85%AC%E7%94%A8/sjdsjA.png?q-sign-algorithm=sha1&q-ak=AKID7P9oga8zVXoFZWmImh8r_O-bY92XThET78BUxKcPTYdd3GMdf2za04G2KJLH_8BD&q-sign-time=1720346385;1720349985&q-key-time=1720346385;1720349985&q-header-list=host&q-url-param-list=ci-process&q-signature=b3e44bdce4c284648df9c088c314b554e4c89d40&x-cos-security-token=EPimACmGSBbVCMyTsfmA87ZCL1mEv67a41ddb8f0c0495ef88c255f4f69deb7b77PNBjuXtUzdsKG-S3WrEywxrhSzvpIw1SsPKh6FjhYF2UvYnR7Tu7N4rrvn4roRNVc7DEUW2QgazUvH6OnOehebmPPO2tG8b0QIothyoDX1pWrGvo5wdnGIEtSe6wRs3La7lKn5LkughLsAK5RClVMIL-YS0dZx7XsXICtEILhz_jFmWLm0-6eJR4TFNo3k3&ci-process=originImage)

#### 指令四：删除默认回复

功能介绍：用来删除指令三添加的默认回复

艾特机器人并输入指令/003 {key}，进行删除。如果删除的key不存在，会进行提示哦~~

![](https://blogimg-1311682597.cos.ap-guangzhou.myqcloud.com/%E5%85%AC%E7%94%A8/777F.png?q-sign-algorithm=sha1&q-ak=AKIDTL77p5havxKF6YHGbRYJ88oErpMjSiqmtaw5VL4R6nKoXzQ_6_8Rxt1UIHJMz1iC&q-sign-time=1720346885;1720350485&q-key-time=1720346885;1720350485&q-header-list=host&q-url-param-list=ci-process&q-signature=322408361aab58b2d1d77bc3dc7126a56fd6de91&x-cos-security-token=EPimACmGSBbVCMyTsfmA87ZCL1mEv67ac0995d0dd4cbf904abbffac995b02be97PNBjuXtUzdsKG-S3WrEy4wNogKeD2xtc2KqhEi1NHHQM8fpGllvwLsOUL0_KTrj4FxnCdVSdu3greCAPVvvmPi5dW7Nrg9AC9_dihGG4_fj_pUYGP9CqqyG7eA_NrY2CcXwRfyH5-x05_zGXqF9apdX5I1FX2Ahr4hMm999y4n-lqdolYgJdVNddl1c_yz5&ci-process=originImage)

#### 指令五：清空游戏状态

功能介绍：如果猜数游戏 或者 成语接龙游戏不想玩了，可以将游戏状态清空。如果不清空游戏状态可能会影响其他指令使用。

猜数游戏：影响艾特机器人参数带整形的。

成语接龙：影响艾特机器人参数带四个字的。

<mark>现在还没有加太多的指令进去，如果不清空其实也不影响使用！</mark>

使用截图：

![](https://blogimg-1311682597.cos.ap-guangzhou.myqcloud.com/%E5%85%AC%E7%94%A8/cc54a9324.png?q-sign-algorithm=sha1&q-ak=AKID3-8Dy3k3ebExOcEuYfKGNTWj89_K2-PKGEwsRRgvt1eaLm4cIjqVrwvr0k5ed9NV&q-sign-time=1720347212;1720350812&q-key-time=1720347212;1720350812&q-header-list=host&q-url-param-list=ci-process&q-signature=8ed9ad15e0d062c65f72a70fe9bc86d7d1e5dac2&x-cos-security-token=EPimACmGSBbVCMyTsfmA87ZCL1mEv67ac6bbe4926ec5c976b9a7cc2492b3720e7PNBjuXtUzdsKG-S3WrEy-zV3w-oMjMRfLRMU0zTUnnXOSxBXe0OxPli1-S6ExxCYMrEDkmazJlIeCo0sOIwFFuw29wxa-zK1vp8lfVeycwg0eWlPScVx2Uv7FdYnrtfslVW8C_j6oK1NM9pymGcYdlsggAweat9I6NkHaogodvvy_Z1OxhM2C2IcO5H6IkX&ci-process=originImage)

### 设计说明

#### 获取配置文件中的配置参数

qq机器人的配置appid和token配置在根目录下的config.yaml文件中。通过代码进行获取。

```go
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

```

#### 事件处理器

采用策略模式的设计思想，抽象出一个事件处理方法，使用map进行存储。根据不同的监听事件注册不同的事件处理器。

```go

var eventHandlerMap = map[dto.OPCode]map[dto.EventType]EventHandler{}

// EventHandler 抽象方法
type EventHandler func(event *dto.WSPayload, message []byte) error

func RgisterHandler(code dto.OPCode, eventT dto.EventType, handler EventHandler) {
    // 添加到map中
    if _, exists := eventHandlerMap[code]; !exists {
        eventHandlerMap[code] = make(map[dto.EventType]EventHandler)
    }
    eventHandlerMap[code][eventT] = handler
}

func HandlerProcess(code dto.OPCode, eventT dto.EventType, payload *dto.WSPayload) error {
    if tempMap, ok := eventHandlerMap[code]; ok {
        if handler, ok1 := tempMap[eventT]; ok1 {
            // 调用 eventHandler
            handler(payload, payload.RawMessage)
        } else {
            log.Infof("没有添加该事件的处理器,opCode:%d eventType:%s", code, eventT)
        }
    }
    return nil
}
```

`RgisterHandler`方法在main程序入口处进行调用注册监听事件处理器。

 `HandlerProcess`方法在websocket有监听事件时进行触发调用。

#### 成语接龙

成语库数据来自数据库，在项目启动后会将数据库中的数据都缓存到map中。

主要数据结构`map[rune]map[string]struct{} `使用汉字作为key，存储所有以该汉字开头的成语。

```go
// UseIdiomMap 已经被使用过的成语，不能再使用了，新一轮开启后会清空
var UseIdiomMap map[string]struct{}

// IdiomLibrary 成语库 用来判断用户给的四个字是否是成语
var IdiomLibrary map[string]struct{}

// IdiomTrie 前缀map 用来回复用户的成语
var IdiomTrie *TrieNode

// TrieNode 节点
type TrieNode struct {
    // 根据汉字 存储所有该汉字开头的成语
    Children map[rune]map[string]struct{} 
}
// 插入
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
        // 如果节点不存在，则没有以该前缀开始的成语
        return "恭喜你赢了，机器人小弟我水平有限，回答不上来了！（管理员该补充成语库啦）<emoji:9>", false
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
```

#### 指令三：默认回复

项目启动后，会默认将数据库fix_reply表中的数据初始化到map中。

处理普通消息时使用消息内容作为key去查询map。如果有则发送固定回复，没有则不处理。

```go
var DefaultMessageMap = map[string]string{
    //"有人吗？":     "大哥，机器人小弟我在！<emoji:16>",
    //"出来聊天了":    "这就来！",
    //"有人窥屏?":    "我没窥屏，我真没窥屏！<emoji:102>",
    //"心情有点不好":   "哪跟小第我玩玩游戏呗！",
    //"群里有机器人吗？": "没有，我是真人！<emoji:33>",
    //"开心":       "看到你开心，我也很开心！<emoji:21>",
}
// 普通消息处理方法
func (p *Processor) ProcessMessage(data *dto.Message) error {
    // 如果缓存中有该key，则说明有该指令的默认回复。
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
```

#### 数据库操作

编写数据库表的dao操作类。供处理类调用。

```go
func GetFixReplyByKey(key string, db *sql.DB) (*entity.FixReply, error) {
    rows, err := db.Query("SELECT `id`, `key`, `reply` FROM fix_reply where `key` = ? ", key)
    if err != nil {
        return nil, err
    }
    var reply = &entity.FixReply{}
    for rows.Next() {
        if err2 := rows.Scan(&reply.Id, &reply.Key, &reply.Reply); err2 != nil {
            return nil, err2
        }
        break
    }
    if reply.Id == 0 {
        reply = nil
    }
    return reply, nil
}

func InsertFixReply(key string, reply string, db *sql.DB) error {
    temp, _ := GetFixReplyByKey(key, db)
    if temp != nil {
        // 不等于空则修改 使用id做where条件
        _, err := db.Exec("UPDATE fix_reply SET `reply` = ? WHERE `id` = ?", reply, temp.Id)
        if err != nil {
            return err
        }
    } else {
        _, err := db.Exec("INSERT INTO fix_reply (`key`, `reply`) VALUES (?, ?)", key, reply)
        if err != nil {
            return err
        }
    }
    return nil
}

func DeleteFixReply(key string, db *sql.DB) (string, error) {
    temp, _ := GetFixReplyByKey(key, db)
    if temp != nil {
        // 不等于空则删除 使用id做where条件
        _, err := db.Exec("DELETE FROM fix_reply WHERE `id` = ?", temp.Id)
        if err != nil {
            return "删除失败！", err
        }
    } else {
        return "删除的key不存在哦！", nil
    }
    return "删除成功！", nil
}
```

### 参考文档

[QQ 机器人 | QQ机器人文档](https://bot.q.qq.com/wiki/)

[账号注册 | QQ机器人文档](https://bot.q.qq.com/wiki/develop/api-v2/)

[GitHub - tencent-connect/botgo: QQ频道机器人 GOSDK](https://github.com/tencent-connect/botgo)
