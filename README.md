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

### 功能介绍

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

#### 指令三：默认回复


