package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

// 接收
type Msg struct {
	Msg string `json:"msg"`
}

// 回傳
type Body struct {
	Larry string `json:"Larry"`
}

func GetAPi(method string, port string, header string, num string) {

	url := "http://192.168.204.76:" + port

	bd := Body{Larry: num}
	b, err := json.Marshal(bd)
	if err != nil {
		return
	}
	payload := strings.NewReader(string(b))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Larry", header)

	res, err := client.Do(req)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		return
	}
	var a Msg
	err = json.Unmarshal(body, &a)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Unmarshal Error")
		return
	}
	fmt.Println(a)
}
