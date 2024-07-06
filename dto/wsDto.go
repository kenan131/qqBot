package dto

type Session struct {
	ID      string
	URL     string
	Token   Token
	Intent  Intent
	LastSeq uint32
	Cnt     int
}

// WSUser 当前连接的用户信息
type WSUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Bot      bool   `json:"bot"`
}
