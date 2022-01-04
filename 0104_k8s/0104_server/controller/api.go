package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type RegisterData struct {
	Username string `json:"username" form:"username" binding:"required,min=6,max=12"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=12"`
}

type LoginData struct {
	Username string `json:"username" form:"username" binding:"required,min=6,max=12"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=12"`
}
type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Page interface{} `json:"page"`
}
type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	UserId string `json:"userid"`
}

const LoginAPI = "http://loginapi:8082"
const TodoAPI = "http://todolist:8083"

// Login-----------------------------------------
func Login(c *gin.Context) {
	var login LoginData
	// 使用者傳帳密進來，解析到struct
	err := c.ShouldBindJSON(&login)
	if err != nil {
		util.Msg(c, 2, "帳號密碼格式錯誤", nil)
		return
	}
	url := LoginAPI + "/login"
	method := "POST"
	// 把前端要傳送的資料放進Body
	b, err := json.Marshal(login)
	if err != nil {
		util.Msg(c, 2, "請求無效", nil)
		return
	}
	payload := strings.NewReader(string(b))
	// 組成新的Request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	var msg util.API_Error
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Unmarshal Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	t := msg.Data
	if t != nil {
		tk := t.(string)
		util.Msg(c, msg.Code, msg.Msg, msg.Data)
		c.SetCookie("token", tk, 3600, "/", "localhost", false, true)
	} else {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效 body Error", nil)
		return
	}

}

func Register(c *gin.Context) {
	var reg RegisterData
	// 使用者傳帳密進來，解析到struct
	err := c.ShouldBindJSON(&reg)
	if err != nil {
		util.Msg(c, util.Code_Param_Invalid, "帳號密碼格式錯誤", nil)
		return
	}
	url := LoginAPI + "/Register"
	method := "POST"

	b, err := json.Marshal(reg)
	if err != nil {
		return
	}
	payload := strings.NewReader(string(b))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	var msg util.API_Error
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Unmarshal Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	util.Msg(c, msg.Code, msg.Msg, nil)
}

func MiddlewareAuth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		util.Msg(c, util.Code_Param_Invalid, "請重新登錄", nil)
		return
	}
	url := LoginAPI + "/middlewareAuth"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	res, err := client.Do(req)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	if string(body) != "" {
		var msg util.API_Error
		err = json.Unmarshal(body, &msg)
		if err != nil {
			log.Warn().Caller().Err(err).Msg("Unmarshal Error")
			util.Msg(c, 5, "請求無效", nil)
			return
		}
		util.Msg(c, msg.Code, msg.Msg, nil)
	}

}
func Logout(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		util.Msg(c, util.Code_Param_Invalid, "請重新登錄", nil)
		return
	}
	url := LoginAPI + "/logout"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	res, err := client.Do(req)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	var msg util.API_Error
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Unmarshal Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	util.Msg(c, msg.Code, msg.Msg, token)
	c.SetCookie("token", token, -1, "/", "localhost", false, true)
}

// TodoList--------------------------------------
func Insert(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		util.Msg(c, util.Code_Param_Invalid, "請重新登錄", nil)
		return
	}
	var todo Todo
	// 使用者傳參數進來，解析到struct
	err = c.ShouldBindJSON(&todo)
	if err != nil {
		util.Msg(c, 2, "帳號密碼格式錯誤", nil)
		return
	}
	url := TodoAPI + "/insert"
	method := "POST"

	b, err := json.Marshal(todo)
	if err != nil {
		util.Msg(c, 2, "請求無效", nil)
		return
	}
	payload := strings.NewReader(string(b))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	res, err := client.Do(req)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	var msg util.API_Error
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Unmarshal Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	util.Msg(c, msg.Code, msg.Msg, nil)
}

// func Get(c *gin.Context) {

// }

func Getpage(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		util.Msg(c, util.Code_Param_Invalid, "請重新登錄", nil)
		return
	}

	page := c.Query("page")

	url := TodoAPI + "/getpage?page=" + page
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	res, err := client.Do(req)
	if err != nil {
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	var msg PageData
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Unmarshal Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Code": msg.Code,
		"Msg":  msg.Msg,
		"Data": msg.Data,
		"Page": msg.Page,
	})
}

func Put(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		util.Msg(c, 2, "請重新登錄", nil)
		return
	}
	id, ok := c.Params.Get("id")
	if !ok {
		util.Msg(c, 2, "id無效", nil)
		return
	}
	var todo Todo
	// 使用者傳參數進來，解析到struct
	err = c.ShouldBindJSON(&todo)
	if err != nil {
		util.Msg(c, 2, "帳號密碼格式錯誤", nil)
		return
	}
	url := TodoAPI + "/put/" + id
	method := "PUT"

	b, err := json.Marshal(todo)
	if err != nil {
		util.Msg(c, 2, "請求無效", nil)
		return
	}
	payload := strings.NewReader(string(b))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	res, err := client.Do(req)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	var msg util.API_Error
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Unmarshal Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	util.Msg(c, msg.Code, msg.Msg, nil)
}

func Del(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		util.Msg(c, 2, "請重新登錄", nil)
		return
	}
	id, ok := c.Params.Get("id")
	if !ok {
		util.Msg(c, 2, "id無效", nil)
		return
	}
	url := TodoAPI + "/del/" + id
	method := "DELETE"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "token", Value: token})

	res, err := client.Do(req)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("NewRequest Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Body Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	var msg util.API_Error
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Unmarshal Error")
		util.Msg(c, 5, "請求無效", nil)
		return
	}
	util.Msg(c, msg.Code, msg.Msg, nil)
}
