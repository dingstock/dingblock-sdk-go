package modal

type SyncUserInfoRuquest struct {
	AppId     string `json:"appId"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Method    string `json:"method"`
	BizData   string `json:"bizData"`
	Sign      string `json:"sign"`
}

type SyncUserInfoBizData struct {
	PhoneNumber string `json:"phoneNumber"`
	UserId      string `json:"userId"`
	RealName    string `json:"realName"`
	IdCard      string `json:"idCard"`
}
