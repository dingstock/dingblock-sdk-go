package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var Request = request{}

type request struct {
}

func (r *request) Post(url string, data interface{}) (err error, result string) {
	// 超时时间：5秒
	jsonStr, _ := json.Marshal(data)
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		//model.Report.ReportError(err)
		return err, ""
	}
	defer func(Body io.ReadCloser) {
		//model.Report.ReportError(err)
		if e := Body.Close(); e != nil {
			println(e.Error())
		}
	}(resp.Body)

	byts, _ := ioutil.ReadAll(resp.Body)
	return nil, string(byts)
}
