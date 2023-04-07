package modal

type SimpleVerifyRequest struct {
	AppId     string `json:"appId"`
	Method    string `json:"method"`
	Nonce     string `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	BizData   string `json:"bizData"`
}

type SimpleVerifyResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Nonce     string `json:"nonce"`
	BizData   string `json:"bizData"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}
