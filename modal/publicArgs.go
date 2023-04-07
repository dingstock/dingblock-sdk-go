package modal

type PublicResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Sign      string `json:"sign"`
	BizData   string `json:"bizData"`
}

type PublicRequest struct {
	AppId     string `json:"appId"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Method    string `json:"method"`
	BizData   string `json:"bizData"`
	Sign      string `json:"sign"`
}
